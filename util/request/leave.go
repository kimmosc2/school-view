package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	// WalkAuthURL is school-walker verification URL
	WalkAuthURL = "http://127.0.0.1:8000/api/v1/leave/verification"
	WalkSaveURL = "http://127.0.0.1:8000/api/v1/leave/save"
)

// Response is common json response struct
type Response struct {
	Data ResponseUser `json:"data"`
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
}

// ResponseUser user infomation form school-walk
type ResponseUser struct {
	ID           uint   `json:"id"`
	SchoolNumber string `json:"school_idcode"`
	// name
	Name       string `json:"name"`
	Teacher    string `json:"teacher"`
	Contact    string `json:"contact"`
	ContactTel string `json:"contact_tel"`
	Direction  string `json:"direction"`
	Reason     string `json:"reason"`
	// hour
	StartTime uint8 `json:"start_time"`
	// hour
	EndTime uint8 `json:"end_time"`
	// 0 as disabled, 1 as enabled
	State int8 `json:"state"`
}

func (user ResponseUser) BuildForm() url.Values {
	u := url.Values{}
	u.Set("school_idcode", user.SchoolNumber)
	u.Set("teacher", user.Teacher)
	u.Set("contact", user.Contact)
	u.Set("contact_tel", user.ContactTel)
	u.Set("direction", user.Direction)
	u.Set("reason", user.Reason)
	u.Set("start_time", strconv.Itoa(int(user.StartTime)))
	u.Set("end_time", strconv.Itoa(int(user.EndTime)))
	return u
}

// GetUserInfo return user information
func GetUserInfo(cookie string) (Response, error) {
	// c := url.Values{}
	// c.Set("cookie", cookie)
	req, err := http.NewRequest(http.MethodGet, WalkAuthURL+"?cookie="+cookie, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, "build request error")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, errors.Wrap(err, "http request error")
	}
	defer resp.Body.Close()
	user := new(Response)
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return Response{}, errors.Wrap(err, "json decode error")
	}
	return *user, nil
}

func SaveUserInfo(user ResponseUser) error {
	user.StartTime = 8
	user.EndTime = 20
	request, err := http.NewRequest(http.MethodPost, WalkSaveURL, strings.NewReader(user.BuildForm().Encode()))
	if err != nil {
		return fmt.Errorf("build request error:%s", err)
	}
	request.Form = user.BuildForm()
	request.Header.Set("Content-Type", `application/x-www-form-urlencoded`)
	// fmt.Println(request.Form)
	client := http.Client{}
	do, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("http request error:%s", err)
	}
	do.Body.Close()
	return nil
}
