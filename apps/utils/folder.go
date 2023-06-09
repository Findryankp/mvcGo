package utils

import (
	"fmt"
	"os"
)

func FolderCreate() {
	folderName := []string{"controllers", "models", "helpers", "configs", "routes", "middlewares"}
	for _, v := range folderName {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			// Create folder
			err = os.Mkdir(v, 0755)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
