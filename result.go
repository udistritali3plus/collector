package collector

import (
	"bytes"
	"fmt"
)

type Result []map[string]string

func (r Result) String() string {

	var buffer bytes.Buffer

	for _, fields := range r {

		for field, value := range fields {
			f := fmt.Sprintf("%s: %s\n", field, value)
			buffer.WriteString(f)
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
