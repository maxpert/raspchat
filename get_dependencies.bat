echo @off
set GOPATH=%cd%

echo "Installing govend"
go get github.com/govend/govend

echo "Installing Go packages"
pushd src
..\bin\govend -v
popd

echo "Installing NPM pacakges"
npm install