package main

import (
	"fmt"
	"regexp"
)

func main() {
	bmatch, err := regexp.Match("hello", []byte("hello world"))
	fmt.Println(bmatch, err)

	wordRx := regexp.MustCompile(`\w+`)
	if matches := wordRx.FindAllString("123abc@dfr@82eri", -1); matches != nil {
		fmt.Println(matches)
	}
}
