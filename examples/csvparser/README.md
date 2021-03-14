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

There are two modes of operation for Go Benchpress - to output all the results together (not separated by parent benchmark),
or to group them (separate them) based on their parent benchmark.

This section covers the latter approach - with the results separated into different files for each benchmark.

The former approach - to show all of the results unseparated is covered later in this guide.

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

You should now see files named as follows`output.svg` within that directory:
1. `output_BenchmarkParseCSVLineFieldLength.svg`
2. `output_BenchmarkParseCSVLineFields.svg`
 
If you open those files you should see the sub-benchmarks laid out in bar charts.

## Streaming Benchmark Results to Go Benchpress

Instead of performing the two-step process outlined above, you can actually achieve
the same results in one command.

Assuming you have Go Benchpress installed, run the following at the terminal:
```bash
go test -bench . | gobenchpress
```

You should now see files named as follows`output.svg` within that directory:
1. `output_BenchmarkParseCSVLineFieldLength.svg`
2. `output_BenchmarkParseCSVLineFields.svg`
 
If you open those files you should see the sub-benchmarks laid out in bar charts.

## Using Other Formats

Go Benchpress can output in different formats, other than bar charts.

Go Benchpress can also output in `JSON`, `CSV` and `XML` formats.

### JSON

To output in JSON format, use the `-renderType JSON` CLI switch, as follows:
```bash
go test -bench . | gobenchpress -renderType JSON
```

You should now see files named as follows within that directory:
1. `output_BenchmarkParseCSVLineFieldLength.json`
2. `output_BenchmarkParseCSVLineFields.json`

An example of the JSON output can be found here: [Example JSON Output](./example_json_output.json).

### CSV

To output in CSV format, use the `-renderType CSV` CLI switch, as follows:
```bash
go test -bench . | gobenchpress -renderType CSV
```

You should now see files named as follows within that directory:
1. `output_BenchmarkParseCSVLineFieldLength.csv`
2. `output_BenchmarkParseCSVLineFields.csv`

An example of the CSV output can be found here: [Example CSV Output](./example_csv_output.csv).

### XML

To output in XML format, use the `-renderType XML` CLI switch, as follows:
```bash
go test -bench . | gobenchpress -renderType XML
```

You should now see files named as follows within that directory:
1. `output_BenchmarkParseCSVLineFieldLength.xml`
2. `output_BenchmarkParseCSVLineFields.xml`

An example of the XML output can be found here: [Example XML Output](./example_xml_output.xml).

## Showing All Benchmarks Results Together (Unseparated)

To view all Go Benchpress results together (for all benchmarks input), you simply use the `-noSep` flag on the command line.

For instance, to render an CSV output with all benchmarks output __together__, you could use the following:
```bash
go test -bench . | gobenchpress -renderType CSV -noSep
```

This results in a CSV file with all of the provided benchmarks output together.

You can find an example of this output here: [Example Unseparated CSV Output](./example_csv_output_all_together.csv)

This same approach can be used for any of the other formats too.
