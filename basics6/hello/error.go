package hello

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
