# functions-keystonemiddleware
OpenStack Identity (Keystone) middleware for IronFunctions


## Building a middleware

Use following commands:
```bash
$ export GOPATH=~/go
$ git clone git@github.com:iron-io/functions.git ${GOPATH}/scr/github.com/iron-io/functions
$ git clone git@github.com:denismakogon/functions-keystonemiddleware.git ${GOPATH}/scr/github.com/denismakogon/functions-keystonemiddleware
$ pushd ${GOPATH}/scr/github.com/iron-io/functions && make dep; popd
$ pushd ${GOPATH}/scr/github.com/denismakogon/functions-keystonemiddleware && make all; popd
```
