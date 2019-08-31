# vs code 需要的工具链不全, 所以经将所有需要的代码手动clone了下来

#! /bin/bash

if ! [ -e ./src/github.com ];then
	mkdir ./src/github.com;
fi
cd ./src/github.com;

# git submodule add https://github.com/rogpeppe/godef.git src/github.com/rogpeppe/godef
# git submodule add https://github.com/ramya-rao-a/go-outline.git src/github.com/ramya-rao-a/go-outline
# git submodule add https://github.com/golang/tools.git src/golang/x/tools
# git submodule add https://github.com/sqs/goreturns src/github.com/sqs/goreturns
git submodule add https://github.com/nsf/gocode src/github.com/nsf/gocode


go install src/github.com/rogpeppe/godef/
go install src/github.com/nsf/gocode/

