package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginUseCase struct {
	UsersGateway gateway.UsersGateway
	JwtSecret    string
}

func NewLoginUseCase(userGateway gateway.UsersGateway, jwtSecret string) *LoginUseCase {
	return &LoginUseCase{
		UsersGateway: userGateway,
		JwtSecret:    jwtSecret,
	}
}

func (u *LoginUseCase) Execute(ctx context.Context, input dto.LoginInputDTO) (*dto.LoginOutputDTO, error) {
	usr, err := u.UsersGateway.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(input.Password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"username": usr.Username,
		"id":       usr.ID.String(),
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	// Create token
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := jwtClaims.SignedString([]byte(u.JwtSecret))
	if err != nil {
		return nil, errors.New("failed to generate token --> " + err.Error())
	}
	return &dto.LoginOutputDTO{Token: t}, nil
}
