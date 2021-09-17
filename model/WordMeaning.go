package model

import "github.com/no-src/golongman/const"

// WordMeaning 单词释义详情
type WordMeaning struct {
	//单词
	WordName string

	//重音描述
	WordStress string

	//单词类型（形容词，副词等）
	WordClass _const.WordClassType

	//音标
	Phonogram string

	//读音-英-本地备份
	PronunciationEN string

	//读音-美-本地备份
	PronunciationUSA string

	//读音-英-原始文件
	PronunciationSrcEN string

	//读音-美-原始文件
	PronunciationSrcUSA string

	//单词标签
	Tags []_const.TagType

	//高频词汇
	TooltipLevel string
}
