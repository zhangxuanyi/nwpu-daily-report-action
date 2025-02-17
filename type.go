package main

type (
	UserInfo struct {
		Name      string
		StudentID string
		Password  string
	}

	LoginForm struct {
		Username    string `naive:"username"`
		Password    string `naive:"password"`
		CurrentMenu string `naive:"currentMenu" default:"1"`
		Execution   string `naive:"execution" default:"e1s1"`
		EventId     string `naive:"_eventId" default:"submit"`
		Geolocation string `naive:"geolocation"`
	}

	ReportForm struct {
		Hsjc        string `naive:"hsjc" default:"1"`
		Xasymt      string `naive:"xasymt" default:"1"`
		ActionType  string `naive:"actionType" default:"addRbxx"`
		UserLoginId string `naive:"userLoginId"`
		Szcsbm      string `naive:"szcsbm" default:"1"`
		Bdzt        string `naive:"bdzt" default:"1"`
		Szcsmc      string `naive:"szcsmc" default:"在学校"`
		Sfyzz       string `naive:"sfyzz" default:"0"`
		Sfqz        string `naive:"sfqz" default:"0"`
		Tbly        string `naive:"tbly" default:"sso"`
		Qtqksm      string `naive:"qtqksm"`
		Ycqksm      string `naive:"ycqksm"`
		UserType    string `naive:"userType" default:"2"`
		Username    string `naive:"userName"`
	}

	ReportResponse struct {
		Status string `json:"state"`
	}
)
