package dao

import (
	"fmt"
	"go-blog/model"
	"log"
)

func GetALLPost(post, postSize int) ([]model.Post, error) {
	post = (post - 1) * postSize
	rows, err := DB.Query("select * from blog_post limit ?,?", post, postSize)
	if err != nil {
		log.Println("GetALLPost出错了", err)
		return nil, nil
	}
	var posts []model.Post
	for rows.Next() {
		var post model.Post
		rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.Type,
			&post.CreateAt,
			&post.UpdateAt,
			&post.ViewCount,
		)

		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPost() (count int) {
	row := DB.QueryRow("select count(*) from blog_post")
	var cnt int
	row.Scan(&cnt)

	return cnt
}

func GetPostPageByCategoryId(cId, page, pageSize int) ([]model.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id=? limit ?,?", cId, page, pageSize)
	if err != nil {
		log.Println("GetALLPost出错了", err)
		return nil, nil
	}
	var posts []model.Post
	for rows.Next() {
		var post model.Post
		rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.Type,
			&post.CreateAt,
			&post.UpdateAt,
			&post.ViewCount,
		)
		posts = append(posts, post)
		fmt.Println(posts)
	}
	return posts, nil
}

func CountGetAllPostByCategoryId(cId int) (count int) {
	row := DB.QueryRow("select count(*) from blog_post where category_id=?", cId)
	var cnt int
	row.Scan(&cnt)
	return cnt
}
