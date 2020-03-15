package datamodels

type BookContent struct {
	Title string 	`json:"title" comment:"章节标题`
	Text string 	`json:"text" comment:"章节正文`
	DetailURL string 	`json:"detail_url" comment:"详情链接`
	PreviousURL string 	`json:"previous_url" comment:"章节链接`
	NextURL string 	`json:"next_url" comment:"章节链接`
	Source string	`json:"source" comment:"搜索结果来源"`
}
