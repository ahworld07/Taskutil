package Taskutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"errors"
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

func GenerateShell(cmdline, filename, finishStr string){
	if finishStr == ""{
		finishStr = "Still_waters_run_deep"
	}
	//for file in glob.glob(shell + '.*'):
	//os.remove(file)
	cmdline = strings.TrimSuffix(cmdline, "\n")
	//cmdline = fmt.Sprintf("echo -e ==========start at : `date +\"%Y-%m-%d %H-%M-%s\"` ==========\n") + cmdline
	//cmdline = cmdline + fmt.Sprintf(" && \\\necho -e ==========end at : `date +\"%Y-%m-%d %H-%M-%s\"` ==========")
	cmdline = cmdline + " && \\\necho " + finishStr + fmt.Sprintf(" >%s.sign", filename)

	var f *os.File
	if CheckFileIsExist(filename) { //如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(filename)
	}
	_, _ = io.WriteString(f, cmdline)
	w := bufio.NewWriter(f)
	w.Flush()
	f.Close()
}

func GenerateShell_PS(cmdline, filename, finishStr string){
	if finishStr == ""{
		finishStr = "Still_waters_run_deep"
	}
	//for file in glob.glob(shell + '.*'):
	//os.remove(file)
	cmdline = strings.TrimSuffix(cmdline, "\n")
	cmdline = fmt.Sprintf("echo -e ==========start at : `date +\"%Y-%m-%d %H-%M-%s\"` ==========\n") + cmdline
	cmdline = cmdline + fmt.Sprintf(" && \\\necho -e ==========end at : `date +\"%Y-%m-%d %H-%M-%s\"` ==========")
	cmdline = cmdline + " && \\\necho " + finishStr + fmt.Sprintf(" >%s.sign", filename)

	var f *os.File
	if CheckFileIsExist(filename) { //如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(filename)
	}
	_, _ = io.WriteString(f, cmdline)
	w := bufio.NewWriter(f)
	w.Flush()
	f.Close()
}

func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	defer f.Close()
	if err == nil {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}

	return err
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name,content string) {
	data :=  []byte(content)
	err := ioutil.WriteFile(name,data,0644)
	CheckErr(err)
}

//使用io.WriteString()函数进行数据的写入
func WriteWithIo(name,content string) {
	fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		CheckErr(err)
		os.Exit(2)
	}
	if  _,err := io.WriteString(fileObj,content);err == nil {
	    CheckErr(err)
	}
}
