package pc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Pfp  string `json:"pfp"`
}
type Message struct {
	User      MessageUser `json:"user"`
	Message   string      `json:"message"`
	Type      messageType `json:"type"`
	ProductId string      `json:"productId"`
	CreatedAt time.Time   `json:"CreatedAt"`
}
type socketConnection struct {
	conn *websocket.Conn
	user MessageUser
	mu   sync.Mutex
}

func (c *socketConnection) sendMessage(msg interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
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

var (
	users    = make(map[string][]*socketConnection)
	UserAccs = make(map[string]map[*websocket.Conn]*socketConnection)
)

var (
	ctx       = context.Background()
	publisher = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":6379", // use default Addr
		Password: "",                                // no password set
		DB:       0,
	})
)

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
		ws.Close()
		return
	}
	user := MessageUser{
		Id:   UserInfo.ID,
		Name: UserInfo.Name,
		Pfp:  UserInfo.Pfp,
	}
	connection := &socketConnection{
		conn: ws,
		user: user,
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
			UserID:    conn.user.Id,
			Message:   finalMsg.Message,
		}
		err = db.DB.Create(&dbMsg).Error
	}
	if err := publisher.Publish(ctx, "messages", finalMsg).Err(); err != nil {
		panic(err)
	}
}

func SendMessage(msg Message) {
	for _, client := range users[msg.ProductId] {
		err := client.sendMessage(msg)
		if err != nil {
			panic(err)
		}
	}
}

func makeMessage(conn *socketConnection, msgType messageType, msg string) *Message {
	return &Message{
		User:      conn.user,
		Message:   msg,
		Type:      msgType,
		ProductId: conn.conn.Params("id"),
		CreatedAt: time.Now(),
	}
}

func HandleDisconnect(conn *socketConnection) {
	var updatedUsers []*socketConnection
	for _, user := range users[conn.conn.Params("id")] {
		if user.conn != conn.conn {
			updatedUsers = append(updatedUsers, user)
		}
	}
	users[conn.conn.Params("id")] = updatedUsers
	_, exists := UserAccs[conn.conn.Params("id")][conn.conn]
	productId := conn.conn.Params("id")
	if exists {
		delete(UserAccs[conn.conn.Params("id")], conn.conn)
	}
	SendUsersList(productId)
}

func SendUsersList(productId string) {
	var allUsers []MessageUser
	for _, conn := range UserAccs[productId] {
		allUsers = append(allUsers, conn.user)
	}
	for _, client := range users[productId] {
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
