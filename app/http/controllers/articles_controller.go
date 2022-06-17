package controllers

import (
	"fmt"
	"goblog/app/http/requests"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)
type ArticlesController struct {

}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 2. 读取对应的文章数据
	a, err := article.Get(id)
	fmt.Printf("%T\n", a.User)
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
		view.Render(w, view.D{"Article" : a}, "articles.show", "articles._article_meta")
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request)  {
	articles, err := article.GetAll()
	if err !=nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}else {
		view.Render(w, view.D{"Articles": articles}, "articles.index","articles._article_meta")
	}
	//tem, err := template.ParseFiles("resources/views/articles/index.gohtml")
	//err = tem.Execute(w, articles)
	//if err != nil {
	//	logger.LogError(err)
	//}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request)  {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	_article := article.Article{
		Title: title,
		Body:  body,
	}
	errors := requests.ValidateArticleForm(_article)
	// 检查是否有错误
	if len(errors) == 0 {

		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Title":  title,
			"Body":   body,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
	}
}

// Edit 文章更新页面
func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

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
		// 4. 读取成功，显示编辑文章表单
		view.Render(w, view.D{
			"Title": _article.Title,
			"Body": _article.Body,
			"Article": _article,
			"Errors":  make(map[string]string),
		}, "articles.edit", "articles._form_field")
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
	art.Title = r.PostFormValue("title")
	art.Body = r.PostFormValue("body")
	errors := requests.ValidateArticleForm(art)
	if len(errors) > 0 {
		// 4.3 表单验证不通过，显示理由
		view.Render(w, view.D{
			"Title":   art.Title,
			"Body":    art.Body,
			"Article": art,
			"Errors":  errors,
		}, "articles.edit", "articles._form_field")
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
