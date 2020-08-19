package db

import (
	"context"
	"fmt"
	"reflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/seheee/PDK/logic-core/domain/model"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type logicRepository struct {
	client *mongo.Client
	collection *mongo.Collection
}

func NewLogicRepository() *logicRepository {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	cli, err := mongo.Connect(context.TODO(), clientOptions)
	fmt.Println("\nresult type:", reflect.TypeOf(cli))
	if err != nil {
		fmt.Println("connect error: ", err.Error())
	}

	// Check the connection
	err = cli.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("connect error: ", err.Error())
	}

	fmt.Println("Connected to MongoDB!")
	return &logicRepository{
		client: cli,
		collection: cli.Database("test").Collection("logics"),
	}
}

func (lr *logicRepository) GetAll() (r []model.Ring, err error) {
	cur, err := lr.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("find error:", err.Error())
	}
	
	for cur.Next(context.TODO()) {
		var elem model.Ring
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("decode error: ", err.Error())
		}
		r = append(r, elem)
	}

	return r, err
}

func (lr *logicRepository) Create(r *model.RingRequest) (string, error) {
	result, err := lr.collection.InsertOne(context.TODO(),r)
	if err != nil {
		fmt.Println("insert error:", err.Error())
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
} 

func (lr *logicRepository)Delete(id string) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	} 
	_, err = lr.collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		fmt.Println("delete error:", err.Error())
		return err
	}
	return nil
}