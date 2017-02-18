# functions-keystonemiddleware
OpenStack Identity (Keystone) middleware for IronFunctions


## Building a middleware

Use following commands:
```bash
$ export GOPATH=~/go
$ git clone git@github.com:iron-io/functions.git ${GOPATH}/src/github.com/iron-io/functions
$ git clone git@github.com:denismakogon/functions-keystonemiddleware.git ${GOPATH}/src/github.com/denismakogon/functions-keystonemiddleware
$ pushd ${GOPATH}/src/github.com/iron-io/functions && make dep build; popd
$ pushd ${GOPATH}/src/github.com/denismakogon/functions-keystonemiddleware && make build; popd
```
