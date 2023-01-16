package model

type PicWord struct {
	IssueToday
	
	//主键Id
	PicWordId int64

	//单词
	WordName string

	//图片地址
	WordPic string

	//单词详情地址
	WordUrl string
}
