package filetypex

func FileType(fileType string) string {
	switch fileType {
	case ".pdf":
		return "pdf"
	case ".doc", ".docx":
		return "docx"
	case ".ppt", ".pptx":
		return "pptx"
	case ".xls", ".xlsx":
		return "xlsx"
	case ".mid", ".mp3", ".m4a", ".ogg", ".flac", ".wav", ".amr", ".aac", ".aiff":
		return "audio"
	case ".mp4", ".m4v", ".mkv", ".webm", ".mov", ".avi", ".wmv", ".mpg", ".flv", ".3gp":
		return "video"
	case ".jpg", ".jp2", ".png", ".gif", ".webp", ".cr2", ".tif", ".bmp", ".jxr", ".psd", ".ico", ".heif", ".dwg":
		return "image"
	case ".zip", ".tar", ".tar.gz", ".rar", ".gz", ".bz2", ".7z", ".xz", ".zst":
		return "zip"
	case ".epub", ".exe", ".swf", ".rtf", ".eot", ".ps", ".sqlite", ".nes", ".crx", ".cab", ".deb", ".ar", ".Z", ".lz", ".rpm", ".elf", ".dcm", ".iso", ".macho":
		return "archive"
	default:
		return "txt"
	}
}
