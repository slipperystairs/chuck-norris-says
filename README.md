# Chuck Norris Says
A simple CLI that will make an HTTP request to the Chuck Norris API based on the inputs given by the user. The HTTP request will return a JSON object, but only the "value" is displayed to the end user.

Go version 1.22.2 is a minimum requirement for this CLI to work. You could edit the `go.mod` file to use your specified version (unless one of the import modules requires a specific version)

To install Go visit: https://go.dev/doc/install

## Cloning/Building/Executing the CLI
```
git clone git@github.com:slipperystairs/chuck-norris-says.git
cd chuck-norris-says

go build -o chuck-norris-says main.go
./chuck-norris-says
```
