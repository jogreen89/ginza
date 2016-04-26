package main

import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    for i := 0; i < 10; i++ {
        l.PushBack(i)
    }
    for j := l.Front(); j != nil; j = j.Next() {
        fmt.Println(i.Value)
    }
}
