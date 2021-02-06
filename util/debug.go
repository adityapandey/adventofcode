package util

import "fmt"

func Indent(d int) string {
	return fmt.Sprintf(fmt.Sprintf("%%%ds", d), "")
}
