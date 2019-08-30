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

func main() {
  //spilit file
  fn,err := file.SplitFile(".temp-",10,"data",ByteHandler)
  if err != nil{
    fmt.Println(err)
  }

  //map:countting&get local top K
  for _,filename := range fn{
    //get byte of {filename}
    bs,err := file.ReadFile(filename)
    if err != nil{
      fmt.Println(err)
    }
    
    //map
    newBs := kernel.Map2TopK(bs,10)
    
    //write back
    _,err = file.WriteFile(filename,newBs)
    if err != nil{
      fmt.Println(err)
    }
  }


  //reduce:reduce into 1 file
  for i,_ := range fn{
    //reduce consective two file
    if i == len(fn)-1{
      break
    }
    //get byte of file a
    a,err := file.ReadFile(fn[i])
    if err != nil{
      fmt.Println(err)
    }
    //get byte of file b
    b,err := file.ReadFile(fn[i+1])
    if err != nil{
      fmt.Println(err)
    }
    //reduce
    newBs := kernel.Reduce(a,b)
    //map
    newBs = kernel.Map2TopK(newBs,10)
    //write back to b file
    _,err = file.WriteFile(fn[i+1],newBs)
    if err != nil{
      fmt.Println(err)
    }
    //remove file a
    err = file.RemoveFile(fn[i])
    if err != nil{
      fmt.Println(err)
    }
  }
}
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
