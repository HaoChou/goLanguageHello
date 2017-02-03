"# golanguageHello" 

先安装 https://golang.org/doc/install?download=go1.7.5.windows-amd64.msi

 set GOPATH D:\goworkspace 设置gopath
 
 在%GOPATH%内建立文件夹 src/hello  在src/hello内建一个go文件 hello.go 内容为
 package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}

执行命令进行编译
go install hello

成功后 发现 %GOPATH%内除了刚才建立的src文件夹外多了一个bin文件夹 bin文件夹内部有个hello.exe
可以通过 %GOPATH%\bin\hello 执行命令 
打印出 hello, world


ps:
 $GOPATH 目录约定有三个子目录：
1.src 存放源代码（比如：.go .c .h .s等）
2.pkg 编译后生成的文件（比如：.a）
3.bin 编译后生成的可执行文件 （比如：.exe）


**************************************************
也可以使用命令直接运行go文件
比如 go run ./hello.go