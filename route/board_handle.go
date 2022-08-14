package route

import (
	"fmt"
	"leadboard/model"

	"github.com/gin-gonic/gin"
)

// TODO:在这里完成handle function，返回所有的leader board内容
func HandleGetBoard(g *gin.Context) {
	_, board := model.GetLeaderBoard()

	g.JSON(200, board)
}

// TODO:在这里完成返回一个用户提交历史的Handle function
func HandleUserHistory(g *gin.Context) {
	name := g.Param("name")
	fmt.Println(name)
	err, subs := model.GetUserSubmissions(name)
	if err == nil {
		g.JSON(200,
			subs,
		)
	} else {
		g.JSON(400, gin.H{
			"code": -1,
		})
	}

}

// TODO:在这里完成接受提交内容，进行评判的handle function
func HandleSubmit(g *gin.Context) {
	type SubmitContent struct {
		User    string `json:"user"`
		Avatar  string `json:"avatar"`
		Content string `json:"content"`
	}

	var js SubmitContent
	err := g.BindJSON(&js)

	if err != nil {
		g.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数不全",
		})
		return
	}

	name := js.User
	avatar := js.Avatar
	content := js.Content

	if len(name) > 255 {
		g.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户名长度超过255",
		})
		return
	}

	if len(avatar) > 102400 {
		g.JSON(200, gin.H{
			"code": -2,
			"msg":  "Avatar大小超过75KiB",
		})
		return
	}

	err, _ = model.CreateSubmission(name, avatar, content)

	if err != nil {
		g.JSON(200, gin.H{
			"code": -3,
			"msg":  err.Error(),
		})
		return
	}
	_, board := model.GetLeaderBoard()

	g.JSON(200, gin.H{
		"code": 0,
		"msg":  "提交成功",
		"data": gin.H{
			"leaderboard": board,
		},
	})

}
