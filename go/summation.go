package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    str := s.Text()
    
    x, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error!")
    }

    s.Scan()
    str = s.Text()
    z := strings.Split(str, " ")
    sum := 0
    for i := 0; i < x; i++ {
        a, err := strconv.Atoi(z[i])
        if err != nil {
            fmt.Println("Error")
        }
        sum = sum + a
    }
    fmt.Println(sum)
}
