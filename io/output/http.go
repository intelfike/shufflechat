package output

import (
	"io"
)

func WriteString(w io.Writer, s string) {
	w.Write([]byte(s))
}
