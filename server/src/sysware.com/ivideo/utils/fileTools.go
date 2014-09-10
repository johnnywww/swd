package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sysware.com/ivideo/common"
)

func ReadFileToBytes(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadFileToString reads data type 'string' from file by given path.
// It returns error when fail to finish operation.
func ReadFileToString(filePath string) (string, error) {
	b, err := ReadFileToBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadFileToStringNoLn(filePath string) (string, error) {
	str, err := ReadFileToString(filePath)
	if err != nil {
		return "", err
	}

	return TrimRightSpace(str), nil
}

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// get filepath base name
func GetFileName(file string) string {
	return path.Base(file)
}

// get filepath dir name
func GetFileDir(file string) string {
	return path.Dir(file)
}

func CreateRecursDir(path string) error {
	if IsPathExist(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

func GetFileExt(file string) string {
	return path.Ext(file)
}

// rename file name
func RenameFile(file string, to string) error {
	return os.Rename(file, to)
}

// delete file
func DeleteFile(file string) error {
	return os.Remove(file)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func IsPathExist(path string) bool {
	return IsFileExist(path)
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsPathExist(fullPath) {
			return
		}
	}
	err = errors.New(fullPath + " not found in paths")
	return
}

// get absolute filepath, based on built executable file
func FileRealPath(file string) (string, error) {
	if path.IsAbs(file) {
		return file, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, file), err
}

// get file modified time
func FileModTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// list dirs under dirPath
func GetDirChildDirs(dirPath string) ([]string, error) {
	if !IsPathExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := []string{}
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil

}

// list files under dirPath
func GetDirFiles(dirPath string) ([]string, error) {
	if !IsPathExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}

	return ret, nil

}

func GetServerConfigFileName(cmdLine string) string {
	return GetFileDir(cmdLine) + "/" + common.DEFAULT_CONFIG_FILE_NAME
}

func IsServerConfigExist(cmdLine string) bool {
	return IsFileExist(GetServerConfigFileName(cmdLine))
}
