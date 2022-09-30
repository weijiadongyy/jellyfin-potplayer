package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket实现
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据
		_, message, err := ws.ReadMessage()
		println(string(message))
		if err != nil {
			break
		}
		attr := &os.ProcAttr{
			// files指定新进程继承的活动文件对象
			// 前三个分别为，标准输入、标准输出、标准错误输出
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			// 新进程的环境变量
			Env: os.Environ(),
		}

		potPath := "D:\\Program Files (x86)\\Pure Codec\\x64\\PotPlayerMini64.exe"
		_, err = os.StartProcess(potPath, []string{potPath, string(message)}, attr)

		//写入ws数据
		//err = ws.WriteMessage(mt, message)
		//if err != nil {
		//	break
		//}
	}
}

func main() {
	r := gin.Default()
	r.GET("/play", ping)
	r.Run(":61142")
}
