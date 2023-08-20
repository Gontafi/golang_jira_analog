package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"math/rand"
	"time"
)

const (
	salt          = "gihodrhgoidrg4259baf8932"
	signingKey    = "ghsdjkrghskdrhgkjsdrgjk3234uFSIOVFH"
	tokenTTL      = 12 * time.Hour
	resetCodeSalt = "resetcode_salt" // Salt for reset code generation
	resetCodeTTL  = 15 * time.Minute // Reset code timeout
)

type AuthService struct {
	r *repos.AuthRepository
}

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthService(repo *repos.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generateHashPasswordHash(user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.r.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.r.GetUser(username, generateHashPasswordHash(password))

	if err != nil {
		slog.Error("Error:", err)
		return "", err
	}
	expTime := time.Now().Add(tokenTTL)
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("token :", token)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func generateHashPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ForgotPassword(username string) (string, error) {
	user, err := s.r.GetUserByUsername(username)
	if err != nil {
		slog.Error("Error:", err)
		return "", err
	}

	resetCode := generateResetCode()
	err = s.r.StoreResetCode(user.ID, resetCode)
	if err != nil {
		slog.Error("Failed to store reset code:", err)
		return "", err
	}

	return resetCode, nil
}

func (s *AuthService) VerifyResetCode(username string, code string) error {
	user, err := s.r.GetUserByUsername(username)

	storedCode, err := s.r.GetResetCode(user.ID)
	if err != nil {
		slog.Error("Error:", err)
		return err
	}
	//TODO
	if storedCode != code {
		return errors.New("invalid reset code")
	}

	return nil
}

func generateResetCode() string {
	const charset = "0123456789"
	codeLength := 6
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func (s *AuthService) ChangeUserPassword(username string, password string) error {
	err := s.r.ChangePassword(username, password)
	if err != nil {
		return err
	}

	return nil
}
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
