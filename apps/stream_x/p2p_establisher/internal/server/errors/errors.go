package errors

type HTTPError struct {
	Err  string `json:"err"`
	Desc string `json:"desc"`
}

func New(err error, desc ...string) *HTTPError {
	description := ""
	if len(desc) > 0 {
		description = desc[0]
	}

	return &HTTPError{err.Error(), description}
}
