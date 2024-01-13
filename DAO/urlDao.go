package dao

import (
	models "URLShortener/Models"
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URLDao struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewURLDao() (*URLDao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, errors.New("MONGODB_URI no está definida en las variables de entorno")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("nombreDeTuBaseDeDatos")

	return &URLDao{client: client, db: db}, nil
}
func (dao *URLDao) Save(url *models.URL) error {
	collection := dao.db.Collection("URLS")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, url)
	return err
}

func (dao *URLDao) Get(shortURL string) (*models.URL, error) {
	collection := dao.db.Collection("URLS")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url models.URL
	err := collection.FindOne(ctx, bson.M{"shortened_url": shortURL}).Decode(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (dao *URLDao) Delete(shortURL string) error {
	collection := dao.db.Collection("URLS")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"shortened_url": shortURL})
	return err
}
