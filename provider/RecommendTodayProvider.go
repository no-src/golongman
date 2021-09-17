package provider

import (
	"github.com/no-src/golongman/html"
	"github.com/no-src/golongman/http"
	"github.com/no-src/golongman/model"
	"strings"
	"time"
)

// RecommendTodayProvider 每日推荐单词抓取器
type RecommendTodayProvider struct {
	//抓取地址
	SpiderUrl string
	//单词名称抓取规则
	WordNameRule string
	//例句抓取规则
	ExampleSentenceRule string
	//单词详情地址抓取规则
	WordUrlRule string
	//热门主题列表规则
	HotTopicListRule string
	//看图识词规则
	PictureWordRule string
	//某一主题规则
	TopicRule string
	//某一主题单词规则
	TopicWordsRule string
}

func NewRecommendTodayProvider() (pro *RecommendTodayProvider) {
	pro = &RecommendTodayProvider{}
	pro.SpiderUrl = "https://www.ldoceonline.com/"
	pro.WordNameRule = "#wotd .title_entry a"
	pro.ExampleSentenceRule = "#wotd .ldoceEntry a"
	pro.WordUrlRule = "#wotd .title_entry a"
	pro.HotTopicListRule = "#hot_topics li a"
	pro.PictureWordRule = ".pictures a"
	pro.TopicRule = "#tcotw .box_title span"
	pro.TopicWordsRule = "#tcotw li a"
	return pro
}

// GetRecommendToday 获取今日推荐信息
func (pro *RecommendTodayProvider) GetRecommendToday() (recommend *model.RecommendToday, err error) {
	recommend = &model.RecommendToday{}
	body, err := _http.Get(pro.SpiderUrl)
	defer body.Close()
	if err != nil {
		return nil, err
	}
	docHtml, err := _html.NewHtmlDocument(body)
	if err != nil {
		return nil, err
	}
	//公共字段
	issueNumber := time.Now().Format("20060102")
	createTime := time.Now()

	//推荐单词
	word := &model.TodayWord{}
	word.IssueNumber = issueNumber
	word.CreateTime = createTime
	word.WordName = docHtml.GetText(pro.WordNameRule)
	word.ExampleSentence = docHtml.GetText(pro.ExampleSentenceRule)
	word.WordUrl = docHtml.GetAttrOrDefault(pro.WordUrlRule, "href")
	recommend.TodayWord = word

	//热门主题列表
	var hotTopicList []*model.Topic
	hotTopicEles := docHtml.GetElements(pro.HotTopicListRule)
	for _, topicEle := range hotTopicEles {
		topic := &model.Topic{
			TopicName: topicEle.GetText(),
			TopicUrl:  topicEle.GetAttrOrDefault("href"),
		}
		topic.IssueNumber = issueNumber
		topic.CreateTime = createTime
		hotTopicList = append(hotTopicList, topic)
	}
	recommend.HotTopics = hotTopicList

	//看图识词
	var picWords []*model.PicWord
	picWordEles := docHtml.GetElements(pro.PictureWordRule)
	for _, picEle := range picWordEles {
		picWord := &model.PicWord{
			WordUrl: picEle.GetAttrOrDefault("href"),
			WordPic: picEle.FirstChild().GetAttrOrDefault("src"),
		}
		if len(picWord.WordUrl) > 0 {
			splitArr := strings.Split(picWord.WordUrl, "/")
			if len(splitArr) > 0 {
				picWord.WordName = splitArr[len(splitArr)-1]
			}
		}
		picWord.IssueNumber = issueNumber
		picWord.CreateTime = createTime
		picWords = append(picWords, picWord)
	}
	recommend.PicWords = picWords

	//指定主题及相关单词
	topic := &model.Topic{}
	topic.TopicName = docHtml.GetText(pro.TopicRule)
	topic.IssueNumber = issueNumber
	topic.CreateTime = createTime
	topicWordEles := docHtml.GetElements(pro.TopicWordsRule)
	for _, topicWordEle := range topicWordEles {
		word := &model.TodayWord{
			WordName: topicWordEle.GetText(),
			WordUrl:  topicWordEle.GetAttrOrDefault("href"),
		}
		word.IssueNumber = issueNumber
		word.CreateTime = createTime
		topic.Words = append(topic.Words, word)
	}
	recommend.Topic = topic

	return recommend, nil
}
