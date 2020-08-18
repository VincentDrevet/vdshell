#!/bin/bash

# Compilation ARM
#CGO_ENABLED=1 CC=/home/vdrevet/Perso/embedded/br-tree/board/rpi-3/cross/bin/arm-linux-gcc GOOS=linux GOARCH=arm go build -ldflags "-linkmode external -extldflags -static" -a vdshell.go show.go authentification.go shell.go display.go fs.go alimentation.go services.go bash.go

# Compilation x86_64
go build -ldflags "-linkmode external -extldflags -static" -a vdshell.go show.go authentification.go shell.go display.go fs.go alimentation.go services.go bash.go sqlmanage.go
