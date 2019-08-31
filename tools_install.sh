#! /bin/bash

if [ !(-e ./src/github.com) ]
    mkdir ./src/github.com
cd ./src/github.com

git clone https://github.com/rogpeppe/godef.git