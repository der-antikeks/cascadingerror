package cascadingerror_test

import (
	"fmt"
	"github.com/der-antikeks/cascadingerror"
)

func ExampleNew() {
	err := cascadingerror.New(
		"the database disappeared",
		nil,
	)

	if err != nil {
		fmt.Print(err)
	}
	// Output: error_test.go (12) github.com/der-antikeks/cascadingerror_test.ExampleNew:	 the database disappeared
}

func ExampleCascadingError_Error() {
	child := cascadingerror.New("no network connection", nil)
	parent := cascadingerror.New("the database disappeared", child)

	if parent != nil {
		fmt.Print(parent.Error())
	}
	// Output: 
	// error_test.go (22) github.com/der-antikeks/cascadingerror_test.ExampleCascadingError_Error:	 the database disappeared
	// error_test.go (21) github.com/der-antikeks/cascadingerror_test.ExampleCascadingError_Error:	 no network connection
}
