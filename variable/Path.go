package variable

import "path/filepath"

var (
	PATH string
	TEMP string
	RESOURCE string
)

func Path(paths ... string) string {
	return filepath.Join(append([]string{PATH}, paths...)...)
}
func Temp(paths ... string) string {
	return filepath.Join(append([]string{TEMP}, paths...)...)
}
func Resource(paths ... string) string {
	return filepath.Join(append([]string{RESOURCE}, paths...)...)
}