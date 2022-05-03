package entity

type ApiErrors struct {
	Errors []ApiError `json:"errors,omitempty"`
	Raw    string     `json:"raw,omitempty"`
}

type ApiError struct {
	Code   string         `json:"code,omitempty"`
	Status string         `json:"status,omitempty"`
	Title  string         `json:"title,omitempty"`
	Detail string         `json:"detail,omitempty"`
	Source ApiErrorSource `json:"source,omitempty"`
	Meta   ApiErrorMeta   `json:"meta,omitempty"`
}

type ApiErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
}

type ApiErrorMeta struct {
	Parameters ApiErrorParameters `json:"parameters,omitempty"`
}

type ApiErrorParameters struct {
	Definition    string `json:"definition,omitempty"`
	ExpectedClass string `json:"expectedClass,omitempty"`
	ActualClass   string `json:"actualClass,omitempty"`
}
