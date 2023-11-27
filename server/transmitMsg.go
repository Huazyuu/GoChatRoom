package server

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// TransmitMsg listen global chan message,and forward data
func TransmitMsg() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("TransmitMsg : ", err)
		}
	}()
	for {
		select {
		case msg := <-message:
			strMsg := string(msg)
			ccF.Println(strMsg)
			// Group messaging without `@`
			// only use @ may be hold logic err in face
			if !strings.Contains(strMsg, "@") {
				if strings.Contains(strMsg, "$NUMGO$") {
					//get NumGoroutine
					arr := strings.Split(strMsg, "]says>:")
					if len(arr) == 2 {
						sender := strings.TrimLeft(arr[0], "[")
						for _, v := range onlineUsers {
							if v.name == strings.Trim(sender, " ") {
								v.perC <- []byte("NumGoroutine:" +
									strconv.Itoa(runtime.NumGoroutine()) + "\n")
								break
							}
						}
					} else if strings.Contains(strMsg, "$NUMCONN$") {
						ccF.Println("$NUMCONN$")
						arr := strings.Split(strMsg, "]says>:")
						if len(arr) == 2 {
							sender := strings.TrimLeft(arr[0], "[")
							for _, v := range onlineUsers {
								if v.name == strings.Trim(sender, " ") {
									v.perC <- []byte("NumConn:" +
										strconv.Itoa(len(onlineUsers)) + "\n")
									break
								}
							}
						}
					} else {
						arr := strings.Split(strMsg, "]says>:")
						if len(arr) == 2 {
							sender := strings.TrimLeft(arr[0], "[")
							for _, v := range onlineUsers {
								if v.name == strings.Trim(sender, " ") {
									v.perC <- []byte("group send successfully")
								} else {
									v.perC <- append(msg, []byte("\n")...)
								}
							}
						}
					}
				}
			} else if strings.Contains(strMsg, "@") {
				//private
				arr := strings.Split(strMsg, "@")
				if len(arr) == 2 {
					arr2 := strings.Split(arr[0], "]says>:")
					if len(arr2) == 2 {
						sender := strings.TrimLeft(arr2[0], "[")
						for _, v := range onlineUsers {
							if v.name == strings.Trim(arr[1], " ") {
								v.perC <- []byte(arr[0] + "\n")
							} else if v.name == strings.Trim(sender, " ") {
								v.perC <- []byte("privately send successfully")
							} else {
								v.perC <- []byte("*******************\n")
							}
						}
					}
				}
			} else {
				ccF.Println("消息未识别")
			}
		}
	}
}
