package route

import (
	"leadboard/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//TODO:完成下方的两个Handle Function，其中第一个作为中间件使用，第二个处理投票逻辑

func CheckUserAgent(g *gin.Context) {
	//用于检查User Agent的中间件
	userAgent := g.Request.UserAgent()
	//DONE_TODO:在这里完成判断User Agent的逻辑，最简单的方法是判断User Agent是否为空字符串
	var theIndex = strings.Index(userAgent, "Mozilla") // -1 if there's no "Mozilla" in userAgent
	if theIndex < 0 {
		g.JSON(http.StatusForbidden, gin.H{
			"msg": "No Robots!",
		})
		g.Abort()
	} else {
		g.Next()
	}
}

func HandleVote(g *gin.Context) {
	type VoteForm struct {
		UserName string `json:"user"`
	}
	var form VoteForm
	if err := g.ShouldBindJSON(&form); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "Invalid Form",
		})
		return
	} else {
		//TODO:完成投票数+1这一操作，注意不要出现并发上的问题，加油 qwq
		//推荐自己完成，也可以使用model/user.go中给出的方法
		err = model.AddVoteForUser(form.UserName)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  "User NOT found",
			})
		} else {
			_, board := model.GetLeaderBoard()
			g.JSON(200, gin.H{
				"code": 0,
				"msg":  "投票成功",
				"data": gin.H{
					"leaderboard": board,
				},
			})
		}
	}
}
