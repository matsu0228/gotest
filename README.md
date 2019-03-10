# gotest

[![CircleCI](https://circleci.com/gh/matsu0228/gotest.svg?style=svg)](https://circleci.com/gh/matsu0228/gotest)

* this repository contains examples for testing method of golang.
* [document at Qiita (japanese)](https://qiita.com/matsu0228/items/c9f69b036f0ab47b617e)

## tips

* table test
* setup/teardown

## integrate test

* unit test with mocking
* switch "unit test" and "integrate test" with build tag
* integrate test with docker on circle.CI

## exercise

* requirement
  * golang
  * docker / docker-compose

* setup
```
go get github.com/matsu0228/gotest

# build mysql on your environment
cd $GOPATH/src/github.com/matsu0228/gotest/infla
docker-compose up -d

# create dabatase/table
go run init/main.go
```

* exercise of some tests
```
# table test & setup
go test -v github.com/matsu0228/gotest/tips 

# unit test with mock
go test -v github.com/matsu0228/gotest/integrate

# switch "unit testing" and "integrate testing" with build tag
go test -v github.com/matsu0228/gotest/integrate/repository
go test -v -tags=integration github.com/matsu0228/gotest/integrate/repository

# atutomation with circle.CI with mysql
# --> see .circle.ci/config.yml
```
