package service

import (
	"errors"
	"time"

	model "HW/app/models"
	"HW/app/repo"
	"HW/config"

	"github.com/dgrijalva/jwt-go"
)

type (
	UserService struct {
		repo repo.UserRepository
	}

	JWTAuthService struct {
		cfg config.Config
	}

	JwtClaims struct {
		UserId int    `json:"id,omitempty"`
		Email  string `json:"email,omitempty"`
		jwt.StandardClaims
	}
)

func NewUserService(r repo.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) GetUser(email string, password string) *model.User {
	return s.repo.GetUserEmailPassword(email, password)
}

func (s *UserService) UserHasRole(id int, role string) bool {
	user := s.repo.GetUserWithRoles(id)
	if user == nil {
		return false
	}

	for i := range user.Roles {
		if user.Roles[i].Name == role {
			return true
		}
	}

	return false
}

func (s *UserService) CreateUser(user *model.User) error {
	userExists := s.repo.GetUserEmail(user.Email)
	if userExists != nil {
		return errors.New("user with same email already exist in database")
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return errors.New("an unknown error occurred during operation")
	}

	return nil
}
func NewJWTAuthService(c config.Config) *JWTAuthService {
	return &JWTAuthService{
		cfg: c,
	}
}

func (s *JWTAuthService) VerifyToken(tokenString string) (bool, *JwtClaims) {
	claims := &JwtClaims{}
	token, _ := getTokenFromString(tokenString, s.cfg.JWTConfig.JWTSecret, claims)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func (s *JWTAuthService) CreateToken(user model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"iss":   s.cfg.JWTConfig.JWTIss,
		"exp":   time.Now().Add(s.cfg.JWTConfig.JWTExp).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.JWTConfig.JWTSecret))
	if err != nil {
		return nil, errors.New("unable to create signed token")
	}

	return &tokenString, nil
}

func getTokenFromString(tokenString string, secret string, claims *JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
}
