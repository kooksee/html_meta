package kts

func NewHtmlPattern() *HtmlPattern {
	return &HtmlPattern{}
}

type HtmlPattern struct {
	Id      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Pattern string `json:"pattern" db:"pattern"`
}

func (t *HtmlPattern) TableName() string {
	return "html_pattern"
}
