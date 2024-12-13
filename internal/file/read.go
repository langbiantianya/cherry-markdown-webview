package file

import (
	"io"
	"os"
	"path/filepath"

	"github.com/wailsapp/mimetype"
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

func JoinAbsPath(basePath string, elem ...string) string {
	dir, _ := filepath.Split(basePath)

	return filepath.Join(dir, filepath.Join(elem...))
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

	f := File{
		Name:  fileInfo.Name(),
		Path:  filePath,
		Bytes: bytes,
	}

	mtype, err := mimetype.DetectFile(filePath)

	if err == nil {
		f.DisplayName = mtype.String()
		f.Pattern = mtype.Extension()
		f.Mime = mtype.String()
	}
	return &f, nil
}
