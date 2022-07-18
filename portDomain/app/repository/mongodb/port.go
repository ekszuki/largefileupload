package mongodb

import (
	"context"
	"ekszuki/uploader/portDomain/app/contracts"
	"ekszuki/uploader/portDomain/app/models"
	"ekszuki/uploader/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database string = utils.GetEnv("MONGO_DB_NAME", "uploads")
)

type PortRepository struct {
	col *mongo.Collection
}

// FindByKey implements contracts.PortRepository
func (r *PortRepository) FindByKey(ctx context.Context, key string) (*models.Port, error) {
	var dbObj Port
	err := r.col.FindOne(ctx, bson.D{{"_id", key}}).Decode(&dbObj)
	if err != nil {
		return &models.Port{}, err
	}

	return dbObj.ToDomain(), nil
}

// SaveOrUpdate implements contracts.PortRepository
func (r *PortRepository) SaveOrUpdate(ctx context.Context, port *models.Port) error {
	var dbObj Port
	if err := dbObj.fromDomain(port); err != nil {
		return fmt.Errorf("could not convert port domain into mongodb obj: %v", err)
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": dbObj.Key}, bson.M{"$set": dbObj}, opts)

	return err
}

func NewPortRepository(ctx context.Context) (contracts.PortRepository, error) {
	client, err := getConnection(ctx)
	if err != nil {
		return nil, err
	}

	return &PortRepository{
		col: client.Database(database).Collection("ports"),
	}, nil
}
