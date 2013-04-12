package cascadingerror

import (
	"fmt"
	"runtime"
	"strings"
)

type CascadingError struct {
	What  string
	Where string
	Child *CascadingError
}

func New(what string, child interface{}) *CascadingError {
	var where string
	var c *CascadingError

	if pc, file, line, ok := runtime.Caller(1); ok {
		parts := strings.Split(file, "/")
		file = parts[len(parts)-1]

		where = fmt.Sprintf("%v (%v) %v", file, line, runtime.FuncForPC(pc).Name())
	}

	switch child.(type) {
	case *CascadingError:
		c = child.(*CascadingError)
	case error:
		c = &CascadingError{child.(error).Error(), "unknown source", nil}
	case nil:
		// do nothing
	default:
		c = &CascadingError{fmt.Sprintf("%v", child), "unknown source", nil}
	}

	return &CascadingError{
		what, where, c,
	}
}

func (e CascadingError) Error() string {
	r := fmt.Sprintf("%v:\t %v", e.Where, e.What)

	if e.Child != nil {
		r += "\n" + e.Child.Error()
	}

	return r
}
