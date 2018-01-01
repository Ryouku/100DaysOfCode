## Motivation
I wanted to explore a bit possibilities around user input, using command line utility, which would be helpful while creating static Go binaries.

Such binaries must act upon user input, like:
* connect to specific server
* port
* api end-point
* etc

Also, all parameters must be logged to the stdout.

## Result
```golang
» go run main.go
Enter server address: some.server.com
Enter server port: 32035
Enter API path: service-name
http://some.server.com:32035/service-name
```
