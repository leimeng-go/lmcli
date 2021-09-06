package mongo2struct

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBModel struct{
	DBClient *mongo.Client
    DBInfo DBInfo
}

type DBInfo struct{
	User string
	Password string
	Host string
	Port int 
}

type CollectionField struct{
	ColumnKey string
	DataType string
	Occurrences int 
	Percents float32
}

func NewDBModel(info *DBInfo)*DBModel{
	return &DBModel{DBInfo:*info}
}
func (m *DBModel)Connect()error{
	var err error
	url:=fmt.Sprintf("mongodb://%s:%s@%s:%d",m.DBInfo.User,m.DBInfo.Password,m.DBInfo.Host,m.DBInfo.Port)
	ctx,cancel:=context.WithTimeout(context.Background(),20*time.Second)
	defer cancel()
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(url))
	if err!=nil{
		return err
	}
	m.DBClient= client
	return nil
}
func (m *DBModel)GetFields(dbName,collectionName string)([]*CollectionField, error) {
	m.DBClient.Database(dbName).RunCommand(nil,bson.M{"eval":""})
	return nil,nil
}