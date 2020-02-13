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
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the pattern",
	Run: func(cmd *cobra.Command, args []string) {
		Length, err := cmd.Flags().GetInt("length")
		CheckErr(err)
		if Length == 0 {
			fmt.Printf("\n[i] ./gottern create -h for help\n\n")
			os.Exit(0)
		}
		fmt.Printf("%v\n", PatternCreate(Length))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().IntP("length", "l", 0, "Lenght of the string to be created")
	createCmd.MarkFlagRequired("length")
}

// PatterCreate will just create the Pattern
// with [A-Za-z0-9]
func PatternCreate(length int) string {
	UpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase := "abcdefghijklmnopqrstuvwxyz"
	Numbers := "0123456789"

	var pattern []string
	for len(pattern) < length {
		for _, A0 := range UpperCase {
			for _, a0 := range LowerCase {
				for _, n0 := range Numbers {
					if len(pattern) < length {
						pattern = append(pattern, string(A0))
					}
					if len(pattern) < length {
						pattern = append(pattern, string(a0))
					}
					if len(pattern) < length {
						pattern = append(pattern, string(n0))
					}
				}
			}
		}
	}
	return strings.Join(pattern, "")
}
