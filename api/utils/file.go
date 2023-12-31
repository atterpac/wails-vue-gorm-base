package utils

import (
	"fmt"
	"os"
)

func CreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	} 
	return nil
}
