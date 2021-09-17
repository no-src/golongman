package model

import "time"

// WordToday 每日单词
type WordToday struct {
	//主键
	WordOfTheDayId int64
	//单词
	WordName string
	//单词例句
	ExampleSentence string
	//单词详情地址
	WordUrl string
	//期号
	IssueNumber string
	//创建时间
	CreateTime time.Time
}
