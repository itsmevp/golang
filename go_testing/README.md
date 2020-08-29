#### get verbose test output
```
go test -v
```
#### if you need to colorize the test outputs, for example, green color for the passed test and red color for failed tests, you can use gotest package
```
go get -u github.com/rakyll/gotest
gotest -v
```
#### if you have lots of test files with test functions but you want to selectively run few
```
go test -v -run TestSomeFunction
```
#### you can specify the selected test files to run but in that case, Go will not include other files into the compilation which might be needed by the test. You need to specify dependency files as well.
```
go test -v hello.go hello_test.go
```