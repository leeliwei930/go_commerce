package validators

import (
	"fmt"
	"time"
)

func IsValidDate() func(value string) error {
	return func(value string) error {
		_, parseError := time.Parse("2006-01-02 15:04:05", value)
		if parseError != nil {
			return fmt.Errorf("datetime format must be based: %s", "2006-01-02 15:04:05")
		}
		return nil
	}
}
