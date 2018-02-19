package shared

import (
	"fmt"
	"os/user"
	"path"
	"strings"
)

// exports
const (
	Address     = "localhost:8080"
	TodoDaoImpl = "txt"
	fileName    = "/todo.txt"
)

// Todotxt file path of the todo.txt
var Todotxt = todoPath()

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// TodoPath export the path to the todo.txt file
func todoPath() string {
	u, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	return u.HomeDir + fileName
}
