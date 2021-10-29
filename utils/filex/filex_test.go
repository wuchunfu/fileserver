package filex

import "testing"

func TestFile(t *testing.T) {
	path := "/tmp/test"
	exists := FilePathExists(path)
	t.Log(exists)
}
