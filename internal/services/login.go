package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	var (
		resp models.LoginResponse
		now  = time.Now()
	)

	userDetail, err := s.UserRepository.GerUserByUsername(ctx, req.Username)
	if err != nil {
		return resp, errors.Wrap(err, "failed get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return resp, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.UserName, userDetail.FullName, "token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	refresh_token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.UserName, userDetail.FullName, "refresh_token", now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refresh_token,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}

	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return resp, errors.Wrap(err, "failed to insert new session")
	}

	resp.ID = userDetail.ID
	resp.Username = userDetail.UserName
	resp.FullName = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refresh_token

	return resp, nil
}
