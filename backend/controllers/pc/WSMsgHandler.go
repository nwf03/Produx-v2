package pc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
	"tutorial/db"

	"github.com/go-redis/redis/v8"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

type MessageUser struct {
	UserId   uint   `json:"userId"`
	UserName string `json:"username"`
	Pfp      string `json:"pfp"`
}
type Message struct {
	User     MessageUser `json:"user"`
	Message  string      `json:"message"`
	Type     messageType `json:"type"`
	DateSent string   `json:"dateSent"`
}
type socketConnection struct {
	conn *websocket.Conn
	user MessageUser
	index int
	mu sync.Mutex
}
func (c *socketConnection) sendMessage(msg interface{}) error{
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("sendin message.....")
	return c.conn.WriteJSON(msg)
}
func (msg *Message) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	return nil
}

func (msg *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(msg)
}

var users = make(map[string][]*socketConnection)
var UserAccs = make(map[string]map[*websocket.Conn]*socketConnection)

var ctx = context.Background()
var publisher = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // use default Addr
	Password: "",               // no password set
	DB:       0,
})

func WSMsgHandler(ws *websocket.Conn, msg string) {
	if conn, ok := UserAccs[ws.Params("id")][ws]; ok {
		publishMessage(conn, msg)
		return
	}

	var UserInfo db.User
	token, err := jwt.Parse(msg, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	err = db.DB.Select("ID, Name, Pfp").Where("id = ?", userId).First(&UserInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("user not found")
	} else {
		fmt.Println(UserInfo.ID)
		fmt.Println(UserInfo.Name)
		fmt.Println(UserInfo.Pfp)
	}
	user := MessageUser{
		UserId:   UserInfo.ID,
		UserName: UserInfo.Name,
		Pfp:      UserInfo.Pfp,
	}
	var connection = &socketConnection{
		conn: ws,
		user: user,
		index: len(users[ws.Params("id")]),
	}
	users[ws.Params("id")] = append(users[ws.Params("id")], connection)
	if _, ok := UserAccs[ws.Params("id")]; ok {
		UserAccs[ws.Params("id")][ws] = connection 
	} else {
		UserAccs[ws.Params("id")] = make(map[*websocket.Conn]*socketConnection)
		UserAccs[ws.Params("id")][ws] = connection 
	}
	SendUsersList(ws.Params("id"))
}

func publishMessage(conn *socketConnection, msg string) {
	finalMsg := makeMessage(conn, chatMsg, msg)
	id := conn.conn.Params("id")
	if id != "" {
		productId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			panic(err)
		}
		dbMsg := db.Message{
			ProductId: uint(productId),
			UserID:    conn.user.UserId,
			Message:   finalMsg.Message,
		}
		db.DB.Create(&dbMsg)
	}
	if err := publisher.Publish(ctx, conn.conn.Params("id"), finalMsg).Err(); err != nil {
		panic(err)
	} 
}
func SendMessage(msg Message, productId string) {
	for _,client := range users[productId] {
		err := client.sendMessage(msg)
		if err != nil {
			panic(err)
		}
	}
}
func makeMessage(conn *socketConnection, msgType messageType, msg string) *Message {
	return &Message{
		User:     conn.user,
		Message:  msg,
		Type:     msgType,
		DateSent: time.Now().Local().String(),
	}
}

func HandleDisconnect(conn *socketConnection) {
	users[conn.conn.Params("id")] = append(users[conn.conn.Params("id")][:conn.index], users[conn.conn.Params("id")][conn.index+1:]...)
	// delete(users[ws.Params("id")], &conn)
	_, exists := UserAccs[conn.conn.Params("id")][conn.conn]
	productId := conn.conn.Params("id")
	if exists {
		delete(UserAccs[conn.conn.Params("id")],conn.conn)
	}
	SendUsersList(productId)

}

func SendUsersList(productId string) {
	var allUsers []MessageUser
	for _, conn := range UserAccs[productId] {
		allUsers = append(allUsers, conn.user)
	}
	fmt.Println("connected users: ", len(allUsers))
	fmt.Println("all users: ", allUsers)
	for _,client  := range users[productId] {
		err := client.sendMessage(struct {
			Users []MessageUser `json:"users"`
			Type  string        `json:"type"`
		}{
			Users: allUsers,
			Type:  string(usersList),
		})
		if err != nil {
			panic(err)
		}
	}
}

type messageType string

const (
	chatMsg   messageType = "chat"
	joinMsg   messageType = "join"
	leaveMsg  messageType = "leave"
	errMsg    messageType = "error"
	usersList messageType = "usersList"
)
