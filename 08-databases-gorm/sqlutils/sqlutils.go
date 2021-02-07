package sqlutils

import (
	"database/sql/driver"
	"errors"
)

type IntBool bool

// Value implements the driver.Valuer interface,
// and turns the IntBool into an integer for MySQL storage.
func (i IntBool) Value() (driver.Value, error) {
	if i {
		return 1, nil
	}
	return 0, nil
}

// Scan implements the sql.Scanner interface,
// and turns the int incoming from MySQL into an IntBool
func (i *IntBool) Scan(src interface{}) error {
	v, ok := src.(int)
	if !ok {
		return errors.New("bad int type assertion")
	}
	*i = v == 1
	return nil
}
