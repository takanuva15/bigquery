package driver

import (
	"database/sql/driver"
	"fmt"
)

const maxValuerUnwrapDepth = 100

func unwrapValuer(namedValue *driver.NamedValue) error {
	value := namedValue.Value
	for depth := 0; depth < maxValuerUnwrapDepth; depth++ {
		valuer, ok := value.(driver.Valuer)
		if !ok {
			namedValue.Value = value
			return nil
		}

		unwrapped, err := valuer.Value()
		if err != nil {
			return err
		}

		value = unwrapped
	}

	return fmt.Errorf("valuer unwrap exceeded max depth %d", maxValuerUnwrapDepth)
}
