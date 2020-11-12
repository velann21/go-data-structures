package main

import (
	"fmt"
	"time"
)

func main() {
fmt.Println(time.Now().UTC().Format(time.RFC3339Nano))
fmt.Println(time.Now().Add(10*time.Minute).UTC().Format(time.RFC3339Nano))
}
