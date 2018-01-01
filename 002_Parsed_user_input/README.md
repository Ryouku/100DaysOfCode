## Motivation
Following previous example, I need less manual work and I like to delegate tasks to the respective Go libraries.

The "net/url" package looks like perfect fit for the job.

## Result

No schema provided

```golang
» go run main.go
Enter service URL in the FQDN:port/path format: some.server.com
2018/01/01 21:58:04 URL scheme not provided!
exit status 1
```

Schema provided

```golang
» go run main.go
Enter service URL in the FQDN:port/path format: https://some.server.com/somepath
Parsed data follows...
Schema: https
Host: some.server.com
Port:
URI: /somepath
Reassembled input: https://some.server.com/somepath
```

## Aftermath
Although, "net/url" does provide convenient methods to access various parts of parsed raw url input string,
but for sure it lacks some needed functionality, like explicitly setting default schema, port or working
with localhost.

Thus, it means that the Gopher should handle such edge cases by him self.
