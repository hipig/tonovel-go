package datamodels

type BookInfo struct {
	Name string 	`json:"name" comment:"小说名称`
	Author string 	`json:"author" comment:"小说作者"`
	Cover string 	`json:"cover" comment:"小说封面"`
	Category string 	`json:"category" comment:"小说分类"`
	Description string	`json:"description" comment:"小说描述"`
	NewChapter string 	`json:"new_chapter" comment:"搜索结果最新章节"`
	URL string 	`json:"url" comment:"搜索结果链接"`
	Source string	`json:"source" comment:"搜索结果来源"`
}