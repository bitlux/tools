# tools
Various tools, mostly for Linux

[![Go Reference](https://pkg.go.dev/badge/github.com/bitlux/tools.svg)](https://pkg.go.dev/github.com/bitlux/tools)
[![Lint Status](https://github.com/bitlux/tools/workflows/golangci-lint/badge.svg)](https://github.com/bitlux/tools/actions?query=workflow%3Agolangci-lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitlux/tools)](https://goreportcard.com/report/github.com/bitlux/tools)

## autotmux

`autotmux` will attach to an existing tmux session, if there is exactly one. If there is more than one,
it will print how many there are. If there are 0, it will do nothing.

```
go install github.com/bitlux/tools/autotmux@latest
```

## red

`red` will print a large, red rectangle to the screen. This is useful for separating the output of
commands, especially when the output is large.

```
go install github.com/bitlux/tools/red@latest
```

## random

`random` will print one of its arguments, chosen at random.

```
$ random 1 2 3
2
```

```
go install github.com/bitlux/tools/random@latest
```
