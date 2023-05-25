$pos = $env:GOOS
$par = $env:GOARCH

$env:GOOS="darwin"
$env:GOARCH="amd64"

go build -C cmd -o ../fedproxy

$env:GOOS=$pos
$env:GOARCH=$par
