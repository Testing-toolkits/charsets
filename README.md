[![Go Reference](https://pkg.go.dev/badge/github.com/Testing-toolkits/charsets.svg)](https://pkg.go.dev/github.com/Testing-toolkits/charsets)

# charsets
Package contains various groups of charsets and utility function to choose random runes from them. 

Example usage:
```go
rndRunes := charsets.RandomRunes(10, []rune(charsets.Japanese))
```