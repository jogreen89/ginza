// bytecounter.go
//
// A Go example of interfaces, from Ch. 7 D&K (2016).
// 2016 (c) zubernetes
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p)) // convert int to ByteCounter
    return len(p), nil
}

func main() {
    var c ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c) // "5", = len("hello")

    c = 0 // reset the counter
    var name = "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c) // "12", = len("hello, Dolly")
}