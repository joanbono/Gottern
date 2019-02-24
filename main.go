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

func PatternOffset(offset string) {
	var maxPattern string
	maxPattern = PatternCreate(20280)
	i := strings.Index(maxPattern, offset)
	//println("Offset found at ", i)
	println("[*]", i)
	//println("Offset len: ", len(offset))
	//s := maxPattern[i : i+len(offset)]
	//println(s)
}

func PatternCreate(lenght int) string {
	UpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase := "abcdefghijklmnopqrstuvwxyz"
	Numbers := "0123456789"

	var pattern []string
	for len(pattern) < lenght {
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
		PatternOffset(offset)
	} else {
		println("./gottern -h for help")
		os.Exit(0)
	}
}
