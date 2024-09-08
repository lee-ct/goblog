package dao

import (
	"go-blog/model"
	"log"
)

func GetAllCategory() ([]model.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Fatalln("GetAllCategory出错", err)
		return nil, err
	}
	var categories []model.Category

	for rows.Next() {
		var category model.Category
		rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)

		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryNameById(categoryId int) string {
	row := DB.QueryRow("select name from blog_category where cid=?", categoryId)
	var name string
	row.Scan(&name)
	return name
}
