// Connects to MongoDB and sets a Stable API version
package dependencies

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the placeholder with your Atlas connection string
const uri = "mongodb://root:secret@localhost:27017/?timeoutMS=5000"

type CommonMongo struct {
}

func NewCommonMongo() (cm *CommonMongo) {
	return &CommonMongo{}
}

func (cm *CommonMongo) ConnectMongo() {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{primitive.E{Key: "bson", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func (cm *CommonMongo) GetMongoClient(dbName string, collectionName string) *mongo.Collection {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, _ := mongo.Connect(context.TODO(), opts)

	return client.Database(dbName).Collection(collectionName)

}
