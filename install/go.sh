#!/bin/bash
#
########## Setting Golang environment on CentOS 7 ######################
###
#
VERSION="$2"

download() {
    echo "================= Downloading Golang $VERSION ==================="
    OS="linux"
    ARCH="amd64"
    wget https://storage.googleapis.com/golang/go$VERSION.$OS-$ARCH.tar.gz
}

install() {
    echo "================= Downloading Golang $VERSION ==================="
    OS="linux"
    ARCH="amd64"
    wget https://storage.googleapis.com/golang/go$VERSION.$OS-$ARCH.tar.gz
    echo "================ Installing Golang on CentOS 7 =================="
    tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
}

setup(){
    echo "================ Setting environment variables =================="
    export GOROOT=/usr/local/go
    export GOBIN=$GOROOT/bin
    export PATH=$PATH:$GOBIN
    export GOPATH=$HOME/gopath
    sleep 2
    echo "======================== Done !!! ===============================" 
    sleep 1
    echo "=============== Setting Linux user profile for go ==============="
    echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
    sleep 2
    echo "======================== Done !!! ==============================="
}

helpFile(){
    echo -e "Help File\n=========\n-d   |   --download : To download Golang version only\n-i   |   --install : Download and install Golang by version parameter.\n\t\t Please add parameter like: sh installGo.sh --install 1.12\n-s   |   --setup : Setup environment variables\n-rm  |   --remove : Remove Golang version. Add Version as second param to sh installGo.sh --remove 1.12"
}

remove(){
    echo "====================== Removing Golang =========================="
    rm -rvf /usr/local/go/
}

while [ -n "$1" ]; do

    case "$1" in
    -d | --download) download ;;

    -i | --install) install ;;

    -s | --setup) setup ;;

    -rm | --remove) remove ;;

    -h | --help) helpFile ;;

    *) helpFile ;;

    esac

    shift

done

if [ -z "$1" ]; then
   helpFile
else
   echo "Don't forget the Shipud in the Pita"
fi
