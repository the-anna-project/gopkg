// Package gopkg provides the string representation of the caller's package as
// being imported.
package gopkg

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-stack/stack"
)

var once sync.Once
var reference string

func init() {
	once.Do(func() {
		reference = lookup()
	})
}

// String returns the package name of the caller as used to import the caller's
// package.
func String() string {
	return reference
}

func lookup() string {
	caller := stack.Caller(5)

	filePath := fmt.Sprintf("%#s", caller)
	funcName := fmt.Sprintf("%+n", caller)
	i := pkgIndex(filePath, funcName)
	j := srcIndex(filePath[:i])
	pkg := filePath[j:i]

	return pkg
}

func pkgIndex(file, funcName string) int {
	sep := string(os.PathSeparator)
	i := len(file)
	for n := strings.Count(funcName, sep) + 1; n > 0; n-- {
		i = strings.LastIndex(file[:i], sep)
		if i == -1 {
			i = -len(sep)
			break
		}
	}
	return i
}

func srcIndex(pkg string) int {
	sep := "/src/"
	return strings.LastIndex(pkg, sep) + len(sep)
}
