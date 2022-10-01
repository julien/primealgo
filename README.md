These are the exercises for [@ThePrimeagen](https://github.com/ThePrimeagen)'s algorithm
[course](https://frontendmasters.com/courses/algorithms) on FrontendMasters,
but instead of using TypeScript it uses Go.

Since the course is free (and good), I thought providing examples with
a different programming language could be useful and nice contribution.

The code can probably be improved so don't hesitate to send a pull request.
The code in this repository has no structure (i.e. different packages, etc...).
I might work on that later.

The test are based on the [code](https://github.com/ThePrimeagen/kata-machine)
used during the course.

Some examples uses generics, which means you need Go >= 1.18.

To run the tests (assuming you have Go installed correctly), use:

```sh
go test -cover -coverprofile=cov -race
```

If you have the `cover` tool installed and want to see the test coverage in
your browser, run the following command:

```sh
go tool cover -html=cov
```

If you want to view coverage details in your terminal, use:

```sh
go tool cover -func=cov
```

There are some basic benchmarks (very few at the moment) that you can run with:

```sh
go test -bench=.
```
