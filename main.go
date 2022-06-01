package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify cnf file.")
		os.Exit(-1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(-1)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	hasToken := sc.Scan()
	if !hasToken {
		os.Exit(-1)
	}

	line := sc.Text()
	fmt.Println(line)
	ary := strings.Split(line, " ")
	if ary[0] != "p" {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}

	if ary[1] != "cnf" {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}

	varCnt, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}
	lineCnt, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}

	fmt.Printf("var count:%d, line count:%d\n", varCnt, lineCnt)
}
