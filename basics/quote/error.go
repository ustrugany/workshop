package quote

// - Err error is underlying error, when to wrap?
// - wrap an error to expose it to callers.
// - do not wrap an error when doing so would expose implementation details.
type ProviderError struct {
	Reason string
	Err    error
}

func (e ProviderError) Error() string {
	return e.Reason
}

func (e ProviderError) Unwrap() error {
	return e.Err
}
