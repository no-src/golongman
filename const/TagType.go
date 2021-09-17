package _const

// TagType 单词标签
type TagType string

const (
	TagS1 TagType = "S1" //最常见的1000个口语词汇
	TagS2 TagType = "S2" //1001～2000最常见的口语词汇
	TagS3 TagType = "S3" //2001～3000最常见的口语词汇
	TagW1 TagType = "W1" //最常见的1000个书面词汇
	TagW2 TagType = "W2" //1001～2000最常见的书面词汇
	TagW3 TagType = "W3" //2001～3000最常见的书面词汇
)
