package netcompete

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"strconv"
)

// GetOrCreateRoom
// 1.用户点击网络对战，发送请求，请求包含，棋盘号，服务端创建或查找房间：(不要包含难度了，不然不好写)
// ①如果遍历后，没有对应房间，则创建一个房间，放入玩家1的信息
// ②如果遍历后，有房间缺人，则加入玩家2
func GetOrCreateRoom(c *gin.Context) {
	username := c.Query("username")

	if username == "" {
		klog.Error("username is nil")
		httpRespone.WriteFailed(c, "username is nil")
		return
	}
	checkerboard, _ := strconv.Atoi(c.Query("checkerboard"))
	mutex.Lock()
	defer mutex.Unlock()
	for _, room := range Rooms {
		if room.Competition.Player2.username == nil {
			room.Competition.Player2 = Player{
				username:   username,
				Prepare:    false,
				connStatus: false,
				Score:      0,
				sumTime:    0,
				start:      false,
				end:        false,
			}
			httpRespone.WriteOK(c, room.Competition.Sign)
			return
		}
	}
	competition := Competition{
		Sign:         len(Rooms) + 1,
		Grade:        "Easy",
		Checkerboard: checkerboard,
		Player1: Player{
			username:   username,
			Prepare:    false,
			connStatus: false,
			Score:      0,
			sumTime:    9999999,
			start:      false,
			end:        false,
		},
		Player2: Player{
			username:   nil,
			Prepare:    false,
			connStatus: false,
			Score:      0,
			sumTime:    9999999,
			start:      false,
			end:        false,
		},
	}
	room := &Room{
		Competition: competition,
		//Clients:     make(map[*websocket.Conn]bool),
	}
	Rooms[competition.Sign] = room
	httpRespone.WriteOK(c, competition.Sign)
}
func getMyConnectionStatus(username string, competition Competition) bool {
	if username == competition.Player1.username {
		return competition.Player1.connStatus
	}
	if username == competition.Player2.username {
		return competition.Player2.connStatus
	}
	return false
}

func getOtherPlayerConnectionStatus(username string, competition Competition) bool {
	if username == competition.Player1.username {
		return competition.Player2.connStatus
	}
	if username == competition.Player2.username {
		return competition.Player1.connStatus
	}
	return false
}

// 0是没找到位置，1是玩家1，2是玩家2
func getPlayerPos(username string, competition Competition) int {
	if username == competition.Player1.username {
		return 1
	}
	if username == competition.Player2.username {
		return 2
	}
	return 0
}

func getOtherPlayerEnd(username string, competition Competition) bool {
	if username == competition.Player1.username {
		return competition.Player2.end
	}
	if username == competition.Player2.username {
		return competition.Player1.end
	}
	return false
}

func setPlayerEnd(username string, competition Competition, status bool) bool {
	if username == competition.Player1.username {
		competition.Player1.end = status
		return true
	}
	if username == competition.Player2.username {
		competition.Player2.end = status
		return true
	}
	return false
}

func getOtherPlayerScore(username string, competition Competition) int {
	if username == competition.Player1.username {
		return competition.Player2.Score
	}
	if username == competition.Player2.username {
		return competition.Player1.Score
	}
	return 0
}

func setPlayerTime(username string, competition Competition, time float64) bool {
	if username == competition.Player1.username {
		competition.Player1.sumTime = time
		return true
	}
	if username == competition.Player2.username {
		competition.Player2.sumTime = time
		return true
	}
	return false
}

func getOtherPlayerTime(username string, competition Competition) float64 {
	if username == competition.Player1.username {
		return competition.Player2.sumTime
	}
	if username == competition.Player2.username {
		return competition.Player1.sumTime
	}
	return 0
}

func compare(competition Competition) string {
	if competition.Player1.sumTime > competition.Player2.sumTime {
		return "fail"
	} else if competition.Player1.sumTime > competition.Player2.sumTime {
		return "win"
	} else {
		return "equality"
	}
}

func FinishedCompetition(c *gin.Context) {
	roomNumber, _ := strconv.Atoi(c.Query("room"))
	mutex.Lock()
	delete(Rooms, roomNumber)
	mutex.Unlock()
}

//func getOtherPlayerStartStatus(username string, competition Competition) bool {
//	if username == competition.Player1.username {
//		return competition.Player2.start
//	}
//	if username == competition.Player2.username {
//		return competition.Player1.start
//	}
//	return false
//}
