package main

import (
	"encoding/hex"
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
	if i == -1 {
		println("[*] Offset not found")
	} else {
		println("[*]", i)
	}
}

func PatternOffsetHex(offset string) {
	var maxPatternHex string
	maxPatternHex = PatternCreate(20280)
	offsetAscii, err := hex.DecodeString(offset[2:])
	if err != nil {
		println("[-] Invalid hex offset")
		os.Exit(0)
	} else {
		j := strings.Index(maxPatternHex, string(offsetAscii))

		if j == -1 {
			println("[*] Offset not found")
		} else {
			println("[*]", j)
		}
	}
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
		println("[i] ./gottern -h for help")
		os.Exit(0)
	} else if offset == "" && create > 0 {
		var patternCreated = PatternCreate(create)
		//var patternCreated = PatternCreate(10)
		println(patternCreated)
	} else if offset != "" && create == 0 {
		if len(offset) < 4 || (len(offset) < 10 && offset[:2] == "0x") {
			println("[!] Offset should be at least 4 bytes")
			os.Exit(0)
		} else if offset[:2] == "0x" {
			PatternOffsetHex(offset)
		} else {
			PatternOffset(offset)
		}
	} else {
		println("[i] ./gottern -h for help")
		os.Exit(0)
	}
}
