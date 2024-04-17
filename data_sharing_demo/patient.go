package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	// Recipient string `json:"recipient"`
	Content []byte `json:"content"`
	DataID  string `json:"dataid"`
}

//AES密钥转换成json字符串
func AESkey2json(Aeskey AESKey) []byte {
	//fmt.Println(Aeskey)
	jsonData, err := json.Marshal(Aeskey)
	//fmt.Println(jsonData)
	if err != nil {
		panic(err)
	}

	return jsonData
}

//生成AES密钥用于加密医疗数据
func generateAESkeys(n int) []AESKey {
	AESkeys := make([]AESKey, n)
	for i := 0; i < n; i++ {
		AESkey, err := NewAESKey("patient_key" + strconv.Itoa(i))
		if err != nil {
			fmt.Println("Failed to create AES key:", err)
			return nil
		}
		AESkeys[i] = *AESkey
	}

	return AESkeys
}

func main() {
	//开启服务
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	var name string
	name = "patient1"

	//生成AES密钥用于加密医疗数据
	datanum := 10
	AESkeys := generateAESkeys(datanum)

	//读取数据请求
	go func() {
		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			//处理发送给自己的请求
			if msg.Receiver != name {
				continue
			}
			fmt.Printf("Received from %s: request for data%s\n", msg.Sender, msg.DataID)

			//读取对方公钥，解析公钥格式
			publicKey, err := x509.ParsePKCS1PublicKey(msg.Content)
			if err != nil {
				panic(err)
			}

			id, _ := strconv.Atoi(msg.DataID)
			jsonAESkeylocal := AESkey2json(AESkeys[id])
			//加密对称密钥
			enlocalaeskey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, jsonAESkeylocal)
			if err != nil {
				fmt.Println("encrypto error:", err)
			}

			//发送加密后的对称密钥
			new_msg := Message{Sender: name, Receiver: msg.Sender, Content: enlocalaeskey, DataID: msg.DataID}
			err = conn.WriteJSON(new_msg)
			if err != nil {
				log.Println("Error sending message:", err)
				return
			}
		}
	}()

	var w1 sync.WaitGroup
	w1.Add(1)
	w1.Wait()

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	message := scanner.Text()
	// 	msg := Message{Sender: name, Content: message}
	// 	err := conn.WriteJSON(msg)
	// 	if err != nil {
	// 		log.Println("Error sending message:", err)
	// 		return
	// 	}
	// }
}
