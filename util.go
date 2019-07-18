package Taskutil

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func Check_lastline_str(filename string, finish_mark string)(strExi bool){
	strExi = false
	file, _ := os.Open(filename)
	defer file.Close()
	ne, _ := file.Seek(0, 2)
	if ne > 100{
		_, _ = file.Seek(ne-100, 0)
	}else{
		_, _ = file.Seek(0, 0)
	}

	INreader := bufio.NewReader(file)
	AllLine_byte, _ := ioutil.ReadAll(INreader)
	AllLine := strings.Split(string(AllLine_byte), "\n")
	if len(AllLine) > 1{
		Last_str := AllLine[len(AllLine)-2]
		if Last_str == finish_mark{
			strExi = true
		}
	}

	return
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CheckSHFinishStatus(FinishStr, filename string)(bool){
	//filename := ScriptPath + ".sign"
	exit_file, _ := PathExists(filename)

	if exit_file == true{
		return Check_lastline_str(filename, FinishStr)
	}

	return false
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}