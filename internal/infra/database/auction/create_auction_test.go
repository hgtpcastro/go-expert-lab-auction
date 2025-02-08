package auction

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGivenAValidParams_WhenCreateAuction_ThenShouldCloseAutomaticallyAfterDefiniteTime(t *testing.T) {
	mockCollection := new(database.MockCollection)
	repo := &AuctionRepository{
		Collection:          mockCollection,
		auctionInterval:     time.Millisecond * 50,
		auctionEndTimeMutex: &sync.Mutex{},
	}

	auction := &auction_entity.Auction{
		Id: "123",
	}

	mockCollection.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).
		Return(&mongo.InsertOneResult{InsertedID: "123"}, nil)
	mockCollection.On("UpdateByID", mock.Anything, "123", mock.MatchedBy(func(update interface{}) bool {
		updateMap, ok := update.(bson.M)
		if !ok {
			return false
		}
		setMap, ok := updateMap["$set"].(bson.M)
		if !ok {
			return false
		}
		return setMap["status"] == auction_entity.Completed
	}), mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil)

	err := repo.CreateAuction(context.Background(), auction)
	assert.Nil(t, err)

	time.Sleep(time.Millisecond * 100)

	mockCollection.AssertExpectations(t)
}
