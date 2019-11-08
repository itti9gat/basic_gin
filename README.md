# Basic GIN
[![CircleCI](https://circleci.com/gh/itti9gat/basic_gin.svg?style=shield)]()

### Usage
1. Make sure you have Golang installed on your local machine. 
For more instructions on how to install Golang, look [here](https://golang.org/doc/install).

2. Install MySQL and insert data from /init/user.txt

3. Run script
```
$ git clone https://github.com/itti9gat/basic_gin.git
$ cd basic_gin
$ go mod init
$ go mod vendor
$ go run main.go
```
Open Browser go to 
http://localhost:8012/user


4. Test script
```
$ ENVIRONMENT=test go test ./...
```
