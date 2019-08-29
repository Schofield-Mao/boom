package main

import (
	"fmt"
  "strconv"
  "os"
  "math/rand"
  "boom/pkg/file"
  "boom/pkg/kernel"
  "strings"
)
const filename = "url"
func main() {
  fn,err := file.PartFile(".temp",10,"url",ByteHandler)
  if err != nil{
    fmt.Println(err)
  }
  fmt.Println(fn)
  for _,filename := range fn{
    bs,err := file.ReadFile(filename)
    if err != nil{
      fmt.Println(err)
    }
    newBs := kernel.Map2TopK(bs,10)
    _,err = file.WriteFile(filename,newBs)
    if err != nil{
      fmt.Println(err)
    }
  }

  for i,filename := range fn{
    if i == len(fn)-1{
      break
    }
    a,err := file.ReadFile(fn[i])
    if err != nil{
      fmt.Println(err)
    }
    b,err := file.ReadFile(fn[i+1])
    if err != nil{
      fmt.Println(err)
    }
    newBs := kernel.Reduce(a,b)
    newBs = kernel.Map2TopK(newBs,10)
    _,err = file.WriteFile(fn[i+1],newBs)
    if err != nil{
      fmt.Println(err)
    }
    err = file.RemoveFile(fn[i])
    fmt.Println(filename)
    if err != nil{
      fmt.Println(err)
    }
  }
}
func CreateURL(N int){
  f,err := os.Create(filename)
  defer f.Close()
  if err != nil{
    fmt.Println(err)
  }
  for i:=0; i<N ; i++{
    _,err = f.Write([]byte(strconv.Itoa(rand.Intn(1000))+" "))
  }
  fmt.Println("done")
}

func ByteHandler(bs []byte) ([]byte,error){
	a := strings.Split(string(bs), " ")
  m := make(map[string]int)
	for _,k := range a{
		m[k]+=1
	}
  bs,err := file.Map2Json(m)
  if err != nil{
    return nil,err
  }
	return bs,nil 
}
