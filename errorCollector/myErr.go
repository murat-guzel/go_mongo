package errorCollector

type myErr struct {
	TIME      string
	err       error
	ErrString string
	Caller    string
}

func (ec myErr) Write(p []byte) (n int, err error) {
	return len(p), err
}
