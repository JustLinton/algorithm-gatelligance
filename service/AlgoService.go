package service

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"gatelligance_algo/entity"
	"gatelligance_algo/utils"

	"github.com/jinzhu/gorm"
)

func CheckLinkTransaction(db *gorm.DB, uuid string) (string, string, string) {
	var transaction []entity.LinkTransaction
	db.Find(&transaction, "id=?", uuid)

	if len(transaction) == 0 {
		fmt.Printf("transcation: not found\n")
		return "-1", "-1", "nil"
	}

	return transaction[0].Progress, transaction[0].Status, transaction[0].Output
}

func CreateLinkTransaction(db *gorm.DB, videoLink string, owner string, serverID int, err *error) string {

	uuid := utils.GenerateUUID()

	var nt = new(entity.Transaction)
	nt.ID = uuid
	nt.Type = "1"
	nt.CreatedAt = time.Now()
	nt.Owner = owner
	nt.Server = serverID
	nt.Avatar = "https://linton-pics.oss-cn-beijing.aliyuncs.com/avatars/dog1.jpg"
	nt.Title = "任务 " + uuid
	db.Create(nt)

	var nlt = new(entity.LinkTransaction)
	nlt.ID = uuid
	nlt.Progress = "10"
	nlt.Status = "0"
	nlt.VideoLink = videoLink
	db.Create(nlt)

	go getAudioText(videoLink, uuid, db)

	return uuid
}

func getAudioText(link string, uuid string, db *gorm.DB) {

	var transaction []entity.LinkTransaction
	db.Find(&transaction, "id=?", uuid)

	if len(transaction) == 0 {
		fmt.Printf("transcation: not found\n")
		return
	}

	// command := exec.Command("python3", "./pyScripts/mainProcess.py", uuid, link)

	fileName := "./tmp/" + uuid + ".mp3"

	//......音频获取

	db.Delete(transaction[0])
	transaction[0].Progress = "15"
	transaction[0].Status = "2"
	db.Create(transaction[0])

	command := exec.Command("python3", "./pyScripts/dnld.py", fileName, link)
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	err := command.Run()

	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		db.Delete(transaction[0])
		transaction[0].Output = command.Stdout.(*bytes.Buffer).String()
		transaction[0].Status = "-2"
		db.Create(transaction[0])
		return
	}

	//......文本提取

	db.Delete(transaction[0])
	transaction[0].Progress = "55"
	transaction[0].Status = "3"
	db.Create(transaction[0])

	command = exec.Command("python3", "./pyScripts/xfr.py", fileName)
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	err = command.Run()

	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		db.Delete(transaction[0])
		transaction[0].Output = command.Stdout.(*bytes.Buffer).String()
		transaction[0].Status = "-3"
		db.Create(transaction[0])
		return
	}

	//......文本摘要

	utils.CreateTxtFileAtTmp(uuid, command.Stdout.(*bytes.Buffer).String())

	db.Delete(transaction[0])
	transaction[0].Progress = "75"
	transaction[0].Status = "4"
	db.Create(transaction[0])

	command = exec.Command("python3", "-u", "../../pythonproject/Text_sum/interface.py", "-name", "./tmp/"+uuid+".txt")
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	err = command.Run()

	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		db.Delete(transaction[0])
		transaction[0].Output = command.Stdout.(*bytes.Buffer).String()
		transaction[0].Status = "-4"
		db.Create(transaction[0])
		return
	}

	//......事务完成

	db.Delete(transaction[0])
	transaction[0].Progress = "100"
	transaction[0].Status = "1"
	transaction[0].Output = command.Stdout.(*bytes.Buffer).String()
	db.Create(transaction[0])

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
