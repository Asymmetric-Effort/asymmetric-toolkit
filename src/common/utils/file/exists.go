package file

import "os"

func Exists(filename string,isFile bool) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if isFile {
		return !info.IsDir()
	}
	return info.IsDir()
}