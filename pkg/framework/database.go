package framework

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/scoir/canis/pkg/datastore"
	"github.com/scoir/canis/pkg/datastore/mongodb"
)

type DatabaseProvider func(name string) *mongo.Database

func NewDatabase(conf *Config) (DatabaseProvider, error) {

	mongoClient, err := getClient(conf)
	if err != nil {
		return nil, err
	}

	return func(name string) *mongo.Database {
		return mongoClient.Database(name)
	}, nil

}

func getClient(conf *Config) (*mongo.Client, error) {
	var err error
	tM := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistryBuilder().RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()
	clientOpts := options.Client().SetRegistry(reg).ApplyURI(conf.Mongo.URL)

	mongoClient, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, errors.Wrap(err, "error creating mongo client")
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to mongo")
	}

	return mongoClient, err
}

func (r *Config) Datastore() (datastore.Store, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.ds != nil {
		return r.ds, nil
	}

	var err error
	switch r.Database {
	case "mongo":
		r.ds, err = r.loadMongo()
	case "postgres":
		r.ds, err = r.loadPostgres()
	}

	return r.ds, errors.Wrap(err, "unable to get datastore from config")
}

func (r *Config) loadMongo() (datastore.Store, error) {
	mongoClient, err := getClient(r)
	if err != nil {
		return nil, err
	}

	return mongodb.NewStore(mongoClient.Database(r.Mongo.Database)), nil
}

func (r *Config) loadPostgres() (datastore.Store, error) {
	return nil, errors.New("not implemented")
}
