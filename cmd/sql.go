package cmd

import (
	"lmcli/internal/sql2struct"
	"log"

	"github.com/spf13/cobra"
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
		Long:  "sql转换和处理",
		Run: func(cmd *cobra.Command, args []string) {
			dbInfo := &sql2struct.DBInfo{
				DBType:   dbType,
				Host:     host,
				UserName: username,
				Password: password,
				Charset:  charset,
			}
			dbModel := sql2struct.NewDBModel(dbInfo)
			err := dbModel.Connect()
			if err != nil {
				log.Fatalf("dbmodel.Connect err: %v", err.Error())
			}
			columns, err := dbModel.GetColumns(dbName, tableName)
			if err != nil {
				log.Fatalf("dbModel.GetColumns err: %v", err)
			}
			template := sql2struct.NewStructTemplate()
			templateColumns := template.AssemblyColumns(columns)
			err = template.Generate(tableName, templateColumns)
			if err != nil {
				log.Fatalf("template.Generate err: %v", err)
			}
		},
	}
)

