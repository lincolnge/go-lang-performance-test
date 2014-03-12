Go 语言性能测试
=============

目录
-------------
- 测试
  - 1 [engine fetch 与 db Scan 性能比较 (一)](01.md)
  - 2 [测试乘法](02.md)
  - 3 [engine fetch 与 db Scan 性能比较 (二)](03.md)
  - 4 [压力测试](04.md)

分析
-------------
### 分析输出数据  

文本类型如下所示：

	14	2.1%	17.2%	58	8.7%	std::_Rb_tree::find

解释每一列的意思：

- Number of profiling samples in this function
- Percentage of profiling samples in this function
- Percentage of profiling samples in the functions printed so far
- Number of profiling samples in this function and its callees
- Percentage of profiling samples in this function and its callees
- Function name

环境
-------------
操作系统：Mac OS 10.9.2 / Ubuntu 13.10 64bit  
go 版本：go1.2 darwin/amd64 / go1.1.2 linux/amd64

References
-------------
- Google CPU Profiler. <http://google-perftools.googlecode.com/svn/trunk/doc/cpuprofile.html>