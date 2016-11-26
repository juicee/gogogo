package main

import (
    "fmt"
    _ "io"
)

type TestInner struct {
    x,y int
}

func (inner *TestInner) inner() int {
    return 123
}

type TestOuter struct  {
    TestInner
    x int

}

type innerInterface interface {
    inner() int
}


func RunInner(inner innerInterface) int {
    return inner.inner()
}

func main()  {
    test := &TestOuter{}
    test.TestInner.x = 1
    test.x = 2
    fmt.Println(test.x)
    fmt.Println(RunInner(test))
}