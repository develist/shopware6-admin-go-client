package entity

type ApiErrors struct {
	Errors []ApiError
}

type ApiError struct {
	Code   string
	Status int
	Detail string
	Source struct {
		Pointer string
	}
}
