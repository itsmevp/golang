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
#### When you have test cases (`_test.go` files) in your executable(`main`) package, you can’t simply execute `go run *.go` to run the project.
#### `*.go` part also matches the test files (`_test.go` files) and go run command can’t run them and returns `go run: cannot run *_test.go` files (`hello_test.go`) error.
