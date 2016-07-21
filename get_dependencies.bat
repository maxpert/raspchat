set GOPATH=%cd%
echo "Installing Go packages"

go get github.com/Azure/go-autorest/autorest
go get golang.org/x/net/context
go get golang.org/x/text
go get github.com/speps/go-hashids
go get github.com/gorilla/websocket
go get github.com/julienschmidt/httprouter
go get github.com/boltdb/bolt/...
go get gopkg.in/natefinch/lumberjack.v2
go get github.com/googollee/go-gcm
go get github.com/Azure/azure-sdk-for-go/management


pushd src/github.com/speps/go-hashids
git checkout -q master
git checkout -q tags/v1.0.0
popd 

pushd src/github.com/gorilla/websocket
git checkout -q master
git checkout -q tags/v1.0.0
popd

pushd src/github.com/julienschmidt/httprouter
git checkout -q master
git checkout -q tags/v1.1
popd

pushd src/github.com/boltdb/bolt
git checkout -q master
git checkout -q tags/v1.2.1
popd

go build -o sibte.so