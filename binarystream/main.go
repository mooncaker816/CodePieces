package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Message struct {
	EventCode   uint8
	IsSystemMsg bool
	Room        string
	Dst         string
	Src         string
	Payload     []byte
}

var p = fmt.Println
var mytpl = template.Must(template.New("index.html").ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":8080", nil)
	// DecodeMessage(EncodeMessage(msg))
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	msg := &Message{
		EventCode:   255,
		IsSystemMsg: true,
		Room:        strings.Repeat("room1", 60),
		Dst:         "",
		Src:         "mingle",
		Payload:     []byte("Hello,小Q👿"),
	}
	DecodeMessage(EncodeMessage(msg))
	if err := mytpl.Execute(w, EncodeMessage(msg)); err != nil {
		fmt.Println(err)
	}
}

func DecodeMessage(data []byte) *Message {
	msg := &Message{}
	buf := bytes.NewBuffer(data)
	fmt.Println(len(data))
	msg.EventCode = uint8(buf.Next(1)[0])
	msg.IsSystemMsg, _ = strconv.ParseBool(string(buf.Next(1)))
	ss := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		fmt.Println(i)
		ss[i] = make([]byte, int(binary.BigEndian.Uint32(buf.Next(4))))
		binary.Read(buf, binary.BigEndian, &ss[i])
	}
	msg.Room = string(ss[0])
	msg.Dst = string(ss[1])
	msg.Src = string(ss[2])
	msg.Payload = ss[3]
	fmt.Printf("%+v", msg)
	return msg
}

func EncodeMessage(msg *Message) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, msg.EventCode)
	binary.Write(buf, binary.BigEndian, msg.IsSystemMsg)
	binary.Write(buf, binary.BigEndian, uint32(len(msg.Room)))
	binary.Write(buf, binary.BigEndian, []byte(msg.Room))
	binary.Write(buf, binary.BigEndian, uint32(len(msg.Dst)))
	binary.Write(buf, binary.BigEndian, []byte(msg.Dst))
	binary.Write(buf, binary.BigEndian, uint32(len(msg.Src)))
	binary.Write(buf, binary.BigEndian, []byte(msg.Src))
	binary.Write(buf, binary.BigEndian, uint32(len(msg.Payload)))
	binary.Write(buf, binary.BigEndian, msg.Payload)
	//binary.Write(buf, binary.BigEndian, msg.EventCode)
	fmt.Println(buf.Bytes())
	return buf.Bytes()
}
