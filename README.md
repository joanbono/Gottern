![](img/Gottern_banner.png)

[![GitHub Issues](https://img.shields.io/github/issues/joanbono/gottern.svg)](https://github.com/gocaio/goca/issues)
[![GitHub tag](https://img.shields.io/github/tag/joanbono/gottern.svg)](https://github.com/gocaio/goca/tags)
[![Go Version](https://img.shields.io/badge/go-1.11-blue.svg?logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

Golang port for Metasploit's `pattern_create` and `pattern_offset`.  
Based on [`haxpattern`](https://github.com/DharmaOfCode/haxpattern)

## Usage
```
$ gottern -h
Usage of haxpattern.exe:
   -c int 	   pattern_create
   -o string   pattern_offset
 ```
 
### Create a pattern 

 Create a pattern using the `-c` flag with the size of the pattern to be created.
 
 ```bash
 $ gottern -c 200
Aa0Aa1Aa2Aa3Aa4[...]Ag1Ag2Ag3Ag4Ag5Ag
```

### Look for an offset

```
$ gottern -o 6Aj7
[*] 290
$ gottern -o 0x36416a37
[*] 290
```

***

## Benchmarks

### Create a Pattern

![](img/benchmark_create.png) 

### Find an offset

![](img/benchmark_offset.png) 
