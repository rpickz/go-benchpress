# Go Benchpress

Go Benchpress is a visualisation utility for comparing the results of benchmarks.

Benchmarks are a really useful tool for comparing a couple of different strategies for achieving a goal, but also for 
an easy understanding of performance regressions within the system.

Sub-benchmarks are particularly useful for understanding how a particular feature operates over different data-set sizes.

Go Benchpress simplifies visualising these benchmark results - designed particularly for sub-benchmarks and demonstrating
how your code performs at different data-set sizes - visually.

## What Does It Look Like?

This is what a Go Benchpress output looks like:

![Example Output](examples/csvparser/example_output.svg "Example Output")

As you can see, it compares all the results for a single benchmark (across its various sub-benchmarks), displaying
their relative values visually.

You can choose between several dimensions (including nanoseconds per operation, bytes per operation, etc.) - for
the most recent advise on this, please consult the help text using:
```bash
gobenchpress -help
```

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
