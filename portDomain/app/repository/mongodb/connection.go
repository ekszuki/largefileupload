package mongodb

import (
	"context"
	"ekszuki/uploader/utils"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection(ctx context.Context) (*mongo.Client, error) {
	host := utils.GetEnv("MONGO_HOST", "localhost")
	port, err := strconv.Atoi(utils.GetEnv("MONGO_PORT", "27017"))
	if err != nil {
		port = 27017
	}

	mongoURL := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	return client, err
}
