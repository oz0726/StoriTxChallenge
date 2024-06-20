package input

type FileReaderPort interface {
	ReadFile(route string)
}
