package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	USER_PROFILE_COLLECTION = "user_profiles"
	USER_DATA_COLLECTION    = "user_data"
	PAGE_THEME_COLLECTION   = "page_theme"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

/* Makes a database connection */
func Connect() (err error) {
	var (
		uri    string = os.Getenv("MONGO_CONNECTION_URI")
		dbname string = os.Getenv("DB_NAME")
	)

	client, err = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	db = client.Database(dbname)
	return err
}

/* Closes active database connection */
func Disconnect() {
	if client != nil {
		client.Disconnect(context.TODO())
	}
}
