package model

import "time"

type IssueToday struct {
	//期号
	IssueNumber string
	//创建时间
	CreateTime time.Time
}
