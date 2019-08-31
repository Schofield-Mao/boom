### boom
--- 
boom 是一个基于go语言实现的map-reduce模仿框架。

**如何使用**
- 切割文件接口
```
/**
 *function:split big file into small files, small filename and number of saml files can be specified
 *parameters:
 *@tempfile:the name of temporary files
 *@numPart: the number of files
 *@mainfile: the big file to be splited
 *@mapper: a hook, do something on the byte of small file before writting
 *returns:
 *@[]string:the array of filenames
 *@error:error
 */
func SplitFile(tempfile string, numPart int64, mainfile string, mapper func([]byte)[]byte) ([]string,error)
```

- shuffle接口
```
  /**
 *function:reorganize the files 
 *parameters:
 *@filenames:file list to shuffle
 *returns:
 *@filenames:file list after shuffle
 *@error:error
 */
Shuffler([]string) ([]string,error)
```

- map接口
```
  /**
 *function:An interface to abstract mapping 
 *parameters:
 *@filenames:file list to map
 *@mapper:custom map function 
 *returns:
 *@error:error
 */
func Mapper(filenames []string, mapper func([]byte)[]byte) error
```

- reduce 接口
```
/**
 *function:An interface to reduce result, it reduce every 2 consective result 
 *parameters:
 *@filenames:file list to reduce
 *@reducer:custom function to reduce to
 *returns:
 *@error:error
 */
 func Reducer(filenames []string, reducer func([]byte,[]byte)[]byte) error
```

**例子**

1）100GB url 文件，使用 1GB 内存计算出出现次数 top100 的 url 和出现的次数

**单机计算**

- 基本思路

把原文件拆成ｎ份(100GB/n<1GB),分别读入内存映射处理后，写入ｎ个临时文件。每两个临时文件做归约操作，再写入临时文件，直到只有最后一个临时文件，即最后的topK。

- map
    - 统计url出现个数
    - 计算局部topK

- reduce
    - 合并两个topK,组合出新的哈希表
  


![flow char](/img/flow.jpg)


![phases description](/img/phases.png)

**优化方法：分布式计算**
