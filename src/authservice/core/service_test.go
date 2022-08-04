package core

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/todolist-project-rhttraining/src/authservice/config"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/authservice/pkg/token"
	"log"
	"testing"
)

func TestService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfiguration()

	repo, err := repository.NewDatabaseRepo(ctx, cfg.Database)
	assert.NoError(t, err)

	tokenMaker := token.NewJWT("f2d448191b404142bd8a45f8a6175b6d567cacec9965bbed355149e5af7d6dad")

	svc := NewAuthService(repo, tokenMaker)

	table := struct {
		username string
		password string
		email    string
	}{
		username: "ifbawieb foanej",
		password: "benalu88ooo",
		email:    "gaerr@benalu.com",
	}

	idAdded, err := svc.AddAuth(ctx, table.username, table.password, table.email)
	assert.NoError(t, err)
	log.Printf("inserted id: %s", idAdded)
}
