package auth

//1. GENERATE TOKEN
//2. VALIDASI

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y") //secret key jgn ditampilkan

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //algoritma HS25

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
	//Security

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
