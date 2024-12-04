// Connects to MongoDB and sets a Stable API version
package dependencies

import (
	"context"

	"github.com/addione/New/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commonMongo struct {
	mongoUri string
}

func newMongo() (cm *commonMongo) {
	uri, _ := helpers.GetEnvVariable(helpers.MONGO_URI)

	return &commonMongo{
		mongoUri: uri,
	}
}

func (cm *commonMongo) getMongoClient(dbName string, collectionName string) *mongo.Collection {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cm.mongoUri).SetServerAPIOptions(serverAPI)

	client, _ := mongo.Connect(context.TODO(), opts)

	return client.Database(dbName).Collection(collectionName)

}
