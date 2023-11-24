package server

import (
	"net"
	"time"
)

// HandleConnect handle conn request
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	// chan overtime
	overtime := make(chan bool)
	bufUName := make([]byte, 4096)
	n, err := conn.Read(bufUName) //get userName
	if err != nil {
		ccF.Println("HandleConnect conn.Read(bufUName) failed err: ", err)
		return
	}

	userName := string(bufUName[:n])
	perC := make(chan []byte)
	perAddUser := make(chan []byte)
	user := userInfo{
		name:    userName,
		perC:    perC,
		AddUser: perAddUser,
	}
	onlineUsers[conn.RemoteAddr().String()] = user

	// new user broadcast after connecting
	go broadcast(userName)

	//listen client self-MsgChan, unique to each client
	go connWrite(conn, user)

	//loop read msg from client
	go connRead(conn, overtime)

	for {
		select {
		case <-overtime:
		case <-time.After(time.Second * 300):
			_, _ = conn.Write([]byte("overtime,exit app"))
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.AddUser <- []byte("user[" + thisUser + "] is exit\n")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return
		}
	}

}
