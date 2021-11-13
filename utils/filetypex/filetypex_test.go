package filetypex

import "testing"

func TestFileType(t *testing.T) {
	t.Run("file type", func(t *testing.T) {
		fileType := FileType(".ppt")
		t.Log(fileType)
	})
}
