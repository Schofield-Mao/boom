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

func TestWriteFile(t *testing.T){
	testcases := []struct{
		a string
		filename string
	}{
		{
			a : "1 2 3 4 5 6 5 ",
			filename: "hello wolrd",
		},
	}
	for _,test := range testcases{
		WriteFile(test.filename,[]byte(test.a))
	}
}

func TestReadFile(t *testing.T){
	testcases := []struct{
		a string
		filename string
	}{
		{
			a : "1 2 3 4 5 6 5 ",
			filename: "hello wolrd",
		},
	}
	for _,test := range testcases{
		bs,_ := ReadFile(test.filename)
		if string(bs) != test.a{
			t.Error("error")
		}
	}
}

func TestSplitFile(t *testing.T){
	_,err := SplitFile(".temp", 10, "url",func(bs []byte)[]byte{
		return bs
	})
	if err != nil{
		t.Error(err)
	}
}

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

