package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/neel229/linktree-clone/internal/shortener"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepo struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int64) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		return nil, fmt.Errorf("error creating mongo client: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewMongoRepo(mongoURL, mongoDB string, mongoTimeout int64) (shortener.RedirectRepository, error) {
	repo := &mongoRepo{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, fmt.Errorf("error creating new mongo repo: %v", err)
	}
	repo.client = client
	return repo, nil
}

func (mr *mongoRepo) Find(code string) (*shortener.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	redirect := &shortener.Redirect{}
	coll := mr.client.Database(mr.database).Collection("redirects")
	filter := bson.M{"code": code}
	err := coll.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no such redirect found: %v", err)
		}
		return nil, err
	}
	return redirect, nil
}

func (mr *mongoRepo) Store(redirect *shortener.Redirect) error {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	coll := mr.client.Database(mr.database).Collection("redirects")
	_, err := coll.InsertOne(ctx, bson.M{
		"code":    redirect.Code,
		"handles": redirect.Handles,
	})
	if err != nil {
		return fmt.Errorf("error storing the redirect: %v", err)
	}
	return nil
}
