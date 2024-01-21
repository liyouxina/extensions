package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/report", report)
	server.GET("/getAllLiveCode", getAllLiveCode)
	server.GET("/showLiveQrCode", showLiveQrCode)
	_ = server.Run("0.0.0.0:10182")
}

func report(context *gin.Context) {
	name := context.Query("name")
	status := context.Query("status")
	qrCode := context.Query("qrCode")
	agent := systemIdAgentPool[systemId]
	if agent == nil {
		context.JSON(200, Resp{
			Msg: "没找到这个设备",
		})
		return
	}
	resp, err := agent.sendMsg(hexContent)
	if err != nil {
		context.JSON(200, Resp{
			Msg: "接收返回数据出错 " + err.Error(),
		})
		return
	}
	context.JSON(200, Resp{
		Msg: *resp,
	})
}

func getAllLiveCode(context *gin.Context) {
	systemId := context.Query("systemId")
	hexContent := context.Query("hex")
	agent := systemIdAgentPool[systemId]
	if agent == nil {
		context.JSON(200, Resp{
			Msg: "没找到这个设备",
		})
		return
	}
	resp, err := agent.sendMsg(hexContent)
	if err != nil {
		context.JSON(200, Resp{
			Msg: "接收返回数据出错 " + err.Error(),
		})
		return
	}
	context.JSON(200, Resp{
		Msg: *resp,
	})
}

func showLiveQrCode(context *gin.Context) {
	systemId := context.Query("systemId")
	hexContent := context.Query("hex")
	agent := systemIdAgentPool[systemId]
	if agent == nil {
		context.JSON(200, Resp{
			Msg: "没找到这个设备",
		})
		return
	}
	resp, err := agent.sendMsg(hexContent)
	if err != nil {
		context.JSON(200, Resp{
			Msg: "接收返回数据出错 " + err.Error(),
		})
		return
	}
	context.JSON(200, Resp{
		Msg: *resp,
	})
}
