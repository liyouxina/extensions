package main

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"sync"
)

var peoplePool map[string]*People
var peoplePoolLock sync.Mutex

type Resp struct {
	Msg        string             `json:"msg"`
	QrCode     string             `json:"qrCode"`
	PeoplePool map[string]*People `json:"peoples"`
}

func main() {
	peoplePool = map[string]*People{}
	peoplePoolLock = sync.Mutex{}
	go wsServe()
	server := gin.Default()
	server.GET("/report", report)
	server.GET("/getAllLiveCode", getAllLiveCode)
	server.GET("/getQrCode", getQrCode)
	server.GET("/showLiveQrCode", showLiveQrCode)
	_ = server.Run("0.0.0.0:10183")
}

func getQrCode(context *gin.Context) {
	name := context.Query("name")
	if name == "" {
		context.JSON(200, Resp{
			Msg: "名称不能为空",
		})
		return
	}
	if peoplePool[name] == nil {
		context.JSON(200, Resp{
			Msg: "找不到这个人",
		})
		return
	}
	context.JSON(200, Resp{
		Msg:    "成功",
		QrCode: peoplePool[name].QrCode,
	})
}

func report(context *gin.Context) {
	name := context.Query("name")
	qrCode := context.Query("qrcode")
	if name == "" {
		context.JSON(200, Resp{
			Msg: "名称不能为空",
		})
		return
	}
	peoplePoolLock.Lock()
	defer peoplePoolLock.Unlock()
	peoplePool[name] = &People{
		QrCode: qrCode,
		Name:   name,
	}
	context.JSON(200, Resp{
		Msg: "上报成功",
	})
}

func getAllLiveCode(context *gin.Context) {
	context.JSON(200, Resp{
		Msg:        "成功",
		PeoplePool: peoplePool,
	})
}

func showLiveQrCode(context *gin.Context) {
	name := context.Query("name")
	if name == "" {
		_, err := context.Writer.Write([]byte("名称不能为空"))
		if err != nil {
			log.Info("showLiveQrCode 出错 {}", err.Error())
		}
		context.JSON(200, Resp{
			Msg: "失败",
		})
		return
	}
	people := peoplePool[name]
	if people == nil {
		_, err := context.Writer.Write([]byte("没有上报信息"))
		if err != nil {
			log.Info("showLiveQrCode 出错 {}", err.Error())
		}
		return
	}
	content, err := ioutil.ReadFile("./backend/health/qrcode.html")
	if err != nil {
		log.Error("打开文件失败：%v\n", err)
		context.JSON(200, Resp{
			Msg: "失败",
		})
		return
	}
	_, _ = context.Writer.Write([]byte(strings.Replace(strings.Replace(string(content), "%s", people.Name, 1), "%s", people.QrCode, 1)))
	return
}
