package db

import (
	"go-blockchain-scope/env"

	. "go-blockchain-scope/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	env *env.Env

	client *mongo.Client
	db     *mongo.Database

	block *mongo.Collection
	tx    *mongo.Collection
}

type DBImpl interface {
}

func NewDB(env *env.Env) (DBImpl, error) {
	d := &DB{env: env}

	ctx := Context()
	var err error

	if d.client, err = mongo.Connect(ctx, options.Client().ApplyURI(env.DB.Uri)); err != nil {
		panic(err)
	} else if err = d.client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		d.db = d.client.Database(env.DB.DB)

		d.block = d.db.Collection(env.DB.Block)
		d.tx = d.db.Collection(env.DB.Tx)

		return d, nil
	}
}
