package chlog

import (
	"fmt"
	"strings"
	"time"
)

type TextFormatter struct {
	IgnoreBasicFields bool
}

func (tf *TextFormatter) Format(e *Entry) error {
	if !tf.IgnoreBasicFields {
		e.Buffer.WriteString(fmt.Sprintf("%s %s ", e.Time.Format(time.RFC3339), LevelNameMapping[e.Level]))

		if e.File != "" {
			short := e.File
			if n := strings.LastIndex(e.File, "/"); n > 0 {
				short = e.File[n+1:]
			}
			// for i := 0; i < len(e.File)-1; i++ {
			// 	if e.File[i] == '/' {
			// 		short = e.File[i+1:]
			// 		break
			// 	}
			// }
			e.Buffer.WriteString(fmt.Sprintf("%s:%d", short, e.Line))
		}

		e.Buffer.WriteString(" ")
	}

	switch e.Format {
	case FmtEmptySeparate:
		e.Buffer.WriteString(fmt.Sprint(e.Args...))
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}

	e.Buffer.WriteString("\n")
	return nil
}
