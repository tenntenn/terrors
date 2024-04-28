# tenntenn/terrors

[![pkg.go.dev][gopkg-badge]][gopkg]

## terrors.Return

Return wraps the error to be returned to the caller.

```go
func Get(id string) (_ int, rerr error) {
	defer terrors.Return(&rerr, func(err error) error {
		return 0, errors.Errorf("f(%q)", err, id)
	})
	
	// ...
	
	return 100, nil
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/tenntenn/terrors
[gopkg-badge]: https://pkg.go.dev/badge/github.com/tenntenn/terrors?status.svg
