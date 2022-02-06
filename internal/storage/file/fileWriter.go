package file

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type FileWriter struct {
}

func NewWriter() *FileWriter {
	return &FileWriter{}
}

func (fw *FileWriter) Write(dirName string, fileName string, data []byte) error {
	currPath, err := os.Getwd()
	if err != nil {
		return err
	}

	filePathName := path.Join(currPath, dirName, fileName)
	log.Println("Filename to be written to: ", filePathName)

	dir := path.Dir(filePathName)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePathName, data, 0644); err != nil {
		return err
	}

	return nil
}
