packages in Go

Important environment variables 


```
export GOROOT=/usr/local/bin
export GOPATH=$HOME/go/workspace
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
```

GOPATH "Workspace" the root of the folder / directory where you're working. 

**src** -- source files
**pkg** -- compiled packages .a files
**bin** -- executable files 

GOBIN
directory where go install and go get will place binaries after building main

GOROOT 

location of Go installation

```
source .bashrc -- to reload the .bashrc file after making any changes. 

go build -- keeps the result in the same directory 

go install -- moves packages to their respective directories 
```