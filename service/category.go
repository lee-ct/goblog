package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/model"
	"html/template"
	"log"
)

func GetAllIndexInfo(page, pageSize int) (*model.HomeResponse, error) {

	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("dao.GetAllCategory出错", err)
		return nil, nil
	}

	posts, err := dao.GetALLPost(page, pageSize)
	if err != nil {
		log.Println("service中GetALLPost出错了", err)
	}
	var postMores []model.PostMore
	for _, post := range posts {
		userName := dao.GetUserNameBYId(post.UserId)
		categoryName := dao.GetCategoryNameById(post.CategoryId)

		postContent := []rune(post.Content)
		if len(postContent) > 100 {
			postContent = postContent[0:100]
		}
		postmore := model.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(postContent),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     model.DateDay(post.CreateAt),
			UpdateAt:     model.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postmore)

	}

	Total := dao.CountGetAllPost()

	pageCount := (Total-1)/pageSize + 1
	var pages []int

	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}
	hr := &model.HomeResponse{
		Viewer:    config.Viewer{},
		Categorys: categorys,
		Posts:     postMores,
		Total:     Total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageCount,
	}
	return hr, nil
}

func GetPostsByCategoryId(cId int, page int, pageSize int) (*model.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("dao.GetAllCategory出错", err)
		return nil, nil
	}

	posts, err := dao.GetPostPageByCategoryId(cId, page, pageSize)
	if err != nil {
		log.Println("service中GetALLPost出错了", err)
	}
	var postMores []model.PostMore
	for _, post := range posts {
		userName := dao.GetUserNameBYId(post.UserId)
		categoryName := dao.GetCategoryNameById(post.CategoryId)

		postContent := []rune(post.Content)
		if len(postContent) > 100 {
			postContent = postContent[0:100]
		}

		postmore := model.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(postContent),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     model.DateDay(post.CreateAt),
			UpdateAt:     model.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postmore)

	}

	Total := dao.CountGetAllPostByCategoryId(cId)

	pageCount := (Total-1)/pageSize + 1
	var pages []int

	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}
	hr := &model.HomeResponse{
		Viewer:    config.Viewer{},
		Categorys: categorys,
		Posts:     postMores,
		Total:     Total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageCount,
	}
	categoryName := dao.GetCategoryNameById(cId)
	CategoryResponse := &model.CategoryResponse{
		hr,
		categoryName,
	}

	return CategoryResponse, nil
}
