### dep
```
$ dep ensure --vendor-only
```

### How to Use
```
$ go run main.go -latest=true ginco
```

## brew
### brew install
```
$ brew tap tomokazukozuma/go-twitter
$ brew install go-twitter
```

### how to use
Search Trend Tweet
```
$ go-twitter SEARCH_WORD
```

Search Latest Tweet
```
$ go-twitter -latest SEARCH_WORD
```