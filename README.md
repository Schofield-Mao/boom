### boom
--- 
boom 是一个基于go语言实现的map-reduce框架。


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
    - 再次计算topK
  

**优化：分布式计算**
-map 
-reduce
100GB分成100份，即每份1GB,每份选出局部topK,然后合并局部topK,得到全局topK

**demo**
X GO url, 分成X/maxM 堆，　存储为临时文件，并分别选出topK
每相邻两个文件A,B　reduce 并写入Ｂ
全部reduce