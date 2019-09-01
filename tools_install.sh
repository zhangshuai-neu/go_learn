# vs code 需要的工具链不全, 所以经将所有需要的代码手动clone了下来

#! /bin/bash

# 添加子目录
# if ! [ -e ./src/github.com ];then
# 	mkdir ./src/github.com;
# fi
# cd ./src/github.com;

# git submodule add https://github.com/rogpeppe/godef.git src/github.com/rogpeppe/godef
# git submodule add https://github.com/ramya-rao-a/go-outline.git src/github.com/ramya-rao-a/go-outline
# git submodule add https://github.com/sqs/goreturns src/github.com/sqs/goreturns
# git submodule add https://github.com/nsf/gocode src/github.com/nsf/gocode
# git submodule add https://github.com/golang/tools.git src/golang.org/x/tools
# git submodule add https://github.com/kisielk/errcheck.git src/github.com/kisielk/errcheck
# git submodule add https://github.com/golang/net.git src/golang.org/x/net
# git submodule add https://github.com/golang/xerrors.git src/golang.org/x/xerrors
# git submodule add https://github.com/golang/sync.git src/golang.org/x/sync



go install github.com/rogpeppe/godef
go install github.com/ramya-rao-a/go-outline
go install github.com/nsf/gocode
go install github.com/sqs/goreturns
go install github.com/kisielk/errcheck

go install golang.org/x/tools/cmd/goimports
go install golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/gopls
go install golang.org/x/tools/cmd/gorename
go install golang.org/x/tools/cmd/gotype
