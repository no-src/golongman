package model

type Topic struct {
	IssueToday
	
	//主键
	TopicId int64
	//主题名称
	TopicName string
	//主题详情地址
	TopicUrl string
	//主题下的单词列表
	Words []*TodayWord
}
