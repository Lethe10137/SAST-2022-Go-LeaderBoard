package model

//hint: 如果你想直接返回结构体，可以考虑在这里加上`json`的tag
type Submission struct {
	ID        uint   `gorm:"not null;autoIncrement"`
	UserName  string `gorm:"type:varchar(255);"`
	Avatar    string //头像base64，也可以是一个头像链接
	CreatedAt int64  //提交时间
	Score     int    //评测成绩
	Subscore1 int    //评测小分
	Subscore2 int    //评测小分
	Subscore3 int    //评测小分
}

//这里提供返回的submission的示例结构
type ReturnSub struct {
	UserName  string `json:"user"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"time"`
	Score     int    `json:"score"`
	UserVotes int    `json:"votes"`
	Subscore1 int
	Subscore2 int
	Subscore3 int
}

/*TODO: 添加相应的与数据库交互逻辑，补全参数和返回值，可以参考user.go的设计思路*/

func CreateSubmission() {

}

func GetUserSubmissions() {
	//返回某一用户的所有提交
	//在查询时可以使用.Order()来控制结果的顺序，详见https://gorm.io/zh_CN/docs/query.html#Order
	//当然，也可以查询后在这个函数里手动完成排序

}

func GetLeaderBoard() []ReturnSub {
	//一个可行的思路，先全部选出submission，然后手动选出每个用户的最后一次提交
	var AllSub []ReturnSub
	DB.Model(&Submission{}).Where("1=1").Find(&AllSub)
	//在这里添加逻辑！
	return AllSub
}
