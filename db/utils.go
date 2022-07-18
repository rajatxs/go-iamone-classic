package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/* Returns reference of database collection */
func Collection(c string) *mongo.Collection {
	return db.Collection(c)
}

/** Finds single document from given collection */
func FindOneDoc(cname string, filter primitive.D, ref interface{}) error {
	return Collection(cname).FindOne(context.TODO(), filter, nil).Decode(ref)
}
