package main

import (
	"fmt"
  "strconv"
  "os"
  "math/rand"
  "strings"
  "boom/pkg/file"
  "boom/pkg/heap"
  "boom/pkg/mapReduce"
)

func main() {
  //spilit file
  filenames,err := file.SplitFile(".temp-",10,"data",Raw2Map)
  if err != nil{
    fmt.Println(err)
  }
  //map
  err = mapReduce.Mapper(filenames, Map2Top10)
  if err != nil{
    fmt.Println(err)
  }
  //reduce
  err = mapReduce.Reducer(filenames, Reduce)
  if err != nil{
    fmt.Println(err)
  }
}

//create data func
func CreateDATA(N int){
  const filename = "data"
  f,err := os.Create(filename)
  defer f.Close()
  if err != nil{
    fmt.Println(err)
  }
  for i:=0; i<N/2 ; i++{
    _,err = f.Write([]byte(strconv.Itoa(rand.Intn(1000))+" "))
  }
  for i:=0; i<N/2 ; i++{
    _,err = f.Write([]byte(strconv.Itoa(rand.Intn(100))+" "))
  }
  fmt.Println("done")
}


//concreate map func
func Raw2Map(bs []byte) []byte{
  a := strings.Split(string(bs), " ")
  m := make(map[string]int)
  for _,k := range a{
	  		m[k]+=1
  }
  bs,err := file.Map2Json(m)
  if err != nil{
    fmt.Println(err)
  }
  return bs
}

func Map2Top10(bs []byte) []byte{
	m,err := file.Json2Map(bs)
	if err != nil{
		fmt.Println(err)
	}
	topK := heap.GetTopK(m,10)
	
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

//concreate reduce func
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
	c = Map2Top10(c)

	if err != nil{
		fmt.Println(err)
	}
	return c
}