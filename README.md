[![Build Status](https://dev.azure.com/mchirico/FirebaseGo/_apis/build/status/mchirico.FirebaseGo?branchName=master)](https://dev.azure.com/mchirico/FirebaseGo/_build/latest?definitionId=35&branchName=master)
[![codecov](https://codecov.io/gh/mchirico/FirebaseGo/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/FirebaseGo)


# FirebaseGo



### Checklist:

1. dockerPassword
2. [CodeCov Token](https://codecov.io/gh/mchirico)
3. No Caps in project
4. MONGO_CONNECTION_STRING
5. MONGO_DATABASE 
6. Make Azure Boards Public
7. More Cobra commands? (cobra add say)



## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


