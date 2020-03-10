package datasource

import "tonovel/datamodels"

var BookSources = map[int64]datamodels.BookSource{
	1: {
		SourceName: "酷小说",
		SourceURL: "https://www.kuxiaoshuo.com",
		SourceKey: "kuxiaoshuo",
		SearchURL: "https://www.kuxiaoshuo.com/modules/article/search.php?searchkey=%s",
		SearchListRule: ".grid",
		SearchItemNameRule: "td:eq(0)&text",
		SearchItemAuthorRule: "td:eq(2)&text",
		SearchItemCoverRule: "",
		SearchItemCategoryRule: "",
		SearchItemNewChapterRule: "td:eq(1)&text",
		SearchItemURLRule: "td:eq(0) a&href",
	},
}