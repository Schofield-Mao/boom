package kernel

import (
	"strconv"
	"testing"
	"boom/pkg/file"
)

func TestMap2TopK(t *testing.T){
	m := make(map[string]int, 10)
	for i:=0;i<10;i++{
		m[strconv.Itoa(i)] = 1000
	}
	for i:=10;i<100;i++{
		m[strconv.Itoa(i)] = 10
	}

	bs,err := file.Map2Json(m)
	if err != nil{
		t.Error(err)
	}
	newBs := Map2TopK(bs,10)
	newM,err := file.Json2Map(newBs)
	if err != nil{
		t.Error(err)
	}
	for _,v := range newM{
		if v != 1000{
			t.Error("error")
		}
	}
}