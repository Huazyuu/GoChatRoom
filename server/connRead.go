package server

import "net"

func connRead(conn net.Conn, overtime chan bool) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		thisUser := onlineUsers[conn.RemoteAddr().String()].name
		if n == 0 {
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.AddUser <- []byte("user[" + thisUser + "] is exit\n")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return
		}
		if err != nil {
			ccF.Println("connRead false err:", err)
			return
		}

		var msg []byte
		if buf[0] != 10 {
			// \n->10 enter->13
			msg = append([]byte("["+thisUser+"] >"), buf[:n-2]...) //remove \n & enter
		} else {
			msg = nil
		}
		// send msg to msg chan
		overtime <- true
		message <- msg
		LogToDb(string(msg), thisUser)
	}
}
