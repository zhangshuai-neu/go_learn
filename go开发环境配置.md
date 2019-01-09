## go 开发环境配置

#### 1.安装vs code

#### 2.安装golang

#### 3.配置环境变量 GOPATH

    ${GOPATH} 指明tools的安装位置
    配置好之后，执行下面的命令

    go get -u -v github.com/ramya-rao-a/go-outline
    go get -u -v github.com/acroca/go-symbols
    go get -u -v github.com/mdempsky/gocode
    go get -u -v github.com/rogpeppe/godef
    go get -u -v golang.org/x/tools/cmd/godoc
    go get -u -v github.com/zmb3/gogetdoc
    go get -u -v golang.org/x/lint/golint
    go get -u -v github.com/fatih/gomodifytags
    go get -u -v golang.org/x/tools/cmd/gorename
    go get -u -v sourcegraph.com/sqs/goreturns
    go get -u -v golang.org/x/tools/cmd/goimports
    go get -u -v github.com/cweill/gotests/...
    go get -u -v golang.org/x/tools/cmd/guru
    go get -u -v github.com/josharian/impl
    go get -u -v github.com/haya14busa/goplay/cmd/goplay
    go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
    go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
    go get -u -v github.com/alecthomas/gometalinter

    将${GOPATH}/bin 加入path环境变量，执行下面命令

    gometalinter --install



#### 4.在vscode中安装 go扩展

    如果vscode提示缺少上述工具，点击 install all即可
    如果安装失败可能是无法翻墙的原因，需要手动安装

