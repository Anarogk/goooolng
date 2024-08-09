
# Iteration

In Go there are *no* ``while``, ``do``, ``until`` keywords, you can only use ``for``. Which is a good thing!


### Declated variables
```go
var repeated string
```

## For Loop 
its a loop that executes a block of code for a specified number of times.Its syntax is similar to that of for loop in C/C++.
```go
var repeated string
for i := 0; i < 5; i++ {
	repeated = repeated + character
}
```

## Benchmarking

The `testing.B` gives you access to the cryptically named `b.N`.

When the benchmark code is executed, it runs `b.N` times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do ``go test -bench=.`` (or if you're in Windows Powershell ``go test -bench="."``)
```go
func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForLoop("a", 3)
	}
}
```

Command to run benchmark tests
```shellscript
go test -bench=.
```

Output:
```shellscript
goos: windows
goarch: amd64
pkg: integers
cpu: Intel(R) Core(TM) i5-10500H CPU @ 2.50GHz
BenchmarkAdd-12         1000000000               0.2358 ns/opPASS
ok      integers        0.386s 
```
