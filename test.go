package main

import (
    _ "bufio"
    "fmt"
    _ "io"
    "sort"
    "strings"
    _ "strings"
    _ "unicode/utf8"
    "math/rand"
    "encoding/json"
    "bytes"
    "time"
)

type FoldedStrings []string

func (slice FoldedStrings) Len() int {
    return len(slice)
}
func (slice FoldedStrings) Less(i, j int) bool {
    return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}
func (slice FoldedStrings) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func createCounter(num int) chan int {
    next := make(chan int)
    go func(i int) {
        for {
            next <- i
            i+=2
        }
    }(num)
    return next
}

func testReturn() int {
    a := 0
    b := 1
    if a == b {
        return a
    } else {
        return a
    }
}

type TestJson struct {
    Test int
    Test2 string
}

func multiArgs(test string, args ...int) {
    fmt.Println(args)
}

func main() {
    //	teststr := "aas一二䶍"
    //	fmt.Printf("%x\n", []rune(teststr))
    //	runes, size := utf8.DecodeRuneInString(teststr[6:])
    //	fmt.Printf("%d %s %d %X %d\n", len(teststr), teststr[9:], strings.Index(teststr, "二"), runes, size)

    //	for i, char := range teststr {
    //		fmt.Printf("%-2d	%U	'%c'	%X\n", i, char, char, []byte(string(char)))
    //	}

    //	as := ""
    //	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
    //		fmt.Printf("[0x%X '%c'] \n", char, char)
    //		as += string(char)
    //	}

    //	type polar struct {
    //		radius,
    //		o float64
    //	}

    //	p := polar{8.32, .49}
    //	fmt.Print(-18.5, 17, "Elephant", -8+.7i, 0x3c7, '\u03c7', "a", "b", p)

    //	helloworld := "hello world here"
    //	fmt.Println(strings.Title(helloworld), strings.ToTitle(helloworld))

    //	strreader := strings.NewReader(helloworld)

    //	for b, pos := strreader.ReadByte(); pos != io.EOF; b, pos = strreader.ReadByte() {
    //		fmt.Printf("%c", b)
    //	}

    teststr := []string{"abc", "Bcd", "Are", "bte", "aqw", "cer"}
    sort.Strings(teststr)
    fmt.Println(teststr)

    testfolded := FoldedStrings{"abc", "Bcd", "Are", "bte", "aqw", "cer"}
    sort.Sort(testfolded)
    fmt.Println(testfolded)

    var i interface{} = 99;
    j := i.(int)
    fmt.Println(j);

    counterA := createCounter(11)
    counterB := createCounter(102)

    for i := 0; i < 1; i++ {
        a := <-counterA
        fmt.Printf("A -> %d, B -> %d\n", a, <- counterB)
    }
    fmt.Println()

    channels := make([]chan bool, 6)
    for i := range channels {
        channels[i] = make(chan bool)
    }
    go func() {
        for {
           channels[rand.Intn(6)] <- true
        }
    }()

    for i := 0; i < 36; i++ {
        var x int
        select {
        case <-channels[0]:
            x = 1
        case <-channels[1]:
            x = 2
        case <-channels[2]:
            x = 3
        case <-channels[3]:
            x = 4
        case <-channels[4]:
            x = 5
        case <-channels[5]:
            x = 6
        }
        fmt.Printf("%d ", x)
    }
    fmt.Println()

    str := `{"Test": 1, "Test2": "hello", "test3": { "testinner": 1 }, "testarr": ["q", "w"]}`
    //var jsonobj interface{}
    var jsonobj TestJson
    fmt.Println(bytes.NewBufferString(str).Bytes())
    json.Unmarshal(bytes.NewBufferString(str).Bytes(), &jsonobj)
    fmt.Println(jsonobj)
    //jsonobj_ := jsonobj.(map[string]interface{})
    fmt.Println(jsonobj)
    //arrobj := jsonobj["testarr"].([]interface{})
    //fmt.Println(arrobj[0])
    strtime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Println(strtime)
    fmt.Println(time.Now().Unix())
    fmt.Println(<-time.After(100000))

    funca := func() string {
        fmt.Println("funca");
        return "123"
    }
    fmt.Println(funca())
}
