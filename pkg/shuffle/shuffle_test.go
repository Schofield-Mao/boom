package shuffle

import (
	"testing"
)

func TestHash(t *testing.T){
	testcase := []struct{
		code string
		hashcode int
	}{
		{
			code:"0",
			hashcode:0,
		},
		{
			code:"1",
			hashcode:1,
		},
		{
			code:"2",
			hashcode:2,
		},
		{
			code:"3",
			hashcode:3,
		},
		{
			code:"4",
			hashcode:4,
		},
		{
			code:"5",
			hashcode:5,
		},
		{
			code:"6",
			hashcode:6,
		},
		{
			code:"7",
			hashcode:7,
		},
		{
			code:"8",
			hashcode:8,
		},
		{
			code:"9",
			hashcode:9,
		},
	}

	for _,test := range testcase{
		if test.hashcode != hash(test.code){
			t.Error("error")
		}
	}
}