package common

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	tempDir  = "./files/temps/"  //临时目录
	draftDir = "./files/drafts/" //草稿目录
	permDir  = "./files/perms/"  //永久目录
)

/*
 * 创建文件
 * 1. 当文件不存在时，创建文件，并返回文件句柄
 * 2. 当文件存在时候，返回文件句柄
 */
func OpenOrCreateFile(fileName string) (file *os.File, err error) {
	if !IsExists(fileName) {
		file, err := os.Create(fileName)
		return file, err
	} else {
		//当文件存在时，以读写的方式打开，权限是rw-rw-rw
		file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
		return file, err
	}
}

//删除文件
func DeleteFile(fileName string) error {
	if IsExists(fileName) {
		err := os.Remove(fileName)
		return err
	}
	return nil
}

//拷贝文件
func CopyFile(sPath string, tPath string) error {
	if !IsExists(sPath) {
		return errors.New("源路径不存在")
	}
	//获取源文件内容
	bytes, err := ReadFile(sPath)
	if err != nil {
		return err
	}
	//删除目标文件
	err = DeleteFile(tPath)
	if err != nil {
		return err
	}
	//创建并打开目标文件
	fi, err := OpenOrCreateFile(tPath)
	defer fi.Close()
	//将源文件的内容写入到目标文件
	err = WriteFile(tPath, bytes)
	if err != nil {
		return err
	}
	return nil
}

//移动文件
func MoveFile(sPath string, tPath string) error {
	err := CopyFile(sPath, tPath)
	if err != nil {
		return err
	}
	err = DeleteFile(sPath)
	return err
}

//获取文件的目录
func GetFileDir(fileName string) (dir string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("get file dir error")
		}
	}()
	sArr := strings.Split(fileName, "/")
	sArrLength := len(sArr)
	tStr := strings.Join(sArr[:sArrLength-2], "/")
	return tStr, nil
}

//获取文件的扩展名
func GetFileExtendName(fileName string) (extendName string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("get file extendName error")
		}
	}()
	if fileName == "" {
		return "", errors.New("file name is null")
	}
	sArr := strings.Split(fileName, "/")
	sArrLength := len(sArr)
	tStr := ""
	if sArrLength >= 2 {
		tStr = strings.Join(sArr[sArrLength-2:], "")
	} else {
		tStr = sArr[0]
	}
	fileNameInfo := strings.Split(tStr, ".")
	extendName = fileNameInfo[1]
	return extendName, nil
}

//读取文件内容
func ReadFile(fileName string) (bytes []byte, err error) {
	bytes = make([]byte, 0)
	fi, err := OpenOrCreateFile(fileName)
	defer fi.Close()
	if err != nil {
		return bytes, err
	}
	_, err = fi.Read(bytes)
	return
}

//向文件中写入内容
func WriteFile(fileName string, bytes []byte) error {
	fi, err := OpenOrCreateFile(fileName)
	defer fi.Close()
	_, err = fi.Write(bytes)
	return err
}

//拷贝文件流
func FileStreamCopy(t io.Writer, s io.Reader) error {
	_, err := io.Copy(t, s)
	return err
}

//判断文件或者文件夹是否存在
func IsExists(name string) bool {
	//获取文件状态
	_, err := os.Stat(name)
	//没有错误说明文件存在
	if err == nil {
		return false
	}
	//如果发生错误，则使用对应的方法判断是不是由于文件不存在引起的错误
	isExists := os.IsExist(err)
	return isExists
}

/*
 * 创建文件夹
 * 当文件夹不存在的时候，创建文件夹
 */
func CreateDir(path string) error {
	if !IsExists(path) {
		err := os.MkdirAll(path, 0666)
		return err
	}
	return nil
}

//获取临时文件夹
func GetTempDir() (path string, err error) {
	path, err = GetPath(tempDir)
	return
}

//获取草稿文件夹
func GetDraftDir() (path string, err error) {
	path, err = GetPath(draftDir)
	return
}

//获取永久文件夹
func GetPermDir() (path string, err error) {
	path, err = GetPath(permDir)
	return
}

//获取制定文件夹路径
func GetPath(name string) (path string, err error) {
	err = CreateDir(name)
	if err != nil {
		return "", err
	}
	//获取文件夹的绝对路径
	path, err = filepath.Abs(name)
	return
}
