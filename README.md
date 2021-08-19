toNovel
------------
tonovel 是一个简洁，干净的小说聚合系统

## 概述
预览：地址已失效

后端地址：[https://github.com/hipig/tonovel-go](https://github.com/hipig/tonovel-go)

前端地址：[https://github.com/hipig/tonovel-vue](https://github.com/hipig/tonovel-vue?_blank)

这是学习 go 以来做的第一个小项目，整体为前后端分离，服务端使用了 colly 爬虫和 iris Web框架，目前只内置了 3 个书源，书源格式为 xpath （ datasource 目录）。
因为涉及小说版权原因，没有做入库操作，所以在源站采集的时候，速度较慢。，目前功能较单一，只支持聚合搜索，查看详情，章节列表，章节内容等功能，后续会考虑做用户中心，历史记录，书架等交互功能。
后期功能会慢慢完善。
go 初学者， 大佬们请轻喷

## 截图
![image](https://user-images.githubusercontent.com/24596908/77449305-bf943900-6e2c-11ea-8513-9237f615a974.png)

![image](https://user-images.githubusercontent.com/24596908/77449504-f79b7c00-6e2c-11ea-84f0-619d6cdb439b.png)
## 运行
```shell script
git clone https://github.com/hipig/tonovel-go.git tonovel
cd tonovel
go run main.go
```
默认端口为 8080

## TODO

### 模块
- [x] 聚合搜索
- [x] 书籍详情
- [x] 内容阅读
- [ ] 用户中心
- [ ] 我的书架
- [ ] 历史记录

### 优化
- ~~手机端自适应~~
- 首屏加载过慢
- 搜索排序聚合
- 阅读页换源

## 鸣谢
* [kataras/iris](https://github.com/kataras/iris)
* [gocolly/colly](https://github.com/gocolly/colly)

## License
Licensed under [The MIT License (MIT)](LICENSE).
