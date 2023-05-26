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
	//username := c.Query("username")
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	username := c.Value("username").(string)
	if username == "" {
		klog.Error("username is nil")
		httpRespone.WriteFailed(c, "username is nil")
		return
	}
	checkerboard, _ := strconv.Atoi(c.Query("checkerboard"))
	mutex.Lock()
	defer mutex.Unlock()
	for _, room := range Rooms {
		if room.Competition.Player1.Username == username {
			httpRespone.WriteOK(c, room.Competition)
			return
		}
		if room.Competition.Player1.Username != username && room.Competition.Player2.Username == "" {
			room.Competition.Player2 = Player{
				Username:   username,
				Prepare:    false,
				ConnStatus: false,
				Score:      0,
				SumTime:    0,
				Start:      false,
				End:        false,
			}
			httpRespone.WriteOK(c, room.Competition)
			return
		}
	}

	competition := Competition{
		Sign:         getRoomNumber(),
		Grade:        "Easy",
		Checkerboard: checkerboard,
		Player1: Player{
			Username:   username,
			Prepare:    false,
			ConnStatus: false,
			Score:      0,
			SumTime:    9999999,
			Start:      false,
			End:        false,
		},
		Player2: Player{
			Username:   "",
			Prepare:    false,
			ConnStatus: false,
			Score:      0,
			SumTime:    9999999,
			Start:      false,
			End:        false,
		},
	}
	room := &Room{
		Competition: competition,
		//Clients:     make(map[*websocket.Conn]bool),
	}
	Rooms[competition.Sign] = room
	httpRespone.WriteOK(c, Rooms[competition.Sign].Competition)
}

func getRoomNumber() int {
	i := 1
	for {
		if _, ok := Rooms[i]; !ok {
			return i
		}
		i++
	}
}

func OpponentFound(c *gin.Context) {
	sign, _ := strconv.Atoi(c.Query("sign"))
	httpRespone.WriteOK(c, Rooms[sign].Competition)
}

func getMyConnectionStatus(username string, competition Competition) bool {
	if username == competition.Player1.Username {
		return competition.Player1.ConnStatus
	}
	if username == competition.Player2.Username {
		return competition.Player2.ConnStatus
	}
	return false
}

func getOtherPlayerConnectionStatus(username string, competition Competition) bool {
	if username == competition.Player1.Username {
		return competition.Player2.ConnStatus
	}
	if username == competition.Player2.Username {
		return competition.Player1.ConnStatus
	}
	return false
}

// 0是没找到位置，1是玩家1，2是玩家2
func getPlayerPos(username string, competition Competition) int {
	if username == competition.Player1.Username {
		return 1
	}
	if username == competition.Player2.Username {
		return 2
	}
	return 0
}

func getOtherPlayerEnd(username string, competition Competition) bool {
	if username == competition.Player1.Username {
		return competition.Player2.End
	}
	if username == competition.Player2.Username {
		return competition.Player1.End
	}
	return false
}

func setPlayerEnd(username string, competition Competition, status bool) bool {
	if username == competition.Player1.Username {
		competition.Player1.End = status
		return true
	}
	if username == competition.Player2.Username {
		competition.Player2.End = status
		return true
	}
	return false
}

func getOtherPlayerScore(username string, competition Competition) int {
	if username == competition.Player1.Username {
		return competition.Player2.Score
	}
	if username == competition.Player2.Username {
		return competition.Player1.Score
	}
	return 0
}

func setPlayerTime(username string, competition Competition, time float64) bool {
	if username == competition.Player1.Username {
		competition.Player1.SumTime = time
		return true
	}
	if username == competition.Player2.Username {
		competition.Player2.SumTime = time
		return true
	}
	return false
}

func getOtherPlayerTime(username string, competition Competition) float64 {
	if username == competition.Player1.Username {
		return competition.Player2.SumTime
	}
	if username == competition.Player2.Username {
		return competition.Player1.SumTime
	}
	return 0
}

func compare(competition Competition) string {
	if competition.Player1.SumTime > competition.Player2.SumTime {
		return "fail"
	} else if competition.Player1.SumTime > competition.Player2.SumTime {
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
	httpRespone.WriteOK(c, nil)
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
