package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Database
}

var (
	ctx context.Context
)

// Connect connects to mongodb
func Connect(connString, database string) (Database, error) {
	// Initiate a session with Mongo
	ctx = context.Background()
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		return Database{}, err
	}
	return Database{db: conn.Database(database)}, nil
}

// Get gets an item given a table and keys
func (db Database) Get(t string, keys map[string]interface{}, out interface{}) error {
	collection := db.db.Collection(t)

	res := collection.FindOne(ctx, keys)
	if res.Err() != nil {
		return res.Err()
	}
	return res.Decode(out)
}

// List retrieves all objects from a collection
// The n parameter is the limit of objects returned, use 0 for no limit
// And the out parameter must be a pointer to a slice
func (db Database) List(table string, keys map[string]interface{}, n int64, out interface{}) error {
	collection := db.db.Collection(table)

	opts := options.FindOptions{Limit: &n}
	cur, err := collection.Find(ctx, keys, &opts)
	if err != nil {
		return err
	}

	return cur.All(ctx, out)
}

// Put inserts an item with the given input
func (db Database) Put(obj interface{}, table string) error {
	collection := db.db.Collection(table)
	_, err := collection.InsertOne(ctx, obj)
	return err
}

// Delete erases at most one element from database
func (db Database) Delete(table string, keys map[string]interface{}) error {
	collection := db.db.Collection(table)

	_, err := collection.DeleteOne(ctx, keys)
	return err
}

// Update changes fields of one object
func (db Database) Update(table string, keys, obj map[string]interface{}) error {
	collection := db.db.Collection(table)
	// Create the update expression
	update := map[string]interface{}{"$set": obj}
	return collection.FindOneAndUpdate(ctx, keys, update).Err()
}

// Replace either inserts or updates an object
func (db Database) Replace(table string, keys map[string]interface{}, obj interface{}) error {
	collection := db.db.Collection(table)
	upsert := true
	opts := options.ReplaceOptions{Upsert: &upsert}
	_, err := collection.ReplaceOne(ctx, keys, obj, &opts)
	return err
}

// RawUpdate changes fields of one object matching the keys
// Here the update if any MongoDB update expression
func (db Database) RawUpdate(table string, keys, up map[string]interface{}) error {
	collection := db.db.Collection(table)

	return collection.FindOneAndUpdate(ctx, keys, up).Err()
}
