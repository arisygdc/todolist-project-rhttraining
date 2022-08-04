package core

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/repository/db"
	tokenProvider "github.com/todolist-project-rhttraining/src/authservice/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	dbRepo     repository.DatabaseRepo
	tokenMaker tokenProvider.JWToken
}

func NewAuthService(dbRepo repository.DatabaseRepo, tokenMaker tokenProvider.JWToken) AuthService {
	return AuthService{
		dbRepo:     dbRepo,
		tokenMaker: tokenMaker,
	}
}

func (as AuthService) AddAuth(ctx context.Context, username, password, email string) (string, error) {
	bitPassword := []byte(password)
	bitHashedPassword, err := bcrypt.GenerateFromPassword(bitPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	param := db.AddAuthParams{
		ID:       uuid.New(),
		Username: username,
		Password: string(bitHashedPassword),
		Email:    email,
	}

	err = as.dbRepo.Query().AddAuth(ctx, param)

	if err != nil {
		return "", err
	}

	return param.ID.String(), nil
}

func (as AuthService) Login(ctx context.Context, username string, password string) (string, error) {
	authRes, err := as.dbRepo.Query().GetAuth(ctx, username)
	if err != nil {
		return "", err
	}

	bitHashedPassword := []byte(authRes.Password)
	bitLoginPassword := []byte(password)

	err = bcrypt.CompareHashAndPassword(bitHashedPassword, bitLoginPassword)
	if err != nil {
		return "", fmt.Errorf("wrong username or password")
	}

	payload := tokenProvider.NewSessionPayload(authRes.Username)
	return as.tokenMaker.Generate(payload)
}

// VerifyToken @param token session token
// return auth id string
func (as AuthService) VerifyToken(ctx context.Context, token string) (string, error) {
	tokenClaim, err := tokenProvider.ParseWithClaimSession(as.tokenMaker, token)
	if err != nil {
		return "", err
	}

	authRes, err := as.dbRepo.Query().GetAuth(ctx, tokenClaim.Username)
	return authRes.ID.String(), err
}
