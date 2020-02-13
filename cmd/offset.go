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

package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// offsetCmd represents the offset command
var offsetCmd = &cobra.Command{
	Use:   "offset",
	Short: "Search for an offset",
	Run: func(cmd *cobra.Command, args []string) {
		offset, err := cmd.Flags().GetString("query")
		CheckErr(err)
		bigendian, err := cmd.Flags().GetBool("bigendian")
		CheckErr(err)

		if len(offset) < 4 || (len(offset) < 10 && offset[:2] == "0x") {
			fmt.Printf("\n[i] ./gottern offset -h for help\n\n")
			os.Exit(0)
		} else if offset[:2] == "0x" {
			PatternOffsetHex(offset)
		} else {
			PatternOffset(offset, bigendian)
		}
	},
}

func init() {
	rootCmd.AddCommand(offsetCmd)
	offsetCmd.PersistentFlags().StringP("query", "q", "", "Query the following pattern. Minimum 4 bytes.")
	offsetCmd.PersistentFlags().BoolP("bigendian", "b", false, "Search for Big Endian Offset")
	offsetCmd.MarkFlagRequired("length")
}

// PatternOffset looks for ASCII string
// inside the pattern
func PatternOffset(offset string, bigendian bool) {
	var maxPattern string
	maxPattern = PatternCreate(20280)
	i := strings.Index(maxPattern, offset)
	if i == -1 && bigendian == false {
		PatternLittleEndian(offset)
	} else if i == -1 && bigendian == true {
		PatternOffsetHex(offset)
	} else {
		fmt.Printf("[*] %v\n", i)
	}
}

// PatternLittleEndian changes the offset
// in plain HEX to Little Endian format
func PatternLittleEndian(offset string) {
	var offsetLE string
	if len(offset) >= 8 {
		offsetLE = offset[6:8] + offset[4:6] + offset[2:4] + offset[0:2]
		PatternOffsetHex(offsetLE)
	} else {
		println("[*] Offset not found")
	}
}

// PatterOffsetHex will look for the offset
// in the pattern if starts by 0x or its
// HEX in Little Endian format
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
