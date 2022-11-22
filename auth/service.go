package auth

import (
	"errors"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID int, userName string, userRole string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {

}

func NewService() *jwtService {
	return &jwtService{}
}


func (s *jwtService) GenerateToken(userID int, userName string, userRole string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["user_name"] = userName
	claim["role"] = userRole
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	err:= godotenv.Load()
	if(err != nil) {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	key := os.Getenv("SECRET_KEY")
	signedToken, err := token.SignedString([]byte(key))

	if(err != nil) {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	err:= godotenv.Load()
	if(err != nil) {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	key := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if(!ok) {
			return nil, errors.New("invalid signing method")
		}

		return []byte(key), nil
	})

	if(err != nil) {
		return token, err
	}

	return token, nil
}