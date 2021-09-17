package _html

import (
	"golang.org/x/net/html"
)

// Element Dom元素
type Element struct {
	innerNode *html.Node
}

func (element *Element) GetText() (text string) {
	text = element.innerNode.FirstChild.Data
	return text
}

func (element *Element) GetAttrOrDefault(attr string) (attrValue string) {
	attrs := element.innerNode.Attr
	for _, attrItem := range attrs {
		if attrItem.Key == attr {
			attrValue = attrItem.Val
			return attrValue
		}
	}
	return attrValue
}

func (element *Element) FirstChild() *Element {
	firstChild := &Element{
		innerNode: element.innerNode.FirstChild,
	}
	return firstChild
}
