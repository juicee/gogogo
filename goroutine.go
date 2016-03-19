package main

import "fmt"
import "time"
import "code.google.com/p/go-tour/tree"

func say(s string)  {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func sum(a []int, c chan int)  {
  sum := 0
  for _, v := range a {
    sum += v
  }
  c <- sum
}

// func fibonacci(n int, c chan int)  {
//   x, y := 0, 1
//   for i := 0; i < n; i++ {
//     c <- x
//     x, y = y, x + y
//   }
//   //close(c)
// }

func fibonacci(c, quit chan int)  {
    x, y := 0, 1
    for {
      select {
      case c <- x:
        x, y = y, x+y
      case <-quit:
        fmt.Println("quit")
        return
      }
    }
}

func WalkRecursive(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  if t.Left != nil {
    WalkRecursive(t.Left, ch)
  }
  ch <- t.Value
  if t.Right != nil {
    WalkRecursive(t.Right, ch)
  }
}

func Walk(t *tree.Tree, ch chan int)  {
  WalkRecursive(t, ch)
  close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for {
    node1, ok1 := <-ch1
    node2, ok2 := <-ch2
    if !ok1 || !ok2 {
      if !ok1 && !ok2 {
        return true
      } else {
        return false
      }
    }
    if ok1 && ok2 {
      if node1 != node2 {
        return false
      } else {
        continue
      }
    }
  }
  return true
}

type Fetcher interface {
  Fetch(url string) (body string, urls []string, err error)
}

var fetchHistory map[string]int

func Crawl(url string, depth int, fetcher Fetcher) {
  if depth <= 0 {
    return
  }
  if _, ok := fetchHistory[url]; ok {
    return
  }
  body, urls, err := fetcher.Fetch(url)
  fetchHistory[url] = 1
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    go Crawl(u, depth - 1, fetcher)
  }
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
  body string
  urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := f[url]; ok {
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher {
"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func main()  {
  // ch := make(chan int)
  // go Walk(tree.New(1), ch)

  // for {
  //   select {
  //   case node := <-ch:
  //     fmt.Println("tree node: ", node)
  //   default:
  //     break;
  //   }
  // }
  fetchHistory = make(map[string]int)
  Crawl("http://golang.org/", 4, fetcher)

  isSame := Same(tree.New(1), tree.New(2))
  fmt.Println("trees are the same: ", isSame)


  // go say("world")
  // say("hello")
  a := []int {1,2,3,4,5}
  c := make(chan int)
  go sum(a[:len(a)/2], c)
  go sum(a[len(a)/2:], c)
  x, y := <-c, <-c

  fmt.Println(x, y, x+y)

  // d := make(chan int, 2)
  // d <- 1
  // d <- 2
  //
  // fmt.Println(<-d)
  // fmt.Println(<-d)
  //
  // e := make(chan int, 10)
  // go fibonacci(cap(e), e)
  // for i := range e {
  //   fmt.Println(i)
  // }

  f := make(chan int)
  quit := make(chan int)

  go func ()  {
    for i := 0; i < 10; i++ {
      if i == 7 {
        quit <- 0
      }
      fmt.Println(<-f)
    }
    quit <- 0
  }()
  fibonacci(f, quit)

  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(500 * time.Millisecond)

  for {
    select {
    case <-tick:
      fmt.Println("tick.")
    case <-boom:
      fmt.Println("BOOM!")
      return
    default:
      fmt.Println("    .")
      time.Sleep(200 * time.Millisecond)
    }
  }
}
