package main

import (
	"flag"
	"os"
	"strings"
)

var offset string
var create int

func init() {
	flag.IntVar(&create, "c", 0, "pattern_create")
	flag.StringVar(&offset, "o", "", "pattern_offset")
}

func PatternOffset(offset, p string) {
	println("In Pattern Offset")
	//var patternCreated = PatternCreate(10)
	//var patternCreated = "Aa0Aa1Aa2Aa3Aa4Aa5Aa6Aa7Aa8Aa9Ab0Ab1Ab2Ab3Ab4Ab5Ab"

	//println(p)
	//i := strings.Index(p, offset)
	//println(i)
}

func PatternCreate(lenght int) string {
	//println("Jere")
	UpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase := "abcdefghijklmnopqrstuvwxyz"
	Numbers := "0123456789"

	var pattern []string
	for len(pattern) < create {
		for _, A0 := range UpperCase {
			for _, a0 := range LowerCase {
				for _, n0 := range Numbers {
					if len(pattern) < lenght {
						pattern = append(pattern, string(A0))
					}
					if len(pattern) < lenght {
						pattern = append(pattern, string(a0))
					}
					if len(pattern) < lenght {
						pattern = append(pattern, string(n0))
					}
				}
			}
		}
	}
	return strings.Join(pattern, "")
}

func main() {
	flag.Parse()

	if (offset == "" && create == 0) || (offset != "" && create != 0) {
		println("./gottern -h for help")
		os.Exit(0)
	} else if offset == "" && create > 0 {
		var patternCreated = PatternCreate(create)
		//var patternCreated = PatternCreate(10)
		println(patternCreated)
	} else if offset != "" && create == 0 {
		if len(offset) < 4 {
			println("Offset should be at least 4 bytes")
			os.Exit(0)
		}
		var i = 10
		p := PatternCreate(i)
		println(len(p))
		PatternOffset(offset, p)
	} else {
		println("./gottern -h for help")
		os.Exit(0)
	}
}
