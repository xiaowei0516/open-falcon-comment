package main

import (
    "os"
    "fmt"
    "path/filepath"
)

func main(){
    wd,err := os.Getwd()
    fmt.Println(wd)
    if err!=nil{
    }

    abs, _ := filepath.Abs("./home/zhang")
    fmt.Println(abs)
    r,err := filepath.Rel(wd, abs)
    if err!=nil{
    }
    fmt.Println(r)
}
