package model

type LogInfo struct {
	Business Business
	Log      Log
	TTLog    TTLog
	UserInfo []UserInfo
}

type Business struct {
	LastMonthAlarm    int      // 上个月告警量
	LastCapture       int      // 上个月抓拍量
	PortraitCount     int      // 本月人像库量
	PortraitIncrement int      // 人像库上月和上上个月的增量
	AllDay            []string // 上个月的所有日期
	Yue               string
	TTAnJianCount     int // 上个月图腾案件量
}

type Log struct {
	List []struct {
		CreateTime   int64  `json:"createTime"`
		Description  string `json:"description,omitempty"`
		OperaName    string `json:"operaName"`
		OrganName    string `json:"organName"`
		Realname     string `json:"realname"`
		ResourceName string `json:"resourceName"`
		Username     string `json:"username"`
	} `json:"list"`
	Total     int `json:"total"` // 日志总数
	TotalPage int `json:"totalPage"`
}

type TTLog struct {
	List []struct {
		Audit        int    `json:"audit"`
		CreateTime   string `json:"createTime"`
		Description  string `json:"description"`
		Ip           string `json:"ip"`
		LogId        int    `json:"logId"`
		OperaName    string `json:"operaName"`
		OrganId      int    `json:"organId"`
		OrganName    string `json:"organName"`
		Realname     string `json:"realname"`
		Remark       string `json:"remark,omitempty"`
		ResourceId   int    `json:"resourceId"`
		ResourceName string `json:"resourceName"`
		Type         int    `json:"type"`
		UserId       int    `json:"userId"`
		Username     string `json:"username"`
	} `json:"list"`
	Page      int `json:"page"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

type UserInfo struct {
	Realname string
	Username string
	OrganId  int64
	Name     string
	State    int
}
