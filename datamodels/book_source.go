package datamodels

type BookSource struct {
	SourceName string	`json:"source_name" comment:"书源名称"`
	SourceURL string 	`json:"source_url" comment:"书源网址"`
	SourceKey string 	`json:"source_key" comment:"书源标识"`

	SearchURL string 	`json:"search_url" comment:"搜索网址"`
	SearchItemRule string 	`json:"search_list_rule" comment:"搜索结果规则"`
	SearchItemNameRule string 	`json:"search_item_name_rule" comment:"搜索结果名称规则"`
	SearchItemAuthorRule string 	`json:"search_item_author_rule" comment:"搜索结果作者规则"`
	SearchItemCoverRule string 	`json:"search_item_cover_rule" comment:"搜索结果封面规则"`
	SearchItemCategoryRule string 	`json:"search_item_category_rule" comment:"搜索结果分类规则"`
	SearchItemNewChapterRule string 	`json:"search_item_new_chapter_rule" comment:"搜索结果最新章节规则"`
	SearchItemURLRule string 	`json:"search_item_url_rule" comment:"搜索结果链接规则"`


	DetailBookItemRule string 	`json:"detail_book_item_rule" comment:"小说详情规则"`
	DetailBookNameRule string 	`json:"detail_book_name_rule" comment:"小说名称规则"`
	DetailBookAuthorRule string 	`json:"detail_book_author_rule" comment:"小说作者规则"`
	DetailBookCoverRule string 	`json:"detail_book_cover_rule" comment:"小说封面规则"`
	DetailBookCategoryRule string 	`json:"detail_book_category_rule" comment:"小说分类规则"`
	DetailBookDescriptionRule string 	`json:"detail_book_description_rule" comment:"小说描述规则"`
	DetailBookNewChapterRule string 	`json:"detail_book_new_chapter_rule" comment:"小说最新章节规则"`

	DetailChapterListURLRule string 	`json:"detail_chapter_url_rule" comment:"小说章节列表链接规则"`
	DetailNewChapterRule string 	`json:"detail_new_chapter_rule" comment:"小说新章节规则"`
	DetailNewChapterTitleRule string 	`json:"detail_new_chapter_rule" comment:"小说新章节名称规则"`
	DetailNewChapterURLRule string 	`json:"detail_new_chapter_rule" comment:"小说新章节名称规则"`
	DetailChapterRule string	`json:"detail_chapter_list_rule" comment:"小说章节列表规则"`
	DetailChapterTitleRule string	`json:"detail_chapter_list_rule" comment:"小说章节名称规则"`
	DetailChapterURLRule string	`json:"detail_chapter_list_rule" comment:"小说章节链接规则"`

	ContentTitleRule string	`json:"chapter_content_rule" comment:"内容标题规则"`
	ContentTextRule string	`json:"chapter_content_rule" comment:"内容正文规则"`
	ContentPreviousURLRule string	`json:"chapter_previous_url_rule" comment:"内容上一章链接规则"`
	ContentNextURLRule string	`json:"chapter_next_url_rule" comment:"内容下一章链接规则"`

	Weight int 	`json:"weight" comment:"权重"`
}