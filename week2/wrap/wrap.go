package wrap

import (
	"database/sql"
	"errors"
	"fmt"
)

func biz0() error {
	return sql.ErrNoRows
}

func biz1() error {
	if err := biz0(); err != nil {
		return fmt.Errorf("biz0 error: %w", err)
	}

	return nil
}

func biz2() error {
	if err := biz1(); err != nil {
		return fmt.Errorf("biz1 error: %w", err)
	}

	return nil
}

func Way1() {
	err := biz2()

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("biz2 error: ", err)
	}
}
