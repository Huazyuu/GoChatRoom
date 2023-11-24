package server

import "github.com/fatih/color"

//global variable

// global msg chan using for handle msgs from individual clients
var message = make(chan []byte)

// store online user msg
var onlineUsers = make(map[string]userInfo)

var ccS = color.New(color.FgGreen, color.Bold)
var ccF = color.New(color.FgRed, color.Bold)
