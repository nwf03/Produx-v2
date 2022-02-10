package pc

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"tutorial/db"

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
	DateSent time.Time   `json:"dateSent"`
}

var users = make(map[string]map[*websocket.Conn]MessageUser)
var userAccs = make(map[string]map[uint]MessageUser)

func WSMsgHandler(ws *websocket.Conn, msg string) {
	if _, ok := users[ws.Params("id")][ws]; ok {
		sendMessage(ws, msg)
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
	//todo make a dictionary of unique users and their connections
	if _, ok := userAccs[ws.Params("id")]; ok {
		if _, exists := userAccs[ws.Params("id")][uint(userId)]; exists {
			users[ws.Params("id")][ws] = userAccs[ws.Params("id")][uint(userId)]
			SendUsersList(ws)
			return
		}
	}
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
	if _, ok := users[ws.Params("id")]; ok {
		users[ws.Params("id")][ws] = user
	} else {
		users[ws.Params("id")] = make(map[*websocket.Conn]MessageUser)
		users[ws.Params("id")][ws] = user
	}
	if _, ok := userAccs[ws.Params("id")]; ok {
		userAccs[ws.Params("id")][uint(userId)] = user
	} else {
		userAccs[ws.Params("id")] = make(map[uint]MessageUser)
		userAccs[ws.Params("id")][uint(userId)] = user
	}
	SendUsersList(ws)
}

func sendMessage(ws *websocket.Conn, msg string) {
	finalMsg := makeMessage(ws, chatMsg, msg)
	id := ws.Params("id")

	if id != "" {
		productId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			panic(err)
		}
		dbMsg := db.Message{
			ProductId: uint(productId),
			UserID:    users[id][ws].UserId,
			Message:   finalMsg.Message,
		}
		db.DB.Create(&dbMsg)
	}
	for client := range users[id] {
		err := client.WriteJSON(finalMsg)
		if err != nil {
			panic(err)
		}
	}
}

func makeMessage(ws *websocket.Conn, msgType messageType, msg string) *Message {
	user := users[ws.Params("id")][ws]
	return &Message{
		User:     user,
		Message:  msg,
		Type:     msgType,
		DateSent: time.Now(),
	}
}

func HandleDisconnect(ws *websocket.Conn) {
	user, ok := users[ws.Params("id")][ws]
	if ok {
		delete(users[ws.Params("id")], ws)
	}
	_, exists := userAccs[ws.Params("id")][user.UserId]
	if exists {
		delete(userAccs[ws.Params("id")], user.UserId)
	}
	SendUsersList(ws)

}

func SendUsersList(ws *websocket.Conn) {
	var allUsers []MessageUser
	for _, user := range userAccs[ws.Params("id")] {
		allUsers = append(allUsers, user)
	}
	fmt.Println("connected users: ", len(allUsers))
	fmt.Println("all users: ", allUsers)
	for client := range users[ws.Params("id")] {
		err := client.WriteJSON(struct {
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
