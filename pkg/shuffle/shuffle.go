package shuffle

import (
	"strconv"
	"strings"
	"boom/pkg/file"
)

type Shuffle interface{
	Shuffler([]string) ([]string,error)
}

//shuffle&merge
//default implementation
func Shuffler(filenames []string) ([]string,error){
	
	//shuffle
	for _,filename := range filenames{
	  _,err := shuffleHelper(filename)
	  if err != nil{
		  return nil,err
	  }  
	  file.RemoveFile(filename)
	}
  
	//merge
	rt := make([]string,10)
	for i:=0;i<10;i++{
	  temp,err := mergeHelper(i,filenames)
	  if err != nil{
		return nil,err
	  }
	  rt[i] = temp
	}
	return rt,nil
  }

  func shuffleHelper(filename string) ([]string,error){
	mList := make([]map[string]int,10)
	shufflefiles := make([]string,10)
	bs,err := file.ReadFile(filename)
	if err != nil{
	  return nil,err
	}
	m,err :=  file.Json2Map(bs)
	if err != nil{
		return nil,err
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
		return nil,err
	  }
	  shufflefiles[i] = filename+"-shuffle-"+strconv.Itoa(i) 
	  file.WriteFile(shufflefiles[i],b)
	}
	return shufflefiles,nil
  }
  
 func mergeHelper(bucket int, filenames []string) (string,error){
	m := make(map[string]int)
	for _,filename := range filenames{
	  fn := filename+"-shuffle-"+strconv.Itoa(bucket)
	  bs,err := file.ReadFile(fn)
	  defer file.RemoveFile(fn)
	  if err != nil{
		return "",err
	  }
	  temp,err := file.Json2Map(bs)
	  if err != nil{
		return "",err
	  }
	  for k,v := range temp{
		m[k] += v
	  }
	}
	b,err := file.Map2Json(m)
	if err != nil{
		return "",err
	}
	mergefile := ".REDUCE-"+strconv.Itoa(bucket)
	file.WriteFile(mergefile,b)
	return mergefile,nil
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
	
  