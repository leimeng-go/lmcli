package mongo2struct

import (
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestRunCommandEval(t *testing.T){
	ctx:=context.Background()
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		t.Error(err.Error())
		return
	}
   client.Database("test")
	testDB:=client.Database("test")
	//result:=client.Database("test").RunCommand(ctx,bson.M{
	//	"eval":"print(hello)",
	//})
	//t.Log(result.Err())
	opts:=make([]*options.ListCollectionsOptions, 0)
	opts=append(opts,options.ListCollections().SetNameOnly(true))

	res,err:=testDB.ListCollections(ctx,bson.M{"name":"users"})
	if err!=nil{
		t.Error(err.Error())
		return
	}
	defer res.Close(ctx)

	for res.Next(ctx){
		next:=&bsonx.Doc{}
		err=res.Decode(next)
		if err!=nil{
			t.Error(err.Error())
			return
		}
		elem,err:=next.LookupErr("va")
	}

	client.Database("test").ListCollections(ctx,bson.D{{"name","users"}}) {
	}
	// t.Log(result.Err().Error())j
	// t.Log(client.Ping(ctx,nil))
}
