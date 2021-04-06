package main

import (
	"flag"
	"fmt"
	"os"
)

//first init var
//second init func
//last main func

var nickName string
var T = a()

func a() int64 {
	fmt.Println("calling a()")
	return 1
}

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	fmt.Println(T)
	flag.StringVar(&nickName, "nickName", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", nickName)
}