/*
	This test file can pass only with {data} file under this path
	with fully tested, this file exists only for usage
*/
package file

import (
	"testing"
	"encoding/json"
	"strconv"
)

func TestReadFile2Map(t *testing.T){
	m := make(map[string]int)
	for i:=0;i<100;i++{
		m[strconv.Itoa(i)] = i;
	}
	b,err := json.Marshal(m)
	a,err := Json2Map(b)
	if err != nil{
		t.Error(err)
	}
	for i:=0;i<100;i++{
		if a[strconv.Itoa(i)] != i{
			t.Error("error")
		}
	}
}

