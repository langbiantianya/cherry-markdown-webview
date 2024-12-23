package file

import (
	"log"
)

type File struct {
	Mime        string
	DisplayName string
	Pattern     string
	Name        string
	Path        string
	Bytes       []byte
	StrData     string
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
			_, ok := <-filech
			if ok {
				close(filech)
			}
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
