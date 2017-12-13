echo Building Linux edition
env GOOS=linux GOARCH=amd64 go build -o builds/linux/example

echo Building Windows edition
env GOOS=windows GOARCH=amd64 go build -o builds/windows/example.exe

echo Building MacOS edition
env GOOS=darwin GOARCH=amd64 go build -o builds/macos/example
