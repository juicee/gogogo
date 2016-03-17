package main

import "fmt"
import "math"
import "runtime"
import "time"

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func Sqrt(x float64) float64 {
  return 0
}

func deferedfun() (i int) {
  defer func() { i++ } ()
  return 1
}

func main() {
  sum := 0
  for i:=0; i<10; i++ {
    sum+=i
  }

  fmt.Print("Go runs on ")

  switch os:=runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    fmt.Printf("%s", os)
  }

  fmt.Println("when's Saturday?")

  today := time.Now().Weekday()

  fmt.Println("weekday: ", int(today))

  switch time.Saturday {
  case today + 0:
    fmt.Println("Today")
  case today + 1:
    fmt.Println("Tomorrow")
  case today + 2:
    fmt.Println("In two days")
  default:
    fmt.Println("Too far away")
  }

  fmt.Println("counting ")
  defer func ()  {
    if i := recover(); i != nil {
      fmt.Println("recover from defer: ", i)
    }
  }()
  for i := 0; i < 10; i++ {
    if i > 5 {
      panic(i)
    }
    defer fmt.Println(i)
  }
  fmt.Println("done")

  fmt.Println(sum)

  fmt.Println(math.Sqrt(3))

  fmt.Println(pow(3,3,10))

  fmt.Println("defered func: ", deferedfun())
}
