package main

type errWrongContentType struct{ msg string }

func (e *errWrongContentType) Error() string { return e.msg }
