![](img/Gottern_banner.png)

[![GitHub Issues](https://img.shields.io/github/issues/joanbono/gottern.svg)](https://github.com/joanbono/gottern/issues)
[![GitHub tag](https://img.shields.io/github/tag/joanbono/gottern.svg)](https://github.com/joanbono/gottern/tags)
[![Go Version](https://img.shields.io/badge/go-1.13.7-blue.svg?logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/joanbono/gottern)](https://goreportcard.com/report/github.com/joanbono/gottern)


Golang port for Metasploit's `pattern_create` and `pattern_offset`.  
Based on [`haxpattern`](https://github.com/DharmaOfCode/haxpattern)

## Usage

```bash
$ gottern help 
Pattern Offset Seeker and Pattern Creator

Usage:
  Gottern [command]

Available Commands:
  create      Create the pattern
  help        Help about any command
  offset      Search for an offset
  version     Prints current Gottern version

Flags:
  -h, --help   help for Gottern

Use "Gottern [command] --help" for more information about a command.
```
 
### Create a pattern 

 Create a pattern using the `create` flag with the size (`-l`) of the pattern to be created.
 
 ```bash
 $ gottern create -h
Create the pattern

Usage:
  Gottern create [flags]

Flags:
  -h, --help         help for create
  -l, --length int   Lenght of the string to be created

$ gottern create -l 200
Aa0Aa1Aa2Aa3Aa4A[...]g1Ag2Ag3Ag4Ag5Ag
```

### Look for an offset

Search an offset using the `offset` flag with the query to perform (`-q`) of the pattern to be created. Use `-b` for big endian search.

```bash
$ gottern offset -h
Search for an offset

Usage:
  Gottern offset [flags]

Flags:
  -b, --bigendian      Search for Big Endian Offset
  -h, --help           help for offset
  -q, --query string   Query the following pattern. Minimum 4 bytes.
```

Examples:

```bash
# ASCII 
$ gottern offset -q 6Aj7
[*] 290
# Plain HEX
$ gottern offset -q 0x36416a37
[*] 290
# Little Endian HEX
$ gottern offset -q 376a4136
[*] 290
# Big Endian HEX
$ gottern offset -q 36416a37 -b
[*] 290
```

***

## Benchmarks

Some benchmarks using [`hyperfine`](https://github.com/sharkdp/hyperfine).

### Create a Pattern

![](img/benchmark_create.png) 

### Find an offset

![](img/benchmark_offset.png) 
