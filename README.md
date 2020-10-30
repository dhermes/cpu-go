# `cpu-go`

[![GoDoc][3]][4]

This code is graciously borrowed from [jonhoo/drwmutex][1] which I learned
about in ["So You Wanna Go Fast?"][2] by Tyler Treat.

## Examples

### Linux

```
$ uname -a
Linux noobuntu 5.4.0-51-generic #56-Ubuntu SMP Mon Oct 5 14:28:49 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
$ go run ./examples/main.go
Current map of CPUs:
| 0 -> 0
| 1 -> 4
| 2 -> 1
| 3 -> 5
| 4 -> 2
| 5 -> 6
| 6 -> 3
| 7 -> 7
Current CPU: 1
```

### macOS

```
$ uname -a
Darwin waddling-duck 19.6.0 Darwin Kernel Version 19.6.0: Mon Aug 31 22:12:52 PDT 2020; root:xnu-6153.141.2~1/RELEASE_X86_64 x86_64
$ go run ./examples/main.go
Current map of CPUs:
| 0 -> 0
| 1 -> 1
| 2 -> 2
| 3 -> 3
| 4 -> 4
| 5 -> 5
| 6 -> 6
| 7 -> 7
| 8 -> 8
| 9 -> 9
| 10 -> 10
| 11 -> 11
Current CPU: 6
```

[1]: https://github.com/jonhoo/drwmutex/
[2]: https://www.youtube.com/watch?v=DJ4d_PZ6Gns
[3]: https://godoc.org/github.com/dhermes/cpu-go?status.svg
[4]: https://godoc.org/github.com/dhermes/cpu-go
