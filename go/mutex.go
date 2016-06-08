package main

import (
    "hash/fnv"
    "log"
    "math/rand"
    "os"
    "sync"
    "time"
)

const hunger = 3
const think  = time.Second / 100
const eat    = time.Second / 100

var ph = []string {
    "Aristotle", 
    "Kant", 
    "Spinoza", 
    "Marx", 
    "Russell",
}
var fmt = log.New(os.Stdout, "", 0)
var dining sync.WaitGroup

func philosopher(philoName string, dominantHand, otherHand *sync.Mutex) {
    fmt.Println(philoName, "seated")
    
    h := fnv.New64a()
    h.Write([]byte(philoName))

    rg := rand.New(rand.NewSource(int64(h.Sum64())))

    rSleep := func(t time.Duration) {
        time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
    }

    for h:= hunger; h > 0; h-- {
        fmt.Println(philoName, "hungry")

        dominantHand.Lock()   // pick up forks
        otherHand.Lock()

        fmt.Println(philoName, "eating")
        rSleep(eat)
        dominantHand.Unlock() // put down forks
        otherHand.Unlock()
        fmt.Println(philoName, "thinking") 
        rSleep(think)
    }

    fmt.Println(philoName, "satisfied")
    dining.Done()
    fmt.Println(philoName, "left the table")
}

func main() {
    fmt.Println("table empty")
    dining.Add(5)
    fork0 := &sync.Mutex{}
    forkLeft := fork0
    for i := 1; i < len(ph); i++ {
        forkRight := &sync.Mutex{}
        go philosopher(ph[i], forkLeft, forkRight)
        forkLeft = forkRight
    }
    go philosopher(ph[0], fork0, forkLeft)
    dining.Wait()            // wait for philosophers to finish
    fmt.Println("table empty")
}
