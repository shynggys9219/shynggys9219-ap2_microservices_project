package mongo

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Config struct {
	Database     string `env:"MONGO_DB"`
	URI          string `env:"MONGO_DB_URI"`
	Username     string `env:"MONGO_USERNAME"`
	Password     string `env:"MONGO_PWD"`
	ReplicaSet   string `env:"MONGO_DB_REPLICA_SET"`
	WriteConcern string `env:"MONGO_WRITE_CONCERN"`
	TLSFilePath  string `env:"MONGO_TLS_FILE_PATH"`
	TLSEnable    bool   `env:"MONGO_TLS_ENABLE" envDefault:"false"`
}

var clientOptions *options.ClientOptions

const writeConcernMajority = "majority"

type DB struct {
	Conn   *mongo.Database
	Client *mongo.Client
}

// NewDB creates connection to mongo and returns the DB struct.
func NewDB(ctx context.Context, cfg Config) (*DB, error) {
	clientOptions = options.Client().ApplyURI(cfg.genConnectURL())

	if cfg.ReplicaSet != "" {
		clientOptions.SetReplicaSet(cfg.ReplicaSet)
	}

	if cfg.WriteConcern != "" {
		var wc *writeconcern.WriteConcern

		if cfg.WriteConcern == writeConcernMajority {
			wc = writeconcern.Majority()
		} else {
			replicasCount, err := strconv.Atoi(cfg.WriteConcern)
			if err != nil {
				return nil, fmt.Errorf("failed to parse replicas count: %w ", err)
			}

			wc = &writeconcern.WriteConcern{W: replicasCount}
		}

		clientOptions.SetWriteConcern(wc)
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connection to mongoDB Error: %w ", err)
	}

	db := &DB{
		Conn:   client.Database(cfg.Database),
		Client: client,
	}

	err = db.Client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ping connection mongoDB Error: %w ", err)
	}

	go db.reconnectOnFailure(ctx)

	return db, nil
}

// reconnectOnFailure implements db reconnection if ping was unsuccessful.
func (db *DB) reconnectOnFailure(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ticker.C:
			err := db.Client.Ping(ctx, nil)
			if err != nil {
				log.Printf("lost connection to mongo: %v", err)
				db.Client, _ = mongo.Connect(ctx, clientOptions)

				err = db.Client.Ping(ctx, nil)
				if err == nil {
					log.Printf("ping to mongo is successful: %v", err == nil)
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			err := db.Client.Disconnect(ctx)
			if err != nil {
				log.Printf("mongo close connection error: %v", err)

				return
			}

			log.Println("mongo connection is closed successfully")
		}
	}
}

func (db *DB) Ping(ctx context.Context) error {
	err := db.Client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("mongo connection error: %w ", err)
	}

	return nil
}
