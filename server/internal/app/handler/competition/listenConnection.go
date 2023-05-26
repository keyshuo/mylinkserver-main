package netcompete

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// 设置最高分，便于直接在接收分数后，直接判断游戏是否结束，结束就进入对应结束流程
const topScore = 100

func Test(c *gin.Context) {
	httpRespone.WriteOK(c, gin.H{
		"msg": Rooms,
	})
}

//设计思路：
//双方交换sign，确定唯一房间号
//直接设置房间属性，便于双方收发，
//

// HandlerWebsocket
// 后面可以把player改为数组形式
// 点击”准备“后会进入当前状态,不需要发送信息，携带token，对token进行鉴权解析即可
// 如果对方掉线不连呢？新开线程对掉线状态检测
func HandlerWebsocket(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	//flag代表对方是否准备
	flag := false
	//flag1代表是否双方是否准备就绪
	flag1 := false
	//flag2代表游戏是否已经开始
	flag2 := false
	//flag3代表游戏是否双方都结束了
	flag3 := false
	//temp用于websocket接收用户信号
	var temp Player
	//获取玩家所在房间号以及玩家名称
	roomNumber, _ := strconv.Atoi(c.Query("sign"))
	username := c.Value("username").(string)
	//升级websocket
	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	//设置用户连接状态，根据玩家名找到玩家位置，并设置其连接状态为已连接
	//Rooms[roomNumber].Mutex.Lock()
	//Rooms[roomNumber].Clients[conn] = true
	pos := getPlayerPos(username, Rooms[roomNumber].Competition)
	setPlayerConnectionStatus(pos, roomNumber, true)
	//Rooms[roomNumber].Mutex.Unlock()

	//掉线超20s时间检测
	disconnected := make(chan bool)
	//websocket通信
	for {
		//每隔5s发送一次
		time.Sleep(2 * time.Second)
		//若对方掉线
		if getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition) {
			//新开线程检测是否超过20s，若超过，则返回true
			go func() {
				counter := 0
				time.Sleep(1 * time.Second)
				counter += 1
				//若对手状态未连接上，且计数大于20，判输
				if !getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition) {
					if counter > 20 {
						setPlayerEnd(username, Rooms[roomNumber].Competition, getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition))
						disconnected <- true
					}
				} else {
					setPlayerEnd(username, Rooms[roomNumber].Competition, getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition))
					disconnected <- false
				}
			}()
		}
		//主线程处理超时20s的用户，默认判输
		select {
		case <-disconnected:
			conn.WriteJSON(gin.H{
				"end":      true,
				"realtime": getOtherPlayerTime(username, Rooms[roomNumber].Competition),
				"rival":    getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition),
			})
			//发送结束状态，真实分数，比赛对手的连接状态，应该直接设置competition对象的内部，然后直接发送competition
			return
		default:
			if !flag && getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition) {
				conn.WriteJSON(gin.H{
					"rival": true,
				})
				flag = true
			} else {
				conn.WriteJSON(gin.H{
					"rival": false,
				})
			}

			//双方准备就绪后开始，发送5后开始
			if flag && !flag1 {
				err := conn.WriteJSON(gin.H{
					"start": 5,
					"rival": getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition),
				})
				if err != nil {
					log.Println("Failed to write JSON from WebSocket:", err)
					break
				}
				flag1 = true
			}

			//接收用户开始游戏标志
			if flag && flag1 && !flag2 {
				err = conn.ReadJSON(&temp)
				if err != nil {
					log.Println("Failed to read message from WebSocket:", err)
					break
				}
				//如果用户开始游戏
				if temp.Start == true {
					//设置玩家一开始游戏状态
					if pos == 1 {
						Rooms[roomNumber].Competition.Player1.Start = true
						flag2 = true
					}
					//设置玩家二开始游戏状态
					if pos == 2 {
						Rooms[roomNumber].Competition.Player2.Start = true
						flag2 = true
					}
				}
			}

			//接收到用户已经开始游戏时，一直读取用户发送的分数信息
			//如果是玩家1就设置玩家1的分数，并将玩家2的分数发给他
			if flag && flag1 && !flag2 && pos == 1 && Rooms[roomNumber].Competition.Player1.Start {
				err := conn.ReadJSON(&temp)
				if err != nil {
					log.Println("Failed to read JSON from WebSocket:", err)
					break
				}
				Rooms[roomNumber].Competition.Player1.Score = temp.Score
				//如果用户达到最高分
				if !flag3 && temp.Score == topScore {
					//设置自己的结束状态，结束运行
					setPlayerEnd(username, Rooms[roomNumber].Competition, true)
					setPlayerTime(username, Rooms[roomNumber].Competition, temp.SumTime)
					flag3 = true
				} else if getOtherPlayerEnd(username, Rooms[roomNumber].Competition) {
					//获取对方结束状态，若对方已结束完成，则服务端告知客户端用户，对方已完成
					conn.WriteJSON(gin.H{
						"end":      true,
						"realtime": getOtherPlayerTime(username, Rooms[roomNumber].Competition),
						"rival":    getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition),
					})
					break
				} else {
					//若双方都未完成，则互相发送比分和连接情况给对方
					err = conn.WriteJSON(gin.H{
						"score": getOtherPlayerScore(username, Rooms[roomNumber].Competition),
						"rival": getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition),
					})
					if err != nil {
						log.Println("Failed to write JSON from WebSocket:", err)
						break
					}
				}
			}
			if flag && flag1 && flag2 && flag3 {
				//双方比赛完成后，统一对比时间
				err = conn.WriteJSON(gin.H{
					"situation": compare(Rooms[roomNumber].Competition),
					"rival":     getOtherPlayerConnectionStatus(username, Rooms[roomNumber].Competition),
				})
				if err != nil {
					log.Println("Failed to write JSON from WebSocket:", err)
					break
				}
			}
		}
		setPlayerConnectionStatus(pos, roomNumber, false)
	}
	//删除websocket连接
	//Rooms[roomNumber].Mutex.Lock()
	//delete(Rooms[roomNumber].Clients, conn)
	//Rooms[roomNumber].Mutex.Unlock()
}

func setPlayerConnectionStatus(pos, roomNumber int, status bool) {
	if pos == 1 {
		Rooms[roomNumber].Competition.Player1.ConnStatus = status
	}
	if pos == 2 {
		Rooms[roomNumber].Competition.Player2.ConnStatus = status
	}
}
