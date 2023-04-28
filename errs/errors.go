package errs

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var BasePath string = ""
var NoLines bool = false

const (
	msgError0         = "something went wrong"
	msgCouldNotFormat = "could not format error"
)

type Error interface {
	error
}

func New(params ...any) Error {
	return factory.Error(params...)
}

var factory Factory

func init() {
	factory = &errorFactory{}
}

type Factory interface {
	Error(params ...any) Error
}

type errorFactory struct{}

func (ef *errorFactory) Error(params ...any) Error {
	var err erro
	err.sep = "\n\t"
	switch len(params) {
	case 0:
		ef.error0(&err)
	case 1:
		ef.error1(&err, params[0])
	case 2:
		ef.error2(&err, params[0], params[1])
	default:
		ef.errorn(&err, params[0], params[1], params[2:])
	}
	ef.caller(&err)
	return &err
}
func (ef *errorFactory) caller(err *erro) {
	pc, file, line, ok := runtime.Caller(3)
	err.lineInfo = lineInfo{
		ProgramCounter: pc,
		Filename:       strings.TrimPrefix(file, BasePath),
		Line:           line,
		Retrieved:      ok,
	}
}
func (ef *errorFactory) error0(err *erro) {
	err.err = errors.New(msgError0)
}
func (ef *errorFactory) errorn(err *erro, error any, msg any, params []any) {
	m := msgCouldNotFormat
	switch error := error.(type) {
	case string:
		m = fmt.Sprintf(error, append([]any{msg}, params...)...)
		ef.error1(err, m)
	default:
		if msg, ok := msg.(string); ok {
			m = fmt.Sprintf(msg, params...)
		}
		ef.error2(err, error, m)
	}
}
func (ef *errorFactory) error2(err *erro, error any, msg any) {
	ef.error1(err, error)
	if m, ok := msg.(string); ok {
		err.msg = m
	} else {
		err.msg = fmt.Sprintf("%v", msg)
	}
}

func (ef *errorFactory) error1(err *erro, param any) {
	switch param := param.(type) {
	case error:
		err.err = param
	case string:
		err.err = fmt.Errorf(param)
	default:
		err.err = fmt.Errorf("%v", param)
	}
}

type lineInfo struct {
	ProgramCounter uintptr
	Filename       string
	Line           int
	Retrieved      bool
}
type erro struct {
	err      error
	msg      string
	sep      string
	lineInfo lineInfo
}

func (e *erro) Is(tgt error) bool {
	if tgt == e {
		return true
	}
	return errors.Is(e.err, tgt)
}
func (e *erro) Error() string {
	var msg string
	if er, ok := e.err.(*erro); ok {
		msg = fmt.Sprintf("(%s)%s%s", e.msg, e.sep, er.Error())
	} else {
		if e.msg != "" {
			msg = fmt.Sprintf("(%s: %s)", e.err.Error(), e.msg)
		} else {
			msg = fmt.Sprintf("(%s)", e.err.Error())
		}
	}
	msg = fmt.Sprintf("%s %s", e.line(), msg)
	return msg
}

func (e *erro) line() string {
	if NoLines {
		return e.lineInfo.Filename
	}
	return fmt.Sprintf("%s:%d ", e.lineInfo.Filename, e.lineInfo.Line)
}
