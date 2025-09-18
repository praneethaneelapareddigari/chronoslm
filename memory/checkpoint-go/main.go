package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    filename := fmt.Sprintf("checkpoint-%d.dump", time.Now().Unix())
    os.WriteFile(filename, []byte("redis snapshot with kafka offsets"), 0644)
    fmt.Println("Checkpoint created:", filename)
}
