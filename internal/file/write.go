package file

import (
	"os"
)

func WriteFile(file File) error {
	if file.StrData != "" && len(file.Bytes) == 0 {
		file.Bytes = []byte(file.StrData)
	}
	return WireByteArray(file.Path, file.Bytes)
}

func WireByteArray(path string, byteArray []byte) error {
	doc, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer doc.Close()

	_, err = doc.Write(byteArray)
	if err != nil {
		return err
	}

	return nil
}
