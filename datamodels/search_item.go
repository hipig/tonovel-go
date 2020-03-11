package datamodels

type SearchItem struct {
	Name string 	`json:"name" comment:"搜索结果名称`
	Author string 	`json:"author" comment:"搜索结果作者"`
	Cover string 	`json:"cover" comment:"搜索结果封面"`
	Category string 	`json:"category" comment:"搜索结果分类"`
	NewChapter string 	`json:"new_chapter" comment:"搜索结果最新章节"`
	URL string 	`json:"url" comment:"搜索结果链接"`
	Source string	`json:"source" comment:"搜索结果来源"`
}
