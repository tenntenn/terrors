package terrors

// Return wraps the error to be returned to the caller.
// The wrapped error is set to the variable via the given pointer.
// Basically, Return is called with a defer statement.
//
//	func Get(id string) (_ int, rerr error) {
//		defer terrors.Return(&rerr, func(err error) error {
//			return 0, errors.Errorf("Get(%q):%w", id, err)
//		})
//		
//		// ...
//		
//		return 100, nil
//	}
//
func Return(errp *error, wrap func(err error) error) {
	if *errp != nil {
		*errp = wrap(*errp)
	}
}
