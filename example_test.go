package cascadingerror_test

import (
	"fmt"
	"github.com/der-antikeks/cascadingerror"
)

func connectdb() error {
	msg := fmt.Errorf("no network connection")

	return cascadingerror.New(
		"the database disappeared",
		msg,
	)
}

func createtables() *cascadingerror.CascadingError {
	if err := connectdb(); err != nil {
		return cascadingerror.New(
			"could not create tables",
			err,
		)
	}

	return nil
}

func Example() {
	if err := createtables(); err != nil {
		fmt.Println(err)
	}
}
