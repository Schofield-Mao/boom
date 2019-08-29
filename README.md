### boolemer filter
--- 
1）100GB url 文件，使用 1GB 内存计算出出现次数 top100 的 url 和出现的次数
**单机计算**
拆成100份分别读入内存处理后，取出topK,再写入磁盘。
**分布式计算**
-map 
-reduce
100GB分成100份，即每份1GB,每份选出局部topK,然后合并局部topK,得到全局topK

**demo**
X GO url, 分成X/maxM 堆，　存储为临时文件，并分别选出topK
每相邻两个文件A,B　reduce 并写入Ｂ
全部reduce