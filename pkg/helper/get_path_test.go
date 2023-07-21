package helper

import (
	"path"
	"path/filepath"
	"runtime"
)

// GetPath 获取路径
/* @description
 * @since 2023/2/205:56
 * @param dir 路径
 * @param filename 文件名
 * @return 根目录的dir或filename的拼接的绝对路径
 *  */
func GetPath(dir, filename string) (string, string) {
	_, workdir, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(workdir)) // 获取当前工作目录
	dirPath := path.Dir(root + dir)

	var filePath string
	// 如果文件名为空, 则该函数为获取目录
	if filename == "" {
		filePath := path.Join(dirPath)
		return dirPath, filePath
	}
	filePath = path.Join(dirPath, filename)

	// 对操作系统进行判断, 如果是Linux则传入的路径全部变为斜杠/
	if runtime.GOOS == "linux" {
		filePath = filepath.FromSlash(filePath)
		dirPath = filepath.FromSlash(dirPath)
		return filePath, dirPath
	}
	filePath = filepath.ToSlash(filePath)
	dirPath = filepath.ToSlash(dirPath)
	return filePath, dirPath
}
