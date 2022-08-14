package model

import (
	"errors"
	// "fmt"
	"time"
)

// hint: 如果你想直接返回结构体，可以考虑在这里加上`json`的tag
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

// 这里提供返回的submission的示例结构
type ReturnSub struct {
	UserName  string `json:"user"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"time"`
	Score     int    `json:"score"`
	UserVotes int    `json:"votes"`
	Subs      [3]int `json:"subs"`
}

type MyReturn struct {
	UserName  string
	Avatar    string
	CreatedAt int64
	Score     int
	Subscore1 int
	Subscore2 int
	Subscore3 int
	Votes     int
}

/*TODO: 添加相应的与数据库交互逻辑，补全参数和返回值，可以参考user.go的设计思路*/

func CreateSubmission(name string, avatar string, content string) (error, uint) {
	var sub Submission

	err, _ := GetUserByName(name)

	if err != nil {
		CreateUser(name)
	}

	total, sub1, sub2, sub3 := Score_calculator(&content)

	if total < 0 {
		return errors.New("提交内容格式非法"), 0
	}

	sub.UserName = name
	sub.Avatar = avatar
	sub.CreatedAt = time.Now().Unix()
	sub.Score = total
	sub.Subscore1 = sub1
	sub.Subscore2 = sub2
	sub.Subscore3 = sub3

	tx := DB.Create(&sub)

	if tx.Error == nil {
		return nil, sub.ID
	}
	return errors.New("提交内容格式非法"), 0

}

func GetUserSubmissions(name string) (error, []ReturnSub) {
	//返回某一用户的所有提交
	//在查询时可以使用.Order()来控制结果的顺序，详见https://gorm.io/zh_CN/docs/query.html#Order
	//当然，也可以查询后在这个函数里手动完成排序
	var ALlsub []ReturnSub
	err, _ := GetUserByName(name)
	if err != nil {
		return err, ALlsub
	}
	tx := DB.Model(&Submission{}).Where("user_name=?", name).Order("created_at desc").Find(&ALlsub)
	if tx.Error == nil {
		return nil, ALlsub
	}
	return tx.Error, ALlsub
}

func GetLeaderBoard() (error, []ReturnSub) {
	//一个可行的思路，先全部选出submission，然后手动选出每个用户的最后一次提交

	sql := "select r.user_name, r.avatar,r.created_at, r.score, r.subscore1, r.subscore2, r.subscore3 , users.votes from (select submissions.user_name, submissions.avatar, submissions.created_at ,submissions.score, submissions.subscore1, submissions.subscore2, submissions.subscore3 from submissions ,( select user_name, MAX(created_at) as mt from submissions group by user_name) sq  where submissions.user_name = sq.user_name and created_at = sq.mt )r, users where r.user_name = users.user_name order by r.score desc, r.created_at asc;"
	var raw_Board []MyReturn
	var Board []ReturnSub
	err := DB.Raw(sql).Scan(&raw_Board).Error
	if err != nil {
		return err, Board
	}

	l := len(raw_Board)
	// fmt.Println(l)
	for i := 0; i < l; i++ {
		var subs = [3]int{}
		subs[0] = raw_Board[i].Subscore1
		subs[1] = raw_Board[i].Subscore2
		subs[2] = raw_Board[i].Subscore3

		// subs := "[" + strconv.Itoa(q1) + " " + strconv.Itoa(q2) + " " + strconv.Itoa(q3) + "]"

		a := ReturnSub{
			UserName:  raw_Board[i].UserName,
			Avatar:    raw_Board[i].Avatar,
			CreatedAt: raw_Board[i].CreatedAt,
			Score:     raw_Board[i].Score,
			UserVotes: raw_Board[i].Votes,
			Subs:      subs,
		}
		Board = append(Board, a)
	}

	return nil, Board

	//DB.Model(&Submission{}).Where("1=1").Find(&AllSub)
	//在这里添加逻辑！
	//DB.Table("user").Joins("right join submisson submission on user.user_name = submission.user_name")

}

// //select * from submissions ,( select user_name, MAX(created_at) as mt from submissions group by user_name) sq  where submissions.user_name = sq.user_name and created_at = sq.mt order by submissions.score desc, submissions.created_at asc;

// select * from (select submissions.user_name, submissions.avatar, submissions.created_at ,submissions.score, submissions.subscore1, submissions.subscore2, submissions.subscore3 from submissions ,( select user_name, MAX(created_at) as mt from submissions group by user_name) sq  where submissions.user_name = sq.user_name and created_at = sq.mt order by submissions.score desc, submissions.created_at asc)r, users where r.user_name = users.user_name;
