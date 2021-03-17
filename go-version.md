# go version 高级用法

本文参考[go version命令的高级用法](https://www.flysnow.org/2020/08/30/golang-version.html)，主要是想测试下文章能否通过issue label来管理

* go version go-torch  
  打印go-torch这个可执行文件的构建版本信息，需要传递绝对路径，否则默认在当前路径下查找，没找到会打印如"stat go-torch: no such file or directory"的错误。
* go version $GOPATH/bin  
  传递文件夹路径，则会打印给定文件夹路径下的所有golang执行文件的构建版本信息。
* go version -m go-torch  
  查看golang可执行文件的mod信息
