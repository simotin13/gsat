package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Clause struct {
	terms map[int64]bool
}

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

	if len(ary) < 4 {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}

	if ary[1] != "cnf" {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}

	varCnt, err := strconv.ParseInt(ary[2], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}
	lineCnt, err := strconv.ParseInt(ary[3], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid format:%s\n", line)
		os.Exit(-2)
	}

	var clauses = []Clause{}
	fmt.Printf("var count:%d, line count:%d\n", varCnt, lineCnt)
	for i := 0; i < int(lineCnt); i++ {
		var clause = Clause{}
		hasToken := sc.Scan()
		if !hasToken {
			os.Exit(-1)
		}

		line = sc.Text()
		ary = strings.Split(line, " ")
		pos := 0
		for {
			val, err := strconv.ParseInt(ary[pos], 10, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid value format:%s\n", line)
				os.Exit(-2)
			}
			if val == 0 {
				break
			}

			if 0 < val {
				// true
				clause.terms[val] = true
			} else {
				// false
				val = val * -1
				clause.terms[val] = false
			}
		}
		clauses = append(clauses, clause)
	}
}
