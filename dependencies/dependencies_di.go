package dependencies

import "go.mongodb.org/mongo-driver/mongo"

type DependenciesDI struct {
	mongo *commonMongo
	mysql *mysql
}

func NewDependenciesDIProvider() *DependenciesDI {
	return &DependenciesDI{
		mongo: newCommonMongo(),
		mysql: newMysql(),
	}
}

func (ddi *DependenciesDI) GetMongoCollection(dbName, collectionName string) *mongo.Collection {
	return ddi.mongo.getMongoClient(dbName, collectionName)
}

func (ddi *DependenciesDI) GetMysql() *mysql {
	return ddi.mysql
}
