package utils

import (
	"fmt"
	"os"
)

func CreateTxtFileAtTmp(fileName string, content string) {
	dstFile, err := os.Create("./tmp/" + fileName + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(content + "\n")
}
