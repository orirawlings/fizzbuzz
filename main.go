// Based roughly on https://themonadreader.files.wordpress.com/2014/04/fizzbuzz.pdf
//
// Terribly impractical and a notably obfuscated implementation
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `usage: %s <i> [<d=word>]
Play FizzBuzz style games with arbitrary rules.
Print the first <i> natural numbers unless that number is divisible by d 
then print the corresponding word.  If number is divisible by multiple d 
values, print words for all matching d, in the same order the rules were
specified.
`

type c func(string) string

type rule struct {
	d int
	s string
}

func Game(rs []rule, n int) string {
	test := func(d int, s string, x c) c {
		if n%d == 0 {
			return func(_ string) string { return s + x("") }
		}
		return x
	}
	run := func(s string) string { return s }
	for i := len(rs) - 1; i >= 0; i-- {
		run = test(rs[i].d, rs[i].s, run)
	}
	return run(strconv.Itoa(n))
}

func printUsageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	fmt.Fprintf(os.Stderr, usage, os.Args[0])
	os.Exit(1)
}

func parseRule(arg string) rule {
	s := strings.SplitN(arg, "=", 2)
	d, err := strconv.Atoi(s[0])
	if err != nil || len(s) != 2 {
		printUsageAndExit(fmt.Sprintf("Unable to parse argument: %q", arg))
	}
	return rule{d, s[1]}
}

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit("Too few command line arguments")
	}
	r, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printUsageAndExit(fmt.Sprintf("Unable to parse argument: %q", os.Args[1]))
	}
	var rules []rule
	for _, arg := range os.Args[2:] {
		rules = append(rules, parseRule(arg))
	}
	for i := 1; i <= r; i++ {
		fmt.Println(Game(rules, i))
	}
}
