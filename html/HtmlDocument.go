package _html

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

// HtmlDocument HTML文档对象
type HtmlDocument struct {
	innerDoc *goquery.Document
}

// NewHtmlDocument 新建一个HTML文档对象
func NewHtmlDocument(reader io.Reader) (htmlDoc *HtmlDocument, err error) {
	htmlDoc = &HtmlDocument{}
	innerDoc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	htmlDoc.innerDoc = innerDoc
	return htmlDoc, nil
}

// GetText 获取指定选择器中的文本内容
func (doc *HtmlDocument) GetText(selector string) (text string) {
	text = doc.innerDoc.Find(selector).Text()
	return text
}

// GetHTML 获取指定选择器中的HTML内容
func (doc *HtmlDocument) GetHTML(selector string) (html string, err error) {
	html, err = doc.innerDoc.Find(selector).Html()
	return html, err
}

// GetAttr 获取指定选择器中指定的Attr内容
func (doc *HtmlDocument) GetAttr(selector string, attr string) (attrValue string, exist bool) {
	attrValue, exist = doc.innerDoc.Find(selector).Attr(attr)
	return attrValue, exist
}

// GetAttrOrDefault 获取指定选择器中指定的Attr内容,若不存在则返回默认值
func (doc *HtmlDocument) GetAttrOrDefault(selector string, attr string) (attrValue string) {
	attrValue, exist := doc.GetAttr(selector, attr)
	if exist == false {
		attrValue = ""
	}
	return attrValue
}

// GetElements 获取指定选择器选中的元素
func (doc *HtmlDocument) GetElements(selector string) (elements []*Element) {
	nodes := doc.innerDoc.Find(selector).Nodes
	if nodes != nil {
		for _, node := range nodes {
			element := &Element{
				innerNode: node,
			}
			elements = append(elements, element)
		}
	}
	return elements
}
