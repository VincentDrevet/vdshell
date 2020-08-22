#!/bin/bash

Buildroot_Path="/home/vincent/br-tree/board/rpi-3/ovl/bin"

if [ -z "$1" ]
then
    echo "no ARCH provided exit"
    exit 1
fi

if [ $1 = "x86_64" ]
then    
    # Compilation x86_64
    go build -ldflags "-linkmode external -extldflags -static" -a vdshell.go show.go authentification.go shell.go display.go fs.go alimentation.go services.go bash.go sqlmanage.go user.go pwmanage.go
fi

if [ $1 = "arm" ]
then
    CGO_ENABLED=1 CC=/home/vincent/br-tree/board/rpi-3/cross/bin/arm-linux-gcc GOOS=linux GOARCH=arm go build -ldflags "-linkmode external -extldflags -static" -a vdshell.go show.go authentification.go shell.go display.go fs.go alimentation.go services.go bash.go sqlmanage.go
    mv ./vdshell $Buildroot_Path
fi

