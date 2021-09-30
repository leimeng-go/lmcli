package sql2struct

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
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
//CollectionField mongo字段映射
type CollectionField struct{
	ColumnKey string //字段名称
	ColumnType string //字段类型
	Description string  //字段说明
}
type CollectionStruct struct{
	CollectionName string  
    List []*CollectionField 
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
func (m *DBModel)GetFields(dbName,collectionName string)(*CollectionStruct, error) {
	ctx:=context.Background()
	dataBase:=m.DBClient.Database(dbName)
	opts:=make([]*options.ListCollectionsOptions,0)
	opts = append(opts, options.ListCollections().SetNameOnly(true))
	res,err:=dataBase.ListCollections(ctx,bson.M{"name":collectionName})
	if err!=nil{
		return nil,err
	}
	defer res.Close(ctx)
    list:=make([]*CollectionField,0)
	for res.Next(ctx){
		next:=&bsonx.Doc{}
		err=res.Decode(next)
		if err!=nil{
			return nil,err
		}
		elem,err:=next.LookupErr("options","validator","$jsonSchema","properties")
		if err!=nil{
			return nil,err
		}
	md,ok:=elem.MDocumentOK()
	if !ok{
		return nil,errors.New("properties类型错误")
	}
	
    for k,v:=range md{
        element:=&CollectionField{}
		element.ColumnKey=k
		for _,value:=range v.Document(){
			switch value.Key{
			case "bsonType":
				element.ColumnType=value.Value.String()
			case "description":
				element.Description=value.Value.String()
			}
		}
		list=append(list, element)
	}
 } 

	return &CollectionStruct{
		CollectionName: collectionName,
		List: list,
	},nil
}