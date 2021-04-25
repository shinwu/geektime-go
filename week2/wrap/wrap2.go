package wrap

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func bar0() error {
	return sql.ErrNoRows
}

func bar1() error {
	return errors.WithMessage(bar0(), "bar0 error")
}

func bar2() error {
	return errors.WithMessage(bar1(), "bar1 error")
}

func Way2() {
	err := errors.Wrap(bar2(), "bar2 error")

	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("%+v", err)
	}
}
