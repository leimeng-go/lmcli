package mongo2struct

import (
	"context"
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
	result:=client.Database("test").RunCommand(ctx,bson.M{
		"mapreduce": "xw_landlords",
		 
	})
	t.Log(result.Err())
	// t.Log(result.Err().Error())
	// t.Log(client.Ping(ctx,nil))
}