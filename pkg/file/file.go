package file

import (
	"io/ioutil"
	"os"
	"strconv"
	"encoding/json"
)
func ReadFile(filename string) ([]byte,error) {
  f,err := os.OpenFile(filename, os.O_RDONLY, 0600)
  defer f.Close()
  if err != nil{
		return nil,err
  }
  bs,err := ioutil.ReadAll(f)
  if err != nil{
    return nil,err
  }
  return bs,nil
}

func WriteFile(filename string, bs []byte) (int,error){
	f,err := os.Create(filename)
	defer f.Close()
	if err != nil{
	  return 0,err
	}
	return f.Write(bs)
}

func RemoveFile(filename string) error{
	return os.Remove(filename)
}

func SplitFile(tempfile string, numPart int64, mainfile string, handler func([]byte)([]byte,error)) ([]string,error){
	f,err := os.OpenFile(mainfile, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil{
	  return nil,err
	}
	fileInfo,err := f.Stat()
	if err != nil{
	  return nil,err
	}
	maxM := fileInfo.Size()/numPart+1
	filenames := make([]string,numPart)
	c := 0
	for{
	  bs := make([]byte,maxM)
	  n,err := f.Read(bs)
	  if err != nil{
			return filenames,nil
	  }
		filenames[c] = tempfile+strconv.Itoa(c)
		newBs,err := handler(bs[:n])
	  WriteFile(filenames[c], newBs)
	  c++
	}
}


func Map2Json(m map[string]int)([]byte,error){
	b,err := json.Marshal(m)
	if err != nil{
		return nil,err
	}
	return b,nil
}

func Json2Map(bs []byte) (map[string]int ,error){
	m := make(map[string]int)
	err := json.Unmarshal(bs, &m)
	if err != nil{
		return m,err
	}
	return m,nil
}



