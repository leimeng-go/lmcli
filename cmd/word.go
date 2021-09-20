package cmd

import (
	"github.com/spf13/cobra"
	"lmcli/internal/word"
	"log"
	"strings"
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

const (
	MODE_UPPER                         = iota + 1 //全部单词转为大写
	MODE_LOWER                                    //全部单词转为小写
	MODE_UNDERSOCRE_TO_UPPER_CAMELCASE            //下划线单词转为大写驼峰单词
	MODE_UNDERSOCRE_TO_LOWER_CAMELCASE            //下划线单词转为小写驼峰单词
	MODE_CAMELCASE_TO_UNDERSCORE                  //驼峰单词转为下划线单词
)

var (
	str  string
	mode int8
)

//命令描述
var (
	desc = strings.Join([]string{
		"1: 全部单词转为大写",
		"2: 全部单词转为小写",
		"3: 下划线单词转为大写驼峰单词",
		"4: 下划线单词转为小写驼峰单词",
		"5: 驼峰单词转为下划线单词",
	}, "\n")
)
var (
	wordCmd = &cobra.Command{
		Use:   "word",
		Short: "单词格式转换",
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			var content string
			switch mode {
			case MODE_UPPER:
				content = word.ToUpper(str)
			case MODE_LOWER:
				content = word.ToLower(str)
			case MODE_UNDERSOCRE_TO_UPPER_CAMELCASE:
				content = word.UnderscoreToUpperCamelCase(str)
			case MODE_UNDERSOCRE_TO_LOWER_CAMELCASE:
				content = word.UnderscoreToLowerCamelCase(str)
			case MODE_CAMELCASE_TO_UNDERSCORE:
				content = word.CamelCaseToUnderscore(str)
			default:
				log.Fatalf("暂不支持改转换模式，请执行help word 查看帮助文档")
			}
			log.Printf("输出结果: %s", content)
		},
	}
)
