package mapReduce

import (
	"boom/pkg/file"
)

type MapReduce interface{
	Mapper([]byte) []byte
	Reducer([]byte,[]byte) []byte
}


/**
 *function:An interface to abstract mapping 
 *parameters:
 *@filenames:file list to map
 *@mapper:custom map function 
 *returns:
 *@error:error
 */
func Mapper(filenames []string, mapper func([]byte)[]byte) error {
	//map:countting&get local top K
	for _,filename := range filenames{
	 //get byte of {filename}
	 bs,err := file.ReadFile(filename)
	 if err != nil{
	   return err
	 }
	 
	 //map
	 newBs := mapper(bs)
	 
	 //write back
	 _,err = file.WriteFile(filename,newBs)
	 if err != nil{
		return err
	}
   }
   return nil
 }
 
 /**
 *function:An interface to reduce result, it reduce every 2 consective result 
 *parameters:
 *@filenames:file list to reduce
 *@reducer:custom function to reduce to
 *returns:
 *@error:error
 */
 func Reducer(filenames []string, reducer func([]byte,[]byte)[]byte) error{
  //reduce:reduce into 1 file
  for i,_ := range filenames{
   //reduce consective two file
   if i == len(filenames)-1{
	 break
   }
   //get byte of file a
   a,err := file.ReadFile(filenames[i])
   if err != nil{
	return err
   }
   //get byte of file b
   b,err := file.ReadFile(filenames[i+1])
   if err != nil{
	return err
   }
   //reduce
   newBs := reducer(a,b)
   //map
 
   //write back to b file
   _,err = file.WriteFile(filenames[i+1],newBs)
   if err != nil{
	return err
   }
   //remove file a
   err = file.RemoveFile(filenames[i])
   if err != nil{
	return err
   }
  }
  return nil
 }