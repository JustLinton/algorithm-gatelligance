package service

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetAudioText(addr string, id string) string {

	command := exec.Command("python3", "./pyScripts/mainProcess.py", id, addr)

	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}

	err := command.Run()
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(command.Stderr.(*bytes.Buffer).String())
		return "nil"
	}
	//打印命令行的标准输出
	// fmt.Println(command.Stdout.(*bytes.Buffer).String())

	return command.Stdout.(*bytes.Buffer).String()
}

// func GetAudioText(addr string, id string) string {
// 	resp, err := http.PostForm("https://app.cupof.beer:8083/download-count",
// 		url.Values{"Count": {"33433"}, "id": {"123"}})

// 	print(addr + " " + id)

// 	if err != nil {
// 		return "nil"
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "nil"
// 	}

// 	// fmt.Println(string(body))
// 	return string(body)
// }
