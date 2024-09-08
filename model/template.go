package model

import (
	"html/template"
	"io"
	"log"
	"time"
)

// 封装模版，方便扩展自己的业务
type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte("error"))
	}
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {

	tp, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"}, templateDir)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]

	return htmlTemplate, err
}

func Date(form string) string {
	return time.Now().Format(form)
}
func IsODD(index int) bool {
	return index%2 == 0
}

func GetNextName(Nav []string, index int) string {
	return Nav[index+1]
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)

		//每个页面都需要加载layout
		home := templateDir + "home.html"
		footer := templateDir + "layout/footer.html"
		header := templateDir + "layout/header.html"
		pagination := templateDir + "layout/pagination.html"
		personal := templateDir + "layout/personal.html"
		postlist := templateDir + "layout/post-list.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		var err error
		t, err = t.ParseFiles(templateDir+viewName, home, footer, header, pagination, personal, postlist)
		if err != nil {
			log.Println("模板解析错误", err)
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}
