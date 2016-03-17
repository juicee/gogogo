package main

import "fmt"
import "math/rand"
import "math"
import "math/cmplx"

func add(x int, y int) int {
  return x + y;
}

func swap(x, y string)(string, string) {
  return y, x;
}

var c, python, java bool;
var test string;


var (
  ToBe    bool  = false
  MaxInt  uint64 = 1 << 64 - 1
  z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  var i int;
  var a, b, c string;

  fmt.Printf("gogogo");
  fmt.Println("my favorite number is ", rand.Intn(100));
  fmt.Printf("now you have %g problems. ", math.Pi);
  fmt.Println("adding x with y: ", add(24, 43));
  a, b = swap("this", "that");
  fmt.Println("string swap: ", a, b);
  fmt.Println(i, c, python, java);

  const f = "%T(%v)\n";
  fmt.Println(f, ToBe, ToBe);
  fmt.Println(f, MaxInt, MaxInt);
  fmt.Println(f, z, z);

  v := 42.1 // change me!
  fmt.Printf("v is of type %T\n", v)
  //v = 42.1 // change me
  fmt.Printf("v is of type %T\n", v)

}
