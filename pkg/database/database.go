package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBError struct {
	message string
}

func NewDBError(message string) *DBError {
	return &DBError{
		message: message,
	}
}

func (err *DBError) Error() string {
	return err.message
}

type database struct {
	client   *mongo.Client
	database *mongo.Database
}

func InitDatabase(ctx context.Context) (*database, error) {
	connectChan := make(chan bool)
	defer close(connectChan)

	c, cancel := context.WithTimeout(ctx, 10*time.Second)

	var database database

	var errResult error
	go func() {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URL")))
		database.client = client

		if err == nil {
			database.database = client.Database(os.Getenv("MONGO_DB_NAME"))
		}

		errResult = err

		connectChan <- true
		cancel()
	}()

	select {
	case <-connectChan:
		return &database, errResult
	case <-c.Done():
		return nil, NewDBError("timeout after long waiting")
	}
}

func (db *database) GetDatabase() *mongo.Database {
	return db.database
}

func (db *database) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
