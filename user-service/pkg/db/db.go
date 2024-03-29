package db

import (
	"fmt"
	"projects/article/user-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDb, nil
}


func ConnectMongoDB(cfg config.Config) (*mongo.Client, error){

	// MongoDB-ga bog'lanish
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// MongoDB-ga ulanishni sinab ko'rish
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("MongoDB-ga ulanish muammo:", err)
	}

	fmt.Println("MongoDB-ga muvaffaqiyatli ulanildi.")	

	return client, nil
}


func ConnectDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, nil
	}

	cleanUpFunc := func() {
		connDb.Close()
	}

	return connDb, cleanUpFunc
}
