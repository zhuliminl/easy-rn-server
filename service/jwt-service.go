package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	zlog "github.com/rs/zerolog/log"
)

var secretKey = "Test1234"
var issuer = "saul"

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

type MyCustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (j jwtService) GenerateToken(userId string) string {
	log.Println("saul GenerateToken UserId", userId)
	if userId == "" {
		zlog.Error().Msgf("userId is empty")
		log.Panicln("no userId")
	}
	mySigningKey := []byte(j.secretKey)
	// Create the claims
	claims := MyCustomClaims{
		userId,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(1000 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Println("saul ==========>>> token %v %v", ss, err)
	if err != nil {
		log.Panic("GenerateTokenError", err)
	}
	return ss
}

func (j jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	return token, err
}
