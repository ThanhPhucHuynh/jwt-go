package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}
var user = "tphuc"
var password = "74488222"
var urlMongo = "mongodb+srv://"+user+":"+password+"@test.onfty.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"


var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	connectStr := fmt.Sprintf(urlMongo)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {panic(err)}

	fmt.Println("Connect done!")
	Mongo.Client = client
	return Mongo
}