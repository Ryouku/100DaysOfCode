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
2018/01/01 20:26:28 Entered: some.server.com
Enter server port: 32003
2018/01/01 20:26:31 Entered: 32003
Enter API path: some-service
2018/01/01 20:26:37 Entered: some-service
http://some.server.com:32003/some-service
```
