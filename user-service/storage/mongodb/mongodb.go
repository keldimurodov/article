package mongodb

import (
	"context"
	"log"
	u "projects/article/user-service/genproto/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type mongoRepo struct {
	db *mongo.Client
}

// NewMongoRepo ...
func NewUserRepo(db *mongo.Client) *mongoRepo {
	return &mongoRepo{db: db}
}

var client *mongo.Client

func (r *mongoRepo) Create( user *u.User) (*u.User, error) {
	collection := client.Database("test").Collection("mongodb")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mongoRepo) Update(user *u.User) (*u.User, error) {
	collection := client.Database("test").Collection("mongodb")
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mongoRepo) Get(userID *u.User) (*u.User, error) {
	var user *u.User
	collection := client.Database("test").Collection("mongodb")
	filter := bson.M{"_id": userID.Id}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func (r *mongoRepo) GetAll(user *u.GetAllRequest) (*u.GetAllResponse, error) {
	var allUsers u.GetAllResponse
	collection := client.Database("test").Collection("mongodb")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user *u.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		allUsers.Users = append(allUsers.Users, user)
	}
	return &allUsers, nil
}

func (r *mongoRepo) Delete(userID *u.GetUserRequest) (*u.User, error) {
    var deletedUser u.User
    collection := client.Database("test").Collection("mongodb")
    filter := bson.M{"_id": userID.UserId}

    // O'chirilayotgan foydalanuvchini olish
    err := collection.FindOneAndDelete(context.Background(), filter).Decode(&deletedUser)
    if err != nil {
		log.Fatal("Error deleting user")
    }
    return &deletedUser, nil
}
