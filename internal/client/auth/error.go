package auth

type ErrorAuth struct {
	Code   int
	Reason string
}

func (err *ErrorAuth) Error() string {
	return err.Reason
}
