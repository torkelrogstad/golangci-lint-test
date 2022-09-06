package main

import "errors"



func main() {
	_ = returnError()
}

func returnError() error {
	return errors.New("i'm an error")
}
