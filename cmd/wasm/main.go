package main

import (  
  "fmt"
  "os"

  "homepage/internal/wasm"
)

func main() {  
    if err := wasm.Run(); err != nil {
      fmt.Fprintf(os.Stderr, "%s\n", err)
      os.Exit(1)
    }
}