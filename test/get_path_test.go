package test

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

// 测试获取路径是否正确
func TestGetPath(t *testing.T) {
	workPath, _ := GetPath("/test/", "")
	t.Log("workPath:", workPath)
	testPath, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	sysType := runtime.GOOS
	// 如果是Linux环境, 则将路径分割符替换成/,并进行比对
	if sysType == "linux" {
		testPath = filepath.FromSlash(testPath)
		if workPath != testPath {
			t.Fatal("获取根目录失败:" + err.Error())
		}
	}
	testPath = filepath.ToSlash(testPath)

	t.Log("testPath:", testPath)
	t.Log("workPath:", workPath)
	if workPath != testPath {
		t.Fatal("获取根目录失败:" + err.Error())
	}

	t.Log("获取根目录成功")
}

func GetPath(dir, filename string) (string, string) {
	_, workdir, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(workdir)) // 获取当前工作目录
	dirPath := path.Dir(root + dir)     // 获取配置文件的目录

	var filePath string
	if filename == "" {
		filePath := path.Join(dirPath) // 获取配置文件
		return dirPath, filePath
	}
	filePath = path.Join(dirPath, filename) // 获取配置文件
	return filePath, dirPath
}
