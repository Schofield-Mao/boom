package kernel

import (
	"fmt"
	"boom/pkg/heap"
	"boom/pkg/file"
)

func Map2TopK(bs []byte, k int) []byte{
	m,err := file.Json2Map(bs)
	if err != nil{
		fmt.Println(err)
	}
	topK := heap.GetTopK(m,k)
	
	newM := make(map[string]int)
	for _,k := range topK{
		newM[k] = m[k]
	}
	newBs,err := file.Map2Json(newM)
	if err != nil{
		fmt.Println(err)
	}
	return newBs
}

func Reduce(a []byte, b []byte) []byte{
	am,err := file.Json2Map(a)
	if err != nil{
		fmt.Println(err)
	}
	bm,err := file.Json2Map(b)
	if err != nil{
		fmt.Println(err)
	}
	cm := make(map[string]int)
	for k,v := range am{
		cm[k]+=v
	}
	for k,v := range bm{
		cm[k]+=v
	}
	c,err := file.Map2Json(cm)
	if err != nil{
		fmt.Println(err)
	}
	return c
}