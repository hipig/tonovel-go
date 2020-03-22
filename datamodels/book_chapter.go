package datamodels

type Chapter struct {
	Title string 	`json:"title" comment:"章节标题`
	ChapterURL string 	`json:"chapter_url" comment:"章节链接`
	DetailURL string 	`json:"detail_url" comment:"详情链接`
	Source string	`json:"source" comment:"搜索结果来源"`
}
