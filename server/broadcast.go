package server

// broadcast if user is on Online,then broadcast
func broadcast(userName string) {
	for _, v := range onlineUsers {
		v.AddUser <- []byte("user[" + userName + "] enters the chat room\n")
	}
}
