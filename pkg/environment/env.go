package environment

import (
	"errors"
	"fmt"
)

var (
	ErrNotFoundKey = errors.New("key not found")
)

func Get[T any](key string) T {
	if !configManager.IsSet(key) {
		panic(fmt.Errorf("%w: %s", ErrNotFoundKey, key))
	}

	return configManager.Get(key).(T)
}
