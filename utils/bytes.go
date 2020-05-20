package utils

import "fmt"

func FormatFileSize(fileSize int64) (size string) {
	fs := float64(fileSize)
	i := 0
	byteUnits := []string{" B", " KB", " MB", " GB", " TB", " PB", " EB", " ZB", " YB", " BB"}
	if fs >= 1024 {
		for i < len(byteUnits) && fs >= 1024 {
			fs = fs / 1024
			i++
		}
	} else {
		return fmt.Sprintf("%.2f%s", fs, " B")
	}
	return fmt.Sprintf("%.2f%s", fs, byteUnits[i])
}
