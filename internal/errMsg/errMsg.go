package errmsg

import "errors"

var ErrWrongURL error = errors.New("wrong url")
var ErrFailedCreateShort error = errors.New("failer to create short link")
var ErrFailedConvertStr error = errors.New("failer convert str to int")
var ErrFailedGetLink error = errors.New("failed to get link")
var ErrFailedCreateLink error = errors.New("failed to create link")