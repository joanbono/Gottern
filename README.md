# Gottern
Golang port for Metasploit's `pattern_create` and `pattern_offset`.  
Based on [`haxpattern`](https://github.com/DharmaOfCode/haxpattern)

## Usage
```
$ gottern -h
Usage of haxpattern.exe:
   -c int 	   pattern_create
   -o string   pattern_offset
 ```
 
 ## Create a pattern 

 Create a pattern using the `-c` flag with the size of the pattern to be created.
 
 ```bash
 $ gottern -c 200
Aa0Aa1Aa2Aa3Aa4[...]Ag1Ag2Ag3Ag4Ag5Ag
```

## Look for an offset

```
$ gottern -o 6Aj7
[*] 290
```

***

## TODO

- [ ] Add Little/Big endian option
- [ ] Look for hex pair in offset