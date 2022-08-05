package core

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/repository/db"
	tokenProvider "github.com/todolist-project-rhttraining/src/authservice/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	log.Println("--- Start Request AddAuth")
	log.Printf("username %s| password %s| email %s\n", username, password, email)

	defer func() { log.Println("--- End Request AddAuth") }()

	bitPassword := []byte(password)
	bitHashedPassword, err := bcrypt.GenerateFromPassword(bitPassword, bcrypt.DefaultCost)

	if err != nil {
		log.Println("error hash password")
		return "", err
	}

	param := db.AddAuthParams{
		ID:       uuid.New(),
		Username: username,
		Password: string(bitHashedPassword),
		Email:    email,
	}

	log.Printf("inserting auth with param %v\n", param)
	err = as.dbRepo.Query().AddAuth(ctx, param)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
		return "", err
	}

	log.Println("success insert")
	return param.ID.String(), nil
}

// Login generate token from username as session
func (as AuthService) Login(ctx context.Context, username string, password string) (string, error) {
	log.Println("--- Start Request Login")
	log.Printf("username %s| password %s\n", username, password)

	defer func() { log.Println("--- End Request Login") }()

	log.Printf("login using username: %s\n", username)
	authRes, err := as.dbRepo.Query().GetAuth(ctx, username)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
		return "", err
	}

	log.Printf("fetched: %v\n", authRes)

	bitHashedPassword := []byte(authRes.Password)
	bitLoginPassword := []byte(password)

	log.Printf("comparing hash. login password: %v, with hashed password %v\n", bitLoginPassword, bitHashedPassword)
	err = bcrypt.CompareHashAndPassword(bitHashedPassword, bitLoginPassword)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
		return "", fmt.Errorf("wrong username or password")
	}

	log.Printf("creating session token from username: %s\n", authRes.Username)

	payload := tokenProvider.NewSessionPayload(authRes.Username)
	sessionToken, err := as.tokenMaker.Generate(payload)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
		return "", err
	}

	log.Printf("session token: %s", sessionToken)

	return sessionToken, err
}

// VerifyToken @param token session token
// return auth id string
func (as AuthService) VerifyToken(ctx context.Context, token string) (string, error) {
	log.Println("--- Start Request VerifyToken")
	log.Printf("Token %s\n", token)

	defer func() { log.Println("--- End Request VerifyToken") }()

	log.Println("claiming token")

	var tokenClaim tokenProvider.SessionPayload
	err := as.tokenMaker.ParseWithClaimToken(token, &tokenClaim)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
		return "", err
	}

	log.Printf("token claimed, value: %v\n", tokenClaim)
	log.Printf("getting auth from username: %s\n", tokenClaim.Username)

	authRes, err := as.dbRepo.Query().GetAuth(ctx, tokenClaim.Username)
	if err != nil {
		log.Printf("failed with error: %v\n", err)
		return "", err
	}

	log.Printf("found: %v\n", authRes)

	return authRes.ID.String(), err
}
