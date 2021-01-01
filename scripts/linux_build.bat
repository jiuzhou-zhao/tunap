pushd

cd /d %~dp0
cd ..

set GOARCH=amd64
set GOOS=linux
go build -o tunaps ./cmd/tunaps
go build -o tunapc ./cmd/tunapc

popd
