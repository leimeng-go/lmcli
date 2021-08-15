package timer

import (
	"testing"
	"time"
)

func TestParse(t *testing.T){
	t.Log(time.Local.String())
	location,_:=time.LoadLocation("Asia/Shanghai")
	inputTime:="2029-09-04 12:02:33"
	layout:="2006-01-02 15:04:05"
	t1,_:=time.ParseInLocation(layout,inputTime,location)
	dataTime:=time.Unix(t1.Unix(),0).In(location).Format(layout)

	t.Logf("输入时间: %s,输出时间: %s",inputTime,dataTime)
}