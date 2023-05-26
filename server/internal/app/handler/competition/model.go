package netcompete

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Player struct {
	//玩家名称
	Username any `json:"username"`
	//玩家就绪状态
	Prepare bool `json:"prepare"`
	//连接状态
	ConnStatus bool `json:"connStatus"`
	//玩家分数
	Score int `json:"score"`
	//消耗时长
	SumTime float64 `json:"sumTime"`
	//玩家是否开始
	Start bool `json:"start"`
	//玩家是否结束
	End bool `json:"end"`
}

type Competition struct {
	//比赛标识
	Sign int `json:"sign"`
	//难度
	Grade string `json:"grade"`
	//棋盘号
	Checkerboard int `json:"checkerboard"`
	//玩家一
	Player1 Player `json:"player1"`
	//玩家二
	Player2 Player `json:"player2"`
}

type Room struct {
	//房间内的比赛
	Competition Competition
	//客户端的websocket连接
	//Clients map[*websocket.Conn]bool
	//锁
	//Mutex sync.Mutex
}

var (
	Rooms    map[int]*Room
	mutex    sync.Mutex
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// 根据需要进行来源检查
			return true
		},
	}
)
