## Motivation
So far, `Scanln` looks like not very usable in terms of Kubernetes deployment, better option would be to use `flag` package instead in conjunction with `goware/urlx` package to parse supplied user input.

## Expectations

During Kubernetes deployment, I would like to supply cli arguments to the Go static binary, as in standard Unix command, e.g:

```yml
containers:
    - image: some.go.image:0.1
      command: [ "/bin/sh", "-c" ]
      args: [ "/mygobinary -server=https://some.server.com:9092/connect" ]
```

## Result
```golang
» go run main.go -server=https://some.server.com:9092/connect
Server: https://some.server.com:9092/connect
Scheme: https
Hostname: some.server.com
Port: 9092
```

## Conclusion

Although it would be better to supply such parameters using environment variables, but this method can also be utilised as more verbose option to visualy define binary options / dependencies.

As a bonus, `flag` and `os` packages can be combined together and `flag` input defined with default value comming from environment variable.