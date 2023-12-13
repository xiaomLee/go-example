package degrade

import (
	"apollo"
)

const (
	Disable = iota
	Enable

	//Section = "degrade"
)

// Enabled determines whether the given section key is enabled
func Enabled(section, key string) bool {
	val := apollo.GetKey(section, key).MustInt(0)
	return val == Enable
}
