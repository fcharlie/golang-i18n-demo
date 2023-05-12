# todo

```go
func detectContentType(data []byte) string {
	mime := mimetype.Detect(data)
	m := mime.String()
	if strings.HasPrefix(m, "image/") {
		return contentType
	}
	// Avoid browser xss attacks.
	// Example: text/html. ⚠️
	if strings.HasPrefix(m, "text/") {
		index := strings.Index(m, ";")
		if index != -1 {
			contentType = "text/plain" + m[index:]
		}
		return m
	}
	return "application/octet-stream"
}
```
