package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/guilhermecosales/security-service/internal/domain/model"
	"github.com/guilhermecosales/security-service/pkg/config"
)

type AuthenticationTokenServicer interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthenticationTokenService struct {
	secretKey  string
	aud        string
	iss        string
	expiration string
}

func NewAuthenticationTokenService(config config.JWTConfig) *AuthenticationTokenService {
	return &AuthenticationTokenService{
		secretKey:  config.SecretKey,
		aud:        config.AUD,
		iss:        config.ISS,
		expiration: config.Duration,
	}
}

func (s *AuthenticationTokenService) GenerateToken(userDetails *model.User) (string, error) {

	claims := s.buildClaims(userDetails.Email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthenticationTokenService) buildClaims(username string) jwt.MapClaims {
	return jwt.MapClaims{
		"iss": s.iss,
		"aud": s.aud,
		"sub": username,
		"exp": time.Now().Add(s.getExpiration()).Unix(),
		"nbf": time.Now().UTC().Unix(),
		"iat": time.Now().UTC().Unix(),
	}
}

func (s *AuthenticationTokenService) getExpiration() time.Duration {
	exp, _ := strconv.Atoi(s.expiration)
	return time.Duration(exp) * time.Second
}

func (s *AuthenticationTokenService) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
