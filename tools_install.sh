#! /bin/bash

if ! [ -e ./src/github.com ];then
	mkdir ./src/github.com;
fi
cd ./src/github.com;

git submodule add https://github.com/rogpeppe/godef.git src/github.com/godef
