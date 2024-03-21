package main

import (
	"net"
	"projects/article/user-service/config"
	pbu "projects/article/user-service/genproto/user"
	"projects/article/user-service/pkg/db"
	"projects/article/user-service/pkg/logger"
	"projects/article/user-service/service"

	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"log"
	"fmt"
)

func main() {


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

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService(connDB,  log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbu.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
