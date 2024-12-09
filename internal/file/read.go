package file

import (
	"io"
	"os"
)

func ReadFileToByteArray(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadFile(filePath string) (*File, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {

		return nil, err
	}
	bytes, err := ReadFileToByteArray(filePath)
	if err != nil {

		return nil, err
	}
	return &File{
		Name:  fileInfo.Name(),
		Path:  filePath,
		Bytes: bytes,
	}, nil
}
