// +build windows

package pinentry

import "github.com/itsonlycode/pinentry/gpgconf"

// GetBinary always returns pinentry.exe
func GetBinary() string {
	if p, err := gpgconf.Path("pinentry"); err == nil && p != "" {
		return p
	}
	return "pinentry.exe"
}
