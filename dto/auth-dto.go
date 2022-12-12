package dto

import "github.com/zhuliminl/easyrn-server/entity"

type ResRegister struct {
	Token string `json:"token"`
	User  entity.User
}
