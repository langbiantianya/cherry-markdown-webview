package file

import (
	"log"
)

type File struct {
	DisplayName string
	Pattern     string
	Name        string
	Path        string
	Bytes       []byte
}

var filech chan *File

func AsynLoadingToRam(filePath string) {
	if len(filePath) != 0 {
		filech = make(chan *File)
		go func() {
			f, err := ReadFile(filePath)
			if err != nil {
				close(filech)
				log.Fatal(err)
			} else {
				filech <- f
			}
		}()
	} else {
		if filech != nil {
			close(filech)
		}
	}
}

func GetFile() *File {
	if filech == nil {
		return new(File)
	}
	for {
		select {
		case f, ok := <-filech:
			if !ok {
				return new(File)
			}
			defer close(filech)
			return f
		}
	}
}
