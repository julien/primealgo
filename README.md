These are the exercises for [ThePrimeagen](https://github.com/ThePrimeagen/), 's
algorithm's [course](https://frontendmasters.com/courses/algorithms) on Frontend Masters,
but instead of using TypeScript it uses Go.

The test are based on the [source code](https://github.com/ThePrimeagen/kata-machine) used 
during the course, but everything is written in Go.

This is not a package, just some examples.
To run the tests (assuming you have go installed correctly), use:

```
go test -cover -coverprofile=cov -race
```

If you have the `cover` tool installed and want to see the test coverage in your browser, 
run the following command: 

```
go tool cover -html=cov
```
