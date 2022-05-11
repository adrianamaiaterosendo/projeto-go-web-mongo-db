package db

import (
	"context"

	"github.com/adrianamaiaterosendo/projeto-go-web-mongo-db.git/model"
)

func Insert(collection string, data model.Subscription) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("tesouro").Collection(collection)

	_, err := c.InsertOne(context.Background(), data)
	return err
}
