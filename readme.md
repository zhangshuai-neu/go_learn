#### Go语言的目录树
go命令依赖一个重要的环境变量：**$GOPATH**  
$GOPATH 下有3个目录 ,src, pkg, bin

- src: 存放源代码(比如：.go .c .h .s等)   
- pkg: 编译时生成的中间文件（比如：.a）
- bin: 编译后生成的可执行文件

#### 新项目

在src目录下，创建一个新的目录，用这个目录来作为一个新的项目。

如果package main存在的文件，则会在pkg目录下生成中间文件，
在bin目录下生成相应的可执行文件。

如果package main不存在的文件，将只会在pkg目录下生成中间文件。

#### 开始
    
执行下列命令，编译基础的golang的工具，主要是vscode自己下载总是失败 :(

    sudo apt-get install golang
    配置GOPATH环境变量
    git config --global http.postBuffer 2000000000
    git submodule init
    git submodule update
    chmod +x tools_install.sh
    ./tools_install.sh

