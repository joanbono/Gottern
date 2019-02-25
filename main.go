/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

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
		PatternLittleEndian(offset)
	} else {
		println("[*]", i)
	}
}

func PatternLittleEndian(offset string) {
	var offsetLE string
	if len(offset) >= 8 {
		offsetLE = offset[6:8] + offset[4:6] + offset[2:4] + offset[0:2]
		PatternOffsetHex(offsetLE)
	} else {
		println("[*] Offset not found")
	}
}

func PatternOffsetHex(offset string) {
	var maxPatternHex string
	var offsetAscii []byte
	var j int
	var err error
	maxPatternHex = PatternCreate(20280)
	if offset[:2] == "0x" {
		offsetAscii, err = hex.DecodeString(offset[2:])
	} else {
		offsetAscii, err = hex.DecodeString(offset)
	}
	if err != nil {
		println("[-] Invalid offset")
	} else {
		j = strings.Index(maxPatternHex, string(offsetAscii))

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
