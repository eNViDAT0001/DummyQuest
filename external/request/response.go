package request

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []*Error    `json:"errors,omitempty"`
}

func (s *Response) WithError(err *Error) *Response {
	s.Errors = append(s.Errors, err)
	return s
}

func NewSuccessResponse(data interface{}) *Response {
	res := &Response{
		Data: data,
	}

	return res
}

func NewErrResponse(errs ...error) *Response {
	result := make([]*Error, 0)
	for _, err := range errs {
		if e, ok := err.(*Error); ok {
			result = append(result, e)
		}
	}

	res := &Response{
		Errors: result,
	}

	return res
}
