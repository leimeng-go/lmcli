package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"lmcli/internal/sql2struct"
)

func init() {
	sql2StructCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2StructCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2StructCmd.Flags().StringVarP(&host, "host", "", "", "请输入数据库的Host")
	sql2StructCmd.Flags().IntVarP(&Port,"port","",3306,"请输入数据库的端口")
	sql2StructCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2StructCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2StructCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2StructCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}

var (
	DBType    string
	username  string
	password  string
	host      string
	Port int
	charset   string
	dbType    string
	dbName    string
	tableName string
)

type Database interface{
	Connect()error   
	GetFields(dbName,tableName string)
}

var (
	sql2StructCmd = &cobra.Command{
		Use:   "sql",
		Short: "sql转换和处理",
		Long:  "sql转换和处理，当前仅支持关系型数据库mysql和非关系型数据库mongo",
		Run: func(cmd *cobra.Command, args []string) {
			dbInfo := &sql2struct.DBInfo{
				Host:     host,
				Port: Port,
				UserName: username,
				Password: password,
				Charset:  charset,
				DBName: dbName,
				TableName: tableName,
			}
			var dbm sql2struct.DataBaseModel
			switch dbType {
			case "mysql":
				dbm=sql2struct.NewMysqlDBModel(dbInfo)
			case "mongodb":
				dbm=sql2struct.NewMongoDBModel(dbInfo)
			}
			err :=dbm.Connect()
			if err!=nil{
				log.Fatal(err.Error())
			}
			fields, err:=dbm.GetFields(dbInfo.DBName,dbInfo.TableName)

			tp:=new(sql2struct.StructTemplate)
			tp.Generate(fields)
		},
	}
)

