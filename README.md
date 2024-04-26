# tenntenn/terrors

[![pkg.go.dev][gopkg-badge]][gopkg]

## terrors.Return

Return wraps the error to be returned to the caller.

```go
func f(id string) (rerr error) {
	defer terrors.Return(&rerr, func(err error) error {
		return errors.Errorf("f(%q)", err, id)
	})
	
	// ...
	
	return nil
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/tenntenn/terrors
[gopkg-badge]: https://pkg.go.dev/badge/github.com/tenntenn/terrors?status.svg
