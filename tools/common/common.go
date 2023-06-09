package common

import (
	"path/filepath"
	"runtime"
)

func GetCurPath() string {
	// 获取当前函数的调用信息
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("无法获取调用信息")
	}

	// 获取相对路径
	relativePath := filepath.Dir(filename)

	parentPath := filepath.Join(relativePath, "..", "..")

	return parentPath
}
