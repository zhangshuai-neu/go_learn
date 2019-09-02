# 学习go进行开发

#### Go语言开发环境

开发环境使用的linux，随便搞个发行版吧 : )

先安装golang的开发包，然后加压即可，会出现 **go** 目录

    wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz

#### 配置golang的环境变量

export GOROOT= **go**目录的绝对路径

export GOPATH= **go_learn**目录的绝对路径

export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

$GOPATH 下有3个目录 ,src, pkg, bin

- src: 存放源代码(比如：.go .c .h .s等)   
- pkg: 编译时生成的中间文件（比如：.a）
- bin: 编译后生成的可执行文件

#### vim开发golang的依赖
    
执行下列命令，编译基础的golang的工具，主要是vscode/vim的自动下载总是失败 : ( ，所以我在用vim并手动解决常用了工具。

    git config --global http.postBuffer 2000000000
    git submodule init
    git submodule update
    chmod +x tools_install.sh
    ./tools_install.sh

#### vim环境搭建

安装 vundle.vim

    mkdir ~/.vim/bundle
    git clone https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim 

修改 ~/.vimrc文件，我用的下面的配置

    set number                    " 行数
    set ts=4                      " 修改tab的缩进
    set nocompatible              " be iMproved, required
    filetype off                  " required


    " 使用vundle配置vim插件
    " set the runtime path to include Vundle and initialize
    set rtp+=~/.vim/bundle/Vundle.vim
    call vundle#begin()

    " let Vundle manage Vundle, required
    " 使用Vundle管理插件 必要的
    Plugin 'gmarik/Vundle.vim'

    " golang的插件 
    Plugin 'fatih/vim-go'

    " you complete me 插件
    Plugin 'Valloric/YouCompleteMe'

    " All of your Plugins must be added before the following line
    call vundle#end()            " required
    filetype plugin indent on    " required

    " YCM settings
    " 让YCM通过回车和向下的箭头来做list item正向选择，通过向上箭头做反向选择。通过ctrl+space来原地触发补齐提示。
    let g:ycm_key_list_select_completion = ['', '']
    let g:ycm_key_list_previous_completion = ['']
    let g:ycm_key_invoke_completion = '<C-Space>'

然后, 执行在vim下执行 PluginInstall 命令，安装 vim-go和YouCompleteMe。

vim-go需要go-tools已经在上面的go依赖中处理了。

YouCompleteMe的包比较大，可以使用命令 du -h .vim/bundle |tail -1 ， 查看下是否下载完成。 下载完成后，会提示需要手动编译编译YCM的support库。
执行下面的命令即可。

    sudo apt-get install build-essential 
    sudo apt-get install python-dev cmake
    # 失败的话执行下  sudo apt-get update --fix-missing
    cd ~/.vim/bundle/YouCompleteMe
    ./install.sh