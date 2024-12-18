package images

import (
	"fmt"
	"os"
)

func DeleteFile(filename string) error {
	fullPath := fmt.Sprintf(filename)
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}
	return nil
}
