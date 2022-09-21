These are the exercises for @ThePrimeagen's algorithm [course](https://frontendmasters.com/courses/algorithms)
on FrontendMasters, but instead of using TypeScript it uses Go.
Since the course is free (and good), I thought providing examples with
another programming language could be useful and contribute.

The test are based on the [source code](https://github.com/ThePrimeagen/kata-machine) used
during the course.

To run the tests (assuming you have Go installed correctly), use:

```sh
go test -cover -coverprofile=cov -race
```

If you have the `cover` tool installed and want to see the test coverage in
your browser, run the following command:

```sh
go tool cover -html=cov
```
