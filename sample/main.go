package main

import (
    "container/list"
    "fmt"
)

func main() {
    // Create a new list and put some numbers in it.
    // Initialization.
    l    := list.New()
    list := list.New()

    // Insertion.
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushBack(4)

    // Iterate through list and print its contents.
    for i := l.Front(); i != nil; i = i.Next() {
        fmt.Println(i.Value)
    }
}
