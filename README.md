# banfunc

`banfunc` is a Go linter that reports the call of a banned function.

# Installing
```bash
$ go install github.com/masakurapa/banfunc/cmd/banfunc@latest
```

# Usage

The -ban option is mandatory and specifies the banned function names.<br>
Multiple function names can be specified separated by commas.

**Example usage:**
```bash
banfunc -ban Println ./...
```

**Example with multiple function names:**
```bash
banfunc -ban Println,Print,Printf ./...
```

# Features to Implement if Possible

- Ban specific functions of a package (e.g., `fmt.Println`)
- Ban functions implemented in specific package structures or interface
- Configuration file loading
