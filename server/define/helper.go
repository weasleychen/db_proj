package define

import (
	"os"
	"path/filepath"
	"strings"
)

func getCurrentFilepath() string {
	path, _ := filepath.Abs(os.Args[0])
	array := strings.Split(path, "/")
	return strings.Join(array[0:len(array)-3], "/")
}
