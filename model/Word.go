package model

import "github.com/no-src/golongman/const"

// Word 单词基本信息
type Word struct {
	//主键Id
	WordId int64

	//单词
	WordName string

	//语言类型
	Language _const.LanguageType
}
