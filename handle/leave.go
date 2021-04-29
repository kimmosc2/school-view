package handle

import (
	"html/template"
	"log"
	"net/http"
	"school-walker/view/util/request"
	"school-walker/view/util/tmpl"
)

// ViewLeave control the leave page, it request
// school-walker API Gateway,this handle only support
// method GET
func ViewLeave(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	// get user information
	useCookie := r.Form.Get("cookie")
	info, err := request.GetUserInfo(useCookie)
	if err != nil {
		log.Printf("get user info error:%+v\ncookie:%v", err, useCookie)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// first use
	if info.Code == 4 {
		first, err := template.New("leave_first.html").ParseFiles("./static/leave_first.html")
		if err != nil {
			log.Printf("parse first page failure:%s", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = first.Execute(rw, info.Data)
		if err != nil {
			log.Printf("first page execute error:%s,data:%v", err, info.Data)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	// no authority
	// if info.Data.State == 0 {
	// 	// rw.Write([]byte("您无权访问此页面"))
	// 	rw.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	files, err := template.New("leave.html").Funcs(template.FuncMap{
		"StartCompute":    tmpl.StartCompute,
		"EndCompute":      tmpl.EndCompute,
		"DurationCompute": tmpl.DurationCompute,
		"ApplyTime":       tmpl.ApplyTime,
	}).ParseFiles("./static/leave.html")
	if err != nil {
		log.Printf("build template error:%s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = files.Execute(rw, info.Data)
	if err != nil {
		log.Printf("excute tempalte error:%s,data:%v", err, info.Data)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func DataSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// fmt.Println(r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	code := r.Form.Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := new(request.ResponseUser)
	user.SchoolNumber = code
	user.Reason = r.Form.Get("reason")
	user.Direction = r.Form.Get("direction")
	user.Contact = r.Form.Get("contact")
	user.ContactTel = r.Form.Get("contact_tel")
	user.Teacher = r.Form.Get("teacher")

	err := request.SaveUserInfo(*user)
	if err != nil {
		log.Printf("save user info error:%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("保存成功,请退出页面重新进入,如果还是显示第一次使用或数据为空,请联系管理员解决"))
	return
}
