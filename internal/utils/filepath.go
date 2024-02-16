package utils

import (
	"os"
	"strings"
)

func MakePath(prefix *string, file string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Can't get CWD")
	}
	var sb strings.Builder
	sb.WriteString(cwd)
	if cwd[len(cwd)-1] != os.PathSeparator {
		sb.WriteRune(os.PathSeparator)
	}
	if prefix != nil {
		sb.WriteString(*prefix)
		tmp := *prefix
		if tmp[len(tmp)-1] != os.PathSeparator {
			sb.WriteRune(os.PathSeparator)
		}
	}
	sb.WriteString(file)
	return sb.String()
}
