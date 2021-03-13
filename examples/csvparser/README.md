# CSV Parser Example

This CSV Parser is a simple example which provides benchmark data which can be compared
using Go Benchpress.

To run the benchmarks in this directory, use the following at the terminal:
```bash
go test -bench .
```

Or, you could use the following if you'd also like memory statistics:
```bash
go test -bench . -benchmem
```

This will output data which can be processed by Go Benchpress.

You could save this output to a file, and use it as the input for Go Benchpress, or
alternatively use STDIN to stream the output directly into Go Benchpress.

## Saving Benchmark Results to File

### 1. Get Benchmark Data

Use the following command to save the result to a file:
```bash
go test -bench . >output.txt
```

Or, use the following to get memory statistics too:
```bash
go test -bench . -benchmem >output.txt
```

### 2. Run Go Benchpress

If you have installed Go Benchpress, you can use the following command to visualise
the results of the benchmarks:
```bash
gobenchpress -input output.txt
```

You should now see a file named `output.svg` within that directory - open that file
and you should see the sub-benchmarks laid out in a bar chart.

## Streaming Benchmark Results to Go Benchpress

Instead of performing the two-step process outlined above, you can actually achieve
the same results in one command.

Assuming you have Go Benchpress installed, run the following at the terminal:
```bash
go test -bench . | gobenchpress
```

You should now see a file named `output.svg` within that directory - open that file
and you should see the sub-benchmarks laid out in a bar chart.
