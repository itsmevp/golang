/*
To upgrade to a latest package:
go get -u github.com/sirupsen/logurus

When you do a go build your tests files are ignored so the dependencies under test files are not resolved.
You can run go mod tidy command that resolves all the dependencies that are inculded under test files as well.

Why do i use this dependency?
Run go mod why github.com/sirupsen/logurus
go mod why -m github.com/sirupsen/logurus to see the chain
*/
