package blockchair

import "net/http"

// Error data structure describing the error.
type Error struct {
	// Hash wrong hash.
	Hash *string
	// MainErr standard error msg.
	MainErr error
	// ExecErr error that occurred during the operation.
	ExecErr error
	// Response http response.
	Response *http.Response
}

// Error compatibility with error interface.
func (e Error) Error() string {
	return e.MainErr.Error()
}

// NewError creates a new Error instance.
func NewError(errorMain error, errorExec error, response *http.Response, hash *string) *Error {
	if errorMain == nil {
		return nil
	}

	return &Error{
		MainErr:  errorMain,
		ExecErr:  errorExec,
		Response: response,
		Hash:     hash,
	}
}

// err1 build error helper.
func (c *Client) err1(errorMain error) error {
	return NewError(errorMain, nil, nil, nil)
}

// err2 build error helper.
func (c *Client) err2(errorMain error, errorExec error) error {
	return NewError(errorMain, errorExec, nil, nil)
}

// err3 build error helper.
func (c *Client) err3(errorMain error, errorExec error, response *http.Response) error {
	return NewError(errorMain, errorExec, response, nil)
}

// err4 build error helper.
func (c *Client) err4(errorMain error, response string) error {
	return NewError(errorMain, nil, nil, &response)
}
