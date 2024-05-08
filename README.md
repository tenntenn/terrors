# tenntenn/terrors

[![pkg.go.dev][gopkg-badge]][gopkg]

## terrors.Return

Return wraps the error to be returned to the caller.

```go
func Get(id string) (_ int, rerr error) {
	defer terrors.Return(&rerr, func(err error) error {
		return errors.Errorf("Get(%q): %w", id, err)
	})
	
	// ...
	
	return 100, nil
}
```

terrors.Return is based on [derrors.Wrap](https://cs.opensource.google/go/x/pkgsite/+/master:internal/derrors/derrors.go;l=240).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/tenntenn/terrors
[gopkg-badge]: https://pkg.go.dev/badge/github.com/tenntenn/terrors?status.svg
