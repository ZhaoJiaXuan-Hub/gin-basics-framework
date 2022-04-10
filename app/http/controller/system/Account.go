package system

import (
	"admin/app/entity"
	"admin/app/service/system/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(Context *gin.Context) {
	//定义一个User变量
	var user entity.SystemAccount
	user.UserName = Context.Query("username")
	user.NickName = Context.Query("nickname")
	user.Password = Context.Query("password")
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	Context.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := account.CreateSystemAccount(&user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		Context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}
