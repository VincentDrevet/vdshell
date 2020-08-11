#!/bin/bash


go build -ldflags "-linkmode external -extldflags -static" -a vdshell.go show.go authentification.go shell.go display.go fs.go
