package main

import "fmt"
//import "code.google.com/p/go-tour/pic"
import "math"
import "math/rand"
import "code.google.com/p/go-tour/wc"
import "code.google.com/p/go-tour/reader"
import "strings"
import "io"

// type Vertex struct {
//   X, Y int
// }

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v \n", s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
  var ret [][]uint8
  for i := 0; i < dx; i++ {
    var sx []uint8
    for j := 0; j < dy; j++ {
      sx = append(sx, uint8(j%(rand.Intn(256)+1)))
    }
    ret = append(ret, sx)
  }
  return ret
}

// type Vertex struct {
//   Lat, Long float64
// }

var m = map[string]Vertex {
  "Bell Labs": Vertex {
    40.68433, -74.39967,
  },
  "Washton DC": Vertex {
    50.72839, -81.12392,
  },
}

func WordCount(s string) map[string]int {
  tokens := strings.Split(s, " ")
  ret := make(map[string]int)
  for i := 0; i < len(tokens); i++ {
    key := tokens[i]
    elem, ok := ret[key]
    if ok {
      ret[key] = elem + 1
    } else {
      ret[key] = 1
    }
  }
  return ret
}

func fibonacci() func() int {
  first := 0
  second := 0
  return func () int {
    if second == 0 {
      second = 1
      return second
    }
    ret := first + second
    first = second
    second = ret
    return ret
  }
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

type Abser interface {
  Abs() float64
}

type IPAddr [4]byte

func (ip *IPAddr) String() string {
  var s string
  fmt.Println("in string func")
  for i := 0; i < len(ip); i++ {
    s += fmt.Sprintf("%s", ip[i])
    s += "."
  }
  return s
}


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %0.0f", e)
}

func Sqrt(x float64) (float64, error) {
  return -2, ErrNegativeSqrt(-2)
}

func main() {
  // v := Vertex{3,4}
  // v.Scale(5)
  // fmt.Println(v.Abs())

  flt := MyFloat(-3.14)
  fmt.Println(flt.Abs())

  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3,4}

  a = f
  a = &v
  fmt.Println(a.Abs())

  addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}

  //pic.Show(Pic)
  fmt.Println(m["Bell Labs"])
  fmt.Println(m)

  wc.Test(WordCount)

  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  fmt.Println(hypot(3,4))

  _, errmsg := Sqrt(-2)
  fmt.Println(errmsg.Error())
  fmt.Println(ErrNegativeSqrt(-3).Error())

  r := strings.NewReader("Hello, Reader!")

  b := make([]byte, 8)

  for {
    n, err := r.Read(b)
    fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])

    if err == io.EOF {
      break
    }
  }

  // f := fibonacci()
  // for i := 0; i < 10; i++ {
  //   fmt.Println(f())
  // }



  // v := Vertex{1,2}
  // v.X = 4
  // //v.test()
  // fmt.Println("123", v)
  //
  // var a [2]string
  // a[0] = "hello"
  // a[1] = "world"
  // fmt.Println(a);
  //
  // var b [10]int
  // fmt.Println(b)
  //
  // //var p = []int {2,5,3,1,2}
  //
  // c := make([]int, 0, 5)
  // fmt.Println(c)
  // printSlice("c", c)
  //
  // d := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
  // fmt.Printf("%s\n", d)
  // e := d[1:4]
  // e[0] = 't'
  // fmt.Printf("%s\n", d)
  // fmt.Printf("%s\n", e)
  //
  // var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
  // for i, v := range pow {
  //   fmt.Printf("2**%d = %d\n", i, v)
  // }



}
