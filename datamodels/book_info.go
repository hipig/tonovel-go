package datamodels

type BookInfo struct {
	Name string 	`json:"name" comment:"小说名称`
	Author string 	`json:"author" comment:"小说作者"`
	Cover string 	`json:"cover" comment:"小说封面"`
	Category string 	`json:"category" comment:"小说分类"`
	Description string	`json:"category" comment:"小说描述"`
}