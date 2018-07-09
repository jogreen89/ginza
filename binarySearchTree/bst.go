package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func printFile(path string, info os.FileInfo, err error) error {
    // Prints FILE that is found in current directory
    // and all sub-directories
    if err != nil {
        log.Print(err)
        return nil
    }

    fmt.Println(path)

    return nil
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1:]

    // if dir == nil {
    //    log.Fatal("Did not provide a directory path to explore");
    // }

    err := filepath.Walk(dir, printFile)

    if err != nil {
        log.Fatal(err)
    }
}
