package sql2struct

import (
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Element struct {
	Key         string `json:"key"`
	BsonType    string `json:"bson_type"`
	Description string `json:"description"`
}

func TestMongoDBModel_Connect(t *testing.T) {
	options.Client().SetAuth(options.Credential{
		AuthSource: "zxw",
		Username:   "zxw",
		Password:   "123456",
	})
	options.Client().SetAppName("lmcli ")
	options.Client().SetConnectTimeout(time.Second)
	options.Client().SetHosts([]string{"localhost:27017"})
	uri := options.Client().GetURI()
	t.Log(uri)
}
func TestRunCommandEval(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	testDB := client.Database("test")
	opts := make([]*options.ListCollectionsOptions, 0)
	opts = append(opts, options.ListCollections().SetNameOnly(true))
	res, err := testDB.ListCollections(ctx, bson.M{"name": "users"})
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer res.Close(ctx)

	for res.Next(ctx) {
		next := &bsonx.Doc{}
		err = res.Decode(next)
		if err != nil {
			t.Error(err.Error())
			return
		}
		elem, err := next.LookupErr("options", "validator", "$jsonSchema")
		if err != nil {
			t.Error(err.Error())
			return
		}
		if elem.Type() != bson.TypeEmbeddedDocument {
			t.Error("类型不对")
			return
		}
		t.Log(elem.String())
		md, ok := elem.MDocumentOK()
		if !ok {
			t.Log("")
			return
		}
		result := SchemaData{}

		result.BsonType = md["bsonType"].StringValue()
		arrayStr := md["required"].Array().String()
		t.Logf(arrayStr)
		array := []Element{}
		for k, v := range md["properties"].MDocument() {
			element := Element{}
			element.Key = k
			//t.Log(v.Document())
			for _, v1 := range v.Document() {
				switch v1.Key {
				case "bsonType":
					element.BsonType = v1.Value.String()
				case "description":
					element.Description = v1.Value.String()

				}
			}
			//element.Description=v.Document()["description"].StringValue()
			//element.BsonType=v.MDocument()["bsonType"].StringValue()
			array = append(array, element)
		}
		t.Log(array)
		//for k,v:=range elem.MDocument() {
		//	t.Logf("文档类型: %s",)
		//}
	}
	res.Close(ctx)
	//client.Database("test").ListCollectionNames(ctx,bson.D{{"name","users"}}) {
	//}
	// t.Log(result.Err().Error())j
	// t.Log(client.Ping(ctx,nil))
}

type SchemaData struct {
	BsonType   string    `json:"bson_type"`
	Required   []string  `json:"required"`
	Properties []Element `json:"properties"`
}

func TestDBModel_Connect(t *testing.T) {
	mm := NewMongoDBModel(&DBInfo{
		//UserName:  "zxw",
		//Password:  "123456",
		Host:      "localhost",
		Port:      27017,
		DBName:    "test",
		TableName: "users",
	})
	err := mm.Connect()
	if err != nil {
		t.Error(err.Error())
	}
	err = mm.DBClient.Ping(context.Background(), nil)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("ping 验证通过")
}
