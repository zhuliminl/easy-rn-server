package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/easyrn-server/constError"
	"github.com/zhuliminl/easyrn-server/constant"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/entity"
	"github.com/zhuliminl/easyrn-server/helper"
	"github.com/zhuliminl/easyrn-server/repository"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx333789da24c1f349"
	appSecret       = "62c790cced311e81907eea2d0b3a6310"
	appBaseLink     = "http://www.baidu.com"
)

var ctx = context.Background()

type WechatService interface {
	GetOpenId(wechatCode dto.WechatCodeDto) (dto.ResJsCode2session, error)
	GenerateAppLink() (dto.WechatAppLink, error)
	ScanOver(loginSessionId string) error
	GetMiniLinkStatus(loginSessionId string) (dto.MiniLinkStatus, error)
	LoginWithEncryptedPhoneData(wxLoginData dto.WxLoginData) (dto.ResWxLogin, error)
	GetUserByLoginSessionId(loginSessionId string) (entity.User, error)
}

type wechatService struct {
	userRepository repository.UserRepository
	userService    UserService
	rdb            redis.Client
}

func (service wechatService) GetMiniLinkStatus(loginSessionId string) (dto.MiniLinkStatus, error) {
	var miniLinkStatus dto.MiniLinkStatus
	value, err := service.rdb.Get(ctx, loginSessionId+constant.PrefixLogin).Result()
	if err == redis.Nil {
		return miniLinkStatus, constError.NewWechatLoginUidNotFound(err, "uid key 不存在")
	} else if err != nil {
		return miniLinkStatus, err
	}
	miniLinkStatus.Status = value
	return miniLinkStatus, nil
}

func (service wechatService) GenerateAppLink() (dto.WechatAppLink, error) {
	loginSessionId := uuid.NewV4().String()
	var linkDto dto.WechatAppLink
	linkDto.Link = appBaseLink + "?login_session_id=" + loginSessionId
	linkDto.LoginSessionId = loginSessionId
	err := service.rdb.Set(ctx, loginSessionId+constant.PrefixLogin, constant.WechatLoginScanReady, constant.MiniLoginExpiredMinute*time.Minute).Err()
	if err != nil {
		// better panic
		return linkDto, err
	}

	return linkDto, nil
}

func (service wechatService) ScanOver(loginSessionId string) error {
	_, err := service.rdb.Get(ctx, loginSessionId+constant.PrefixLogin).Result()
	if err == redis.Nil {
		return constError.NewWechatLoginUidNotFound(err, "loginSessionId key 不存在")
	} else if err != nil {
		return err
	}

	err = service.rdb.Set(ctx, loginSessionId+constant.PrefixLogin, constant.WechatLoginScanOver, constant.MiniLoginExpiredMinute*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (service wechatService) GetOpenId(wechatCodeDto dto.WechatCodeDto) (dto.ResJsCode2session, error) {
	var session dto.ResJsCode2session
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, wechatCodeDto.Code)
	log.Println("code2sessionURL", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("GetOpenIdHttpGetError", err)
		return session, err
	}

	wxMap := map[string]interface{}{"openid": "", "session_key": "", "errcode": 0, "errmsg": ""}

	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		fmt.Println("GetOpenIdHttpDecodeError", err)
		return session, err
	}
	defer resp.Body.Close()

	log.Println("code2sessionWechatMap", wxMap)

	var errorCode int
	if _, ok := wxMap["errcode"].(int); ok {
		errorCode = wxMap["errcode"].(int)
	} else {
		errorCode = int(wxMap["errcode"].(float64))
	}

	session.OpenId = wxMap["openid"].(string)
	session.SessionKey = wxMap["session_key"].(string)
	session.Errcode = errorCode
	session.Errmsg = wxMap["errmsg"].(string)

	if session.SessionKey != "" {
		// 绑定 sessionKey 到 loginSessionId
		err := service.rdb.Set(ctx, wechatCodeDto.LoginSessionId+constant.PrefixWechatSessionKey, session.SessionKey, constant.MiniLoginExpiredMinute*time.Minute).Err()
		if err != nil {
			// better panic
			return session, err
		}
	}

	if session.OpenId != "" {
		// 绑定 openId
		err := service.rdb.Set(ctx, wechatCodeDto.LoginSessionId+constant.PrefixWechatOpenId, session.OpenId, constant.MiniLoginExpiredMinute*time.Minute).Err()
		if err != nil {
			// better panic
			return session, err
		}
	}

	return session, nil
}

func (service wechatService) LoginWithEncryptedPhoneData(wxLoginData dto.WxLoginData) (dto.ResWxLogin, error) {
	var resWxLogin dto.ResWxLogin
	_, err := service.rdb.Get(ctx, wxLoginData.LoginSessionId+constant.PrefixLogin).Result()
	if err == redis.Nil {
		return resWxLogin, constError.NewWechatLoginUidNotFound(err, "loginSessionId key 不存在")
	} else if err != nil {
		return resWxLogin, err
	}

	sessionKey, err := service.rdb.Get(ctx, wxLoginData.LoginSessionId+constant.PrefixWechatSessionKey).Result()
	if err == redis.Nil {
		return resWxLogin, errors.New("wechat sessionKey 不存在，可能已过期")
	} else if err != nil {
		return resWxLogin, err
	}

	openId, err := service.rdb.Get(ctx, wxLoginData.LoginSessionId+constant.PrefixWechatOpenId).Result()
	if err == redis.Nil {
		return resWxLogin, errors.New("wechat openId 不存在，可能已过期")
	} else if err != nil {
		return resWxLogin, err
	}

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	aesIv, err := base64.StdEncoding.DecodeString(wxLoginData.Iv)

	base64Ciphertext := wxLoginData.EncryptedData
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		log.Println("WechatDecodeStringError", err)
	}
	raw, err := helper.AESDecryptData(ciphertext, aesKey, aesIv)
	if err != nil {
		log.Println("WechatAESDecryptDataError", err)
	}

	var resOfNumber dto.WxGetPhoneNumberRes
	log.Println("AESDecryptData", string(raw))
	err = json.Unmarshal(raw, &resOfNumber)
	if err != nil {
		return resWxLogin, err
	}

	resWxLogin.Phone = resOfNumber.PurePhoneNumber
	// 将该用户写入数据库
	err = service.CreateWxUser(openId, resWxLogin.Phone)
	if err != nil {
		return resWxLogin, err
	}

	return resWxLogin, nil
}

func (service wechatService) CreateWxUser(openId string, phone string) error {
	// 检查是否登陆注册过，登陆过则更新用户
	user, err := service.userRepository.GetUserByPhone(phone)
	if constError.Is(err, constError.UserNotFound) {
		err := service.userRepository.CreateUser(entity.User{
			Username: "微信用户",
			Phone:    phone,
			OpenId:   openId,
		})

		if err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}
	}

	log.Println("现存用户", user)
	return nil
}

func (service wechatService) GetUserByLoginSessionId(loginSessionId string) (entity.User, error) {
	var user entity.User
	openId, err := service.rdb.Get(ctx, loginSessionId+constant.PrefixWechatOpenId).Result()
	if err == redis.Nil {
		return user, errors.New("wechat openId 不存在，可能已过期")
	} else if err != nil {
		return user, err
	}

	user, err = service.userRepository.GetUserByOpenId(openId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewWechatService(userRepo repository.UserRepository, userService UserService, rdb *redis.Client) WechatService {
	return &wechatService{
		userRepository: userRepo,
		userService:    userService,
		rdb:            *rdb,
	}
}
