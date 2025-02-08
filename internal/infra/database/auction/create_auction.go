package auction

import (
	"context"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection          database.CollectionAPI
	auctionInterval     time.Duration
	auctionEndTimeMutex *sync.Mutex
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection:          database.Collection("auctions"),
		auctionInterval:     getAuctionInterval(),
		auctionEndTimeMutex: &sync.Mutex{},
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	go ar.ValidatesExpiredAuction(ctx, auctionEntity, ar.auctionInterval)

	return nil
}

func (ar *AuctionRepository) ValidatesExpiredAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction,
	duration time.Duration) {
	time.AfterFunc(duration, func() {
		ar.auctionEndTimeMutex.Lock()
		defer ar.auctionEndTimeMutex.Unlock()

		update := bson.M{
			"$set": bson.M{
				"status": auction_entity.Completed,
			},
		}

		_, err := ar.Collection.UpdateByID(ctx, auctionEntity.Id, update)
		if err != nil {
			logger.Error("Error trying to update auction", err)
		}

		logger.Info("Status auction update to Completed")
	})
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}
