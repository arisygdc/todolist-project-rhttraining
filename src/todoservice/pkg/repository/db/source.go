package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/todoservice/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	DriverMongo      string = "mongodb"
	DriverCloudMongo string = "mongodb+srv"
	CollectionList   string = "list"
)

type IListColl interface {
	// InsertList param context, user id, todo, done.
	// return _id after inserted list
	InsertList(context.Context, uuid.UUID, string, bool) (primitive.ObjectID, error)
	GetList(context.Context, uuid.UUID) ([]List, error)
}

type DBSource struct {
	Client *mongo.Client
	dbName string
}

// NewDatabaseSource create new client connection database mongodb to run CRUD operation
// constant default port mongodb 27017
// Driver mongodb connect to self-hosted database
// Driver mongodb+srv connect to cloud mongodb.com
func NewDatabaseSource(cfg config.DbConfig) (DBSource, error) {
	var serverAPIOptions *options.ServerAPIOptions = nil
	var connString = fmt.Sprintf("%s://%s:%s@%s", cfg.Driver, cfg.User, cfg.Password, cfg.Host)

	if cfg.Driver == DriverMongo {
		connString = fmt.Sprintf("%s:%s", connString, "27017")
	}

	if cfg.Driver == DriverCloudMongo {
		serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)
		connString = fmt.Sprintf("%s/?retryWrites=true&w=majority", connString)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := NewConnection(ctx, connString, serverAPIOptions)
	if err != nil {
		return DBSource{}, err
	}

	return DBSource{
		Client: conn,
		dbName: cfg.DbName,
	}, err
}

// NewConnection given connString is uri database and return connection to mongodb
func NewConnection(ctx context.Context, connString string, opt ...*options.ServerAPIOptions) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connString)
	for _, v := range opt {
		if v != nil {
			clientOptions.SetServerAPIOptions(v)
		}
	}

	var conn *mongo.Client

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return conn, nil
}

func (ds DBSource) InsertList(ctx context.Context, userId uuid.UUID, todo string, done bool) (primitive.ObjectID, error) {
	doc := bson.D{
		{"auth_id", userId},
		{"todo", todo},
		{"done", done},
	}

	res, err := ds.Client.Database(ds.dbName).Collection(CollectionList).InsertOne(ctx, doc)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	if id == primitive.NilObjectID {
		return primitive.NilObjectID, fmt.Errorf("something went wrong")
	}

	return id, nil
}

type List struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID uuid.UUID          `json:"auth_id" bson:"auth_id"`
	Todo   string             `json:"todo" bson:"todo"`
	Done   bool               `json:"done" bson:"done"`
}

func (ds DBSource) GetList(ctx context.Context, userId uuid.UUID) ([]List, error) {
	filter := primitive.D{
		{"auth_id", userId},
	}

	var resultCollection []List

	cursor, err := ds.Client.Database(ds.dbName).Collection(CollectionList).Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return resultCollection, err
	}

	for cursor.Next(ctx) {
		result := new(List)
		if err := cursor.Decode(result); err != nil {
			return []List{}, err
		}
		resultCollection = append(resultCollection, *result)
	}

	return resultCollection, err
}
