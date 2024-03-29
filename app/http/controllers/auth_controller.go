package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		// 3. 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		_user.Create()

		if _user.ID > 0 {
			auth.Login(_user)
			flash.Success("恭喜您注册成功！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
		    w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}

	// 5. 表单不通过 —— 重新显示表单
}

// Login 登录页面
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 登录
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	//var data = user.User{
	//	Email: email,
	//	Password: password,
	//}
	//errs := requests.ValidateLoginForm(data)
	//if len(errs) > 0 {
	//	view.RenderSimple(w, view.D{"Error":errs["password"]}, "auth.login")
	//}
	err := auth.Attempt(email, password)
	if err != nil {
		logger.LogError(err)
		view.RenderSimple(w, view.D{"Email" : email, "Error":err.Error()}, "auth.login")
	}
	flash.Success("欢迎回来！")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (*AuthController) DoLogout(w http.ResponseWriter,r *http.Request)  {
	auth.Logout()
	flash.Success("您已退出登录")
	http.Redirect(w, r, "/", http.StatusFound)
}