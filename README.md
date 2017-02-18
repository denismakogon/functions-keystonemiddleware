# functions-keystonemiddleware
OpenStack Identity (Keystone) middleware for IronFunctions


## Building a middleware

Use following commands:
```bash
$ export GOPATH=~/go

$ export MIDDLEWARE_PATH=${GOPATH}/src/github.com/denismakogon/functions-keystonemiddleware

$ git clone -b master git@github.com:denismakogon/functions-keystonemiddleware.git ${MIDDLEWARE_PATH}

$ pushd ${MIDDLEWARE_PATH} && make dep build; popd
```
