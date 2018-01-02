## Motivation
Looking at the previous example `002_Parsed_user_input` while using "net/url" package, it required to do additional / manual work in order to ensure, that all URL parts were defined / and or extracted from the supplied user input.

After searching documentation and "the Internets", I've found one package, which does exactly what I needed: [goware/urlx](https://github.com/goware/urlx).

## Expectations

As in `002_Parsed_user_input` example, I expect that all the parts needed as described in `RFC3986` can be extracted and later on be used in safe manner inside Go static binary.

## Result

```golang
» go run main.go
Enter service URL in the FQDN:port/path format: some.server.com:8081/some-path
Parsed data follows...
Schema: http
Host: some.server.com
Port: 8081
Path: /some-path
```

## Conclusion

The [goware/urlx](https://github.com/goware/urlx) indeed does fit for the job required and handled most used cases in the way I would expect from the Go package. Although, validity of data should be handled as a separate case:

```golang
» go run main.go
Enter service URL in the FQDN:port/path format: httpo://some.server:1/?param=1
Parsed data follows...
Schema: httpo
Host: some.server
Port: 1
Path: /
```

Which suggests to use `ParseWithDefaultScheme` func, which is also available in `urlx` package, to hint the parser about URL scheme.

Also, it is worth to look at `urlx` [test suite](https://github.com/goware/urlx/blob/master/urlx_test.go), which covers pretty much every edge case.