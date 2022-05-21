# Shoulders

Inspired by the [SHOULDERS.md](https://github.com/gobuffalo/buffalo/blob/master/SHOULDERS.md) several projects have started to create their own implementations. This tool will create a custom SHOULDERS.md file for any given Go project allowing OSS maintainers to recognize those who's OSS contributions helped them.

## Installation

```bash
$ go install github.com/gobuffalo/shoulders@latest
```

## Usage

When run without any flags the `shoulders` command will print the `SHOULDERS.md` to the `STDOUT`.

```bash
$ shoulders
# github.com/gobuffalo/shoulders Stands on the Shoulders of Giants

github.com/gobuffalo/shoulders does not try to reinvent the wheel! Instead, it uses the already great wheels developed by the Go community and puts them all together in the best way possible. Without these giants, this project would not be possible. Please make sure to check them out and thank them for all of their hard work.

Thank you to the following **GIANTS**:

* [github.com/yuin/goldmark](https://godoc.org/github.com/yuin/goldmark)
* [golang.org/x/crypto](https://godoc.org/golang.org/x/crypto)
* [golang.org/x/mod](https://godoc.org/golang.org/x/mod)
* [golang.org/x/net](https://godoc.org/golang.org/x/net)
* [golang.org/x/sync](https://godoc.org/golang.org/x/sync)
* [golang.org/x/sys](https://godoc.org/golang.org/x/sys)
* [golang.org/x/text](https://godoc.org/golang.org/x/text)
* [golang.org/x/tools](https://godoc.org/golang.org/x/tools)
* [golang.org/x/xerrors](https://godoc.org/golang.org/x/xerrors)
```

### The `-w` (Write) Flag

To write the `SHOULDERS.md` file to disk use the `-w` flag.

```bash
$ shoulders -w
```

### The `-n` (Name) Flag

By default the "name" of the project is the current package name. To change that use the `-n` flag.

```bash
$ shoulders -n shoulders
# shoulders Stands on the Shoulders of Giants

shoulder does not try to reinvent the wheel! Instead, it uses the already great wheels developed by the Go community and puts them all together in the best way possible. Without these giants, this project would not be possible. Please make sure to check them out and thank them for all of their hard work.

<...>
```

### The `-j` (JSON) Flag

To get a JSON array of the dependencies of the project use the `-j` flag.

```bash
$ shoulders -j
["github.com/yuin/goldmark","golang.org/x/crypto","golang.org/x/mod","golang.org/x/net","golang.org/x/sync","golang.org/x/sys","golang.org/x/text","golang.org/x/tools","golang.org/x/xerrors"]
```

