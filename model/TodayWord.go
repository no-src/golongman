package model

// TodayWord 每日单词
type TodayWord struct {
	IssueToday
	//主键
	WordOfTheDayId int64
	//单词
	WordName string
	//单词例句
	ExampleSentence string
	//单词详情地址
	WordUrl string
}
