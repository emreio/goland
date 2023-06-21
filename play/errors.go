package main

func ReturnError() error {
	panic(MyError{})
}

type MyError struct {
}

func (err MyError) Error() string {
	return "hata oldu"
}
