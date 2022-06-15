package controllers

import (
	"fmt"
	"goblog/app/http/request"
	"goblog/app/models/article"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
)
type ArticlesController struct {

}

type articleFormData struct {
	article.Article
	URL string
	Errors map[string]string
}

var db = database.DB

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 2. 读取对应的文章数据
	a, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示文章
		view.Render(w,"articles.show", a)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request)  {
	articles, err := article.GetAll()
	if err !=nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}else {
		view.Render(w, "articles.index", articles)
	}
	//tem, err := template.ParseFiles("resources/views/articles/index.gohtml")
	//err = tem.Execute(w, articles)
	//if err != nil {
	//	logger.LogError(err)
	//}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request)  {
	tmp, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err  != nil {
		logger.LogError(err)
		return
	}
	var formData articleFormData
	storeUrl := route.Name2URL("articles.store")
	formData = articleFormData{
		URL: storeUrl,
		Errors: nil,
	}
	err = tmp.Execute(w, formData)
	if err!=nil {
		logger.LogError(err)
	}
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request)  {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	req := request.ArticleRequest{}
	errors := req.Validate(title, body)

	// 检查是否有错误
	if len(errors) == 0 {
		_article := article.Article{
			Title: title,
			Body:  body,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {

		storeURL := route.Name2URL("articles.store")

		data := articleFormData{
			Article: article.Article{Title: title, Body: body},
			URL:    storeURL,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")

		logger.LogError(err)

		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}
}

func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request)  {
	articleId := route.GetRouteVariable("id", r)
	art, err :=article.Get(articleId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章不存在")
		}else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器错误")
		}

	} else {
		updateUrl := route.Name2URL("articles.update", "id", articleId)
		fmt.Println(updateUrl)

		data := articleFormData{
			Article: art,
			URL: updateUrl,
			Errors: nil,
		}
		tem, err :=template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)
		_ = tem.Execute(w, data)
	}
}

func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request)  {
	articleId := route.GetRouteVariable("id", r)
	art, err :=article.Get(articleId)
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "文章不存在")
	}
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	vail := request.ArticleRequest{}
	errors := vail.Validate(title, body)
	updateUrl := route.Name2URL("articles.update", "id", articleId)
	art.Title = title
	art.Body = body
	data := articleFormData{
		Article: art,
		URL: updateUrl,
		Errors: errors,
	}
	if len(errors) > 0 {
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}else {
		rowsAffected, err := art.Update()

		if err != nil {
			// 数据库错误
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
			return
		}

		// √ 更新成功，跳转到文章详情页
		if rowsAffected > 0 {
			showURL := route.Name2URL("articles.show", "id", articleId)
			http.Redirect(w, r, showURL, http.StatusFound)
		} else {
			fmt.Fprint(w, "您没有做任何更改！")
		}
	}
}

func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request)  {
	articleId := route.GetRouteVariable("id", r)
	_article, err := article.Get(articleId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "文章不存在")
		return
	}
	rowsAffected, err := _article.Delete()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error")
		return
	}else {
		if rowsAffected > 0 {
			//fmt.Fprint(w, "删除成功")
			indexUrl := route.Name2URL("articles.index")
			http.Redirect(w,r, indexUrl, http.StatusFound)
		}else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "删除失败")
		}
	}
}
