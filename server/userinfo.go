package server

// userInfo stores global userinfo
type userInfo struct {
	name    string
	perC    chan []byte // Message channel for each user
	AddUser chan []byte // Broadcast users enter or exit
}
