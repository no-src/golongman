package model

// RecommendToday 今日推荐信息
type RecommendToday struct {
	//今日推荐单词
	TodayWord *TodayWord
	//今日热门主题列表
	HotTopics []*Topic
	//看图识词
	PicWords []*PicWord
	//今日推荐主题
	Topic *Topic
}
