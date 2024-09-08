package model

import "go-blog/config"

// 主页数据
type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int //用户当前的页码
	Pages     []int
	PageEnd   bool
}
