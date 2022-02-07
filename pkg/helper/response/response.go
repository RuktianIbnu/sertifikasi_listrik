package response

// Body ...
type Body struct {
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *pagination `json:"pagination,omitempty"`
}

type pagination struct {
	CurrentPage  interface{} `json:"current_page,omitempty"`
	Limit        interface{} `json:"limit,omitempty"`
	TotalEntries interface{} `json:"total_entries,omitempty"`
}

// Format ...
func Format(code int, err error, data ...interface{}) (statusCode int, b *Body) {
	var (
		msg string
		d   interface{}

		pg = pagination{}
	)

	if err != nil {
		msg = err.Error()
	}

	if len(data) >= 1 {
		d = data[0]
	}

	b = &Body{
		Code:    code,
		Data:    d,
		Message: msg,
	}

	if len(data) > 1 {
		pg.TotalEntries = data[1]
		pg.CurrentPage = data[2]
		pg.Limit = data[3]

		b.Pagination = &pg
	}

	statusCode = code

	return
}
