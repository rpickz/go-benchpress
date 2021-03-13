# Go Benchpress

Go Benchpress is a visualisation utility for comparing the results of benchmarks.

Benchmarks are a really useful tool for comparing a couple of different strategies for achieving a goal, but also for 
an easy understanding of performance regressions within the system.

Sub-benchmarks are particularly useful for understanding how a particular feature operates over different data-set sizes.

Go Benchpress simplifies visualising these benchmark results - designed particularly for sub-benchmarks and demonstrating
how your code performs at different data-set sizes - visually.

## How to Install?

Run the following command at a terminal:
```bash
go get github.com/rpickz/go-benchpress
```

## How to Use?

See the example [CSV Parser](./examples/csvparser) package for instructions on how to use.

For more detailed instructions, see the CLI usage info with the following:
```bash
gobenchpress -help
```

## License?

MIT License.
