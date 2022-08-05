package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/todolist-project-rhttraining/src/todoservice/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sync"
	"testing"
)

var wg sync.WaitGroup

type list struct {
	userId uuid.UUID
	todo   string
	done   bool
}

func TestDatabase(t *testing.T) {
	cfg := config.NewConfiguration()
	source, err := NewDatabaseSource(cfg.Database)
	assert.NoError(t, err)

	user1 := uuid.New()
	user2 := uuid.New()

	testTable := []list{
		{
			userId: user1,
			todo:   "bangun",
			done:   true,
		},
		{
			userId: user2,
			todo:   "kerja",
			done:   true,
		},
		{
			userId: user1,
			todo:   "ngopi",
			done:   false,
		},
		{
			userId: user2,
			todo:   "tidur",
			done:   false,
		},
		{
			userId: user1,
			todo:   "garuk garuk",
			done:   true,
		},
	}

	ctx := context.TODO()
	wg.Add(len(testTable))

	for _, v := range testTable {
		go func(ctx context.Context, source DBSource, v list, t *testing.T) {
			id, err := source.InsertList(ctx, v.userId, v.todo, v.done)
			assert.NoError(t, err)
			assert.NotEqual(t, primitive.NilObjectID, id)
			wg.Done()
		}(ctx, source, v, t)
	}

	wg.Wait()

	a, err := source.GetList(ctx, user1)
	assert.NoError(t, err)
	log.Println(a)

	b, err := source.GetList(ctx, user2)
	assert.NoError(t, err)
	log.Println(b)
}
