# Go Benchpress
[![codecov](https://codecov.io/gh/rpickz/go-benchpress/branch/main/graph/badge.svg?token=D3097TFF89)](https://codecov.io/gh/rpickz/go-benchpress)

Go Benchpress is a visualisation utility for comparing the results of benchmarks.

Benchmarks are a really useful tool for comparing a couple of different strategies for achieving a goal, but also for 
an easy understanding of performance regressions within the system.

Sub-benchmarks are particularly useful for understanding how a particular feature operates over different data-set sizes.

Go Benchpress simplifies visualising these benchmark results - designed particularly for sub-benchmarks and demonstrating
how your code performs at different data-set sizes - visually.

## What Is Output?

Go Benchpress outputs, by default, a file of results for each benchmark passed in (via the Benchmark result data).

This is so that sub-benchmarks can be compared more easily.

A second mode of operation exists - 'no separation'.  This is where the results of __all__ the passed in benchmarks
are output into a single file.

You will find more information on the specifics of this in the [CSV Parser](./examples/csvparser) package - the README
contains details instructions on CLI usage, and how to make use of these different modes.

So, given Go Benchpress will either output a file for each benchmark, or a single file for all benchmarks, what formats
are available?

## What Does It Look Like?

Go Benchpress output can look like a variety of things.

Graphically, this is what a Go Benchpress output looks like:

![Example Output](examples/csvparser/example_output.svg "Example Output")

As you can see, it compares all the results for a single benchmark (across its various sub-benchmarks), displaying
their relative values visually.

You can choose between several dimensions (including nanoseconds per operation, bytes per operation, etc.) - for
the most recent advice on this, please consult the help text using:
```bash
gobenchpress -help
```

There are also other formats to choose from - overall the following formats are supported:
1. SVG (as a bar chart)
2. PNG (as a bar chart)
3. JSON
4. CSV
5. XML

See the example [CSV Parser](./examples/csvparser) package for instructions on how to use different formats, and what
they look like.

## How to Install?

Run the following command at a terminal:
```bash
go get github.com/rpickz/go-benchpress/cmd/gobenchpress
```

## How to Use?

See the example [CSV Parser](./examples/csvparser) package for instructions on how to use.

For more detailed instructions, see the CLI usage info with the following:
```bash
gobenchpress -help
```

## License?

MIT License.
