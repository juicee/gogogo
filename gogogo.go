package main

import "fmt"
import "math/rand"
import "math"

func add(x int, y int) int {
  return x + y;
}

func main() {
  fmt.Printf("gogogo");
  fmt.Println("my favorite number is ", rand.Intn(100));
  fmt.Printf("now you have %g problems. ", math.Pi);
  fmt.Println("adding x with y: ", add(24, 43));
}
