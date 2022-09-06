To reproduce: 

```bash
$ cd nested
$ golangci-lint --fix --path-prefix nested
```

Expected: file is formatted, no errors reported from the linters. 

What happens: 

```
$ golangci-lint run --fix --path-prefix nested
ERRO Failed to fix issues in file nested/main.go: failed to get file bytes for nested/main.go: can't read file nested/main.go: open nested/main.go: no such file or directory 
nested/main.go:5: File is not `gofmt`-ed with `-s` (gofmt)
```