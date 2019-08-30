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

  shuffle(filenames)

  filenames = Reducer(filenames)
  //map
  err = mapReduce.Mapper(filenames, Map2Top10)
  if err != nil{
    fmt.Println(err)
  }
  // //reduce
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

//
func shuffle(filenames []string){
  for _,filename := range filenames{
    shuffleHelper(filename)
    file.RemoveFile(filename)
  }
}

func shuffleHelper(filename string){
  mList := make([]map[string]int,10)
  
  bs,err := file.ReadFile(filename)
  if err != nil{
    fmt.Println(err)
  }
  m,err :=  file.Json2Map(bs)
  if err != nil{
    fmt.Println(err)
  }
  for k,v := range m{
    if mList[hash(k)] == nil{
      mList[hash(k)] = make(map[string]int)
    } 
    mList[hash(k)][k] = v 
  }
  for i:=0;i<10;i++{
    b,err := file.Map2Json(mList[i])
    if err != nil{
      fmt.Println(err)
    }
    file.WriteFile(filename+"-shuffle-"+strconv.Itoa(i),b)
  }

}

func Reducer(filenames []string) []string{
  rt := make([]string,10)
  for i:=0;i<10;i++{
    rt[i] = reduceHelper(i,filenames)
  }
  return rt
}

func reduceHelper(bucket int, filenames []string) string{
  m := make(map[string]int)
  for _,filename := range filenames{
    fn := filename+"-shuffle-"+strconv.Itoa(bucket)
    bs,err := file.ReadFile(fn)
    defer file.RemoveFile(fn)
    if err != nil{
      fmt.Println(err)
    }
    temp,err := file.Json2Map(bs)
    if err != nil{
      fmt.Println(err)
    }
    for k,v := range temp{
      m[k] += v
    }
  }
  b,err := file.Map2Json(m)
  if err != nil{
    fmt.Println(err)
  }
  file.WriteFile(".REDUCE-"+strconv.Itoa(bucket),b)
  return ".REDUCE-"+strconv.Itoa(bucket)
}

func hash(key string) int{
  if strings.HasPrefix(key, "0"){
    return 0;
  }
  if strings.HasPrefix(key, "1"){
    return 1;
  }
  if strings.HasPrefix(key, "2"){
    return 2;
  }
  if strings.HasPrefix(key, "3"){
    return 3;
  }
  if strings.HasPrefix(key, "4"){
    return 4;
  }
  if strings.HasPrefix(key, "5"){
    return 5;
  }
  if strings.HasPrefix(key, "6"){
    return 6;
  }
  if strings.HasPrefix(key, "7"){
    return 7;
  }
  if strings.HasPrefix(key, "8"){
    return 8;
  }
  if strings.HasPrefix(key, "9"){
    return 9;
  }
  return 0
}
  
