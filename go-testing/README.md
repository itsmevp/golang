#### Get verbose test output.
```
go test -v
```
#### If you need to colorize the test outputs, for example, green color for the passed test and red color for failed tests, you can use gotest package.
```
go get -u github.com/rakyll/gotest
gotest -v
```
#### If you have lots of test files with test functions but you want to selectively run few.
```
go test -v -run TestSomeFunction
```
#### You can specify the selected test files to run but in that case, Go will not include other files into the compilation which might be needed by the test. You need to specify dependency files as well.
```
go test -v hello.go hello_test.go
```
**When you have test cases (`_test.go` files) in your executable(`main`) package, you can’t simply execute `go run *.go` to run the project.**  
**`*.go` part also matches the test files (`_test.go` files) and go run command can’t run them and returns `go run: cannot run *_test.go` files (`hello_test.go`) error.**

#### Go provide built-in functionality to check your code coverage.
```
go test -v -cover .
```

#### To check what did we miss to cover in the test we use `-coverprofile` option.
```
go test -v -coverprofile=cover.txt .
```

#### Out of many built in tools, go provides `cover` tool to analyze the code coverage details.
```
go tool cover -html=cover.txt -o cover.html
```

#### To test all the packages in a module, you can use `go test ./...` command in which `./...` matches all the packages in the module.

#### There are two ways to run tests `local directory mode` and `package list mode`.
- **local directory mode:**
    - Local directory mode occurs when `go test` is invoked with no package arguments.
    - In this mode, caching is disabled.
    - In this mode, `go test` compiles the package sources and tests found in the current directory and then runs the resulting test binary.
    - Example: `go test` or `go test -v`
- **Package list mode:**
    - Package list mode occurs when go test is invoked with explicit package arguments.
    - In this mode go test caches successfully package test results to avoid unnecessary repeated running of tests.
    - Whenever Go run tests on a package, Go creates a test binary and runs it. 
    - You can output this binary using -c flag with go test command. This will output .test file but won’t run it. 
    - Additionally, rename this file using -o flag.

```
v01@Vishnus-Mac go_testing$ go test -v -c
v01@Vishnus-Mac go_testing$ ./go_testing.test -test.v
=== RUN   TestGreet
--- PASS: TestGreet (0.00s)
    hello_test.go:12: Expected Hello Dude! and got Hello Dude!
PASS
```