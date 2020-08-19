# vdshell
small shell for embedded device


![alt text](https://github.com/VincentDrevet/vdshell/blob/master/demo/screenshot.png "screenshot vdshell")

## requirements ##
```
go get github.com/abiosoft/ishell
go get github.com/jaypipes/ghw
go get github.com/shirou/gopsutil
go get github.com/olekukonko/tablewriter
go get github.com/mattn/go-sqlite3
go get golang.org/x/crypto/bcrypt
go get golang.org/x/sys/unix
```

## Build ##

If you want to use this shell in embedded device you should compile this program with static library. For do that just run the build.sh script file.

```
./build.sh
```

You should see vdshell executable file in the folder.

## TODO ##

 authentication
  - auth with sqlite db ?
  - checksum update firmware ?
  - logging ?


## Thanks ##


abiosoft for ishell library

aypipes for ghw library

shirou for gopsutil library

olekukonko for tablewriter library

 mattn go-sqlite3 library
