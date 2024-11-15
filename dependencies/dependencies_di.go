package dependencies

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type DependenciesDI struct {
	mongo *commonMongo
	mysql *commonMysql
}

func NewDependenciesDIProvider() *DependenciesDI {
	return &DependenciesDI{
		mongo: newMongo(),
		mysql: newMysql(),
	}
}

func (ddi *DependenciesDI) GetMongoCollection(dbName, collectionName string) *mongo.Collection {
	return ddi.mongo.getMongoClient(dbName, collectionName)
}

func (ddi *DependenciesDI) GetMysql(dbName string) *sql.DB {
	return ddi.mysql.getMysqlClient(dbName)
}
