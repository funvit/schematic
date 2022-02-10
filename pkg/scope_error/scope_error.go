// Package scope_error defines a scoped error.
//
// Sometimes you need a scoped error, which can help to group multiple errors
// for later simple check via errors.Is().
//
// Example:
//
//    var ErrRepoError = scope_error.New("repo error")
//
//    // some sql method returns error
//    rErr := repo.GetOrder(id)
//
//    // you just wrap it at return...
//    return ErrRepoError.WithCause(rErr)
//
//    // later check
//    if errors.Is(err, ErrRepoError) {
//        // any error wrapped via WithCause will ends here
//        log.Println("ERROR", err)
//    }
//
package scope_error

import "fmt"

type ScopeError struct {
	scope string
	cause error
}

func New(scope string) *ScopeError {
	return &ScopeError{
		scope: scope,
	}
}

func (s ScopeError) Error() string {
	if s.cause == nil {
		return s.scope
	}

	return fmt.Sprintf("%s: %s", s.scope, s.cause.Error())
}

func (s *ScopeError) Is(err error) bool {
	switch e := err.(type) {
	case *ScopeError:
		return e.scope == s.scope
	default:
		return false
	}
}

func (s *ScopeError) Unwrap() error {
	return s.cause
}

// WithCause returns a new scope error clone with cause set.
//
// Note: Is() will returns true for clone and original.
func (s *ScopeError) WithCause(err error) error {
	e := *s
	e.cause = err

	return &e
}
