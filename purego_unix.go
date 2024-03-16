//go:build !cgoopenh264 && (darwin || freebsd || linux)

package openh264

import "github.com/ebitengine/purego"

func openLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
func closeLibrary(handle uintptr) error {
	return purego.Dlclose(handle)
}
