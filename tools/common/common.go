package common

import (
	"blg/blg/model"
	"path/filepath"
	"runtime"

	"gorm.io/gorm"
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

func FindCategoryId(db *gorm.DB, category string) (int, error) {
	dbcategory := &model.ArticleCategory{}
	if err := db.Model(&model.ArticleCategory{}).Where("categoryname = ?", category).First(category).Error; err != nil {
		return -1, err
	}

	return dbcategory.Id, nil
}
