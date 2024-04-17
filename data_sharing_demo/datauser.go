package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	// Recipient string `json:"recipient"`
	Content []byte `json:"content"`
	DataID  string `json:"dataid"`
}

// 生成RSA密钥（用于后续传输）
func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// 生成RSA私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// 从私钥中提取公钥
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

//json字符串转换为AES密钥
func json2AESkey(j []byte) AESKey {
	//fmt.Println(Aeskey)
	Aeskey := AESKey{}
	err := json.Unmarshal(j, &Aeskey)
	//fmt.Println(jsonData)
	if err != nil {
		panic(err)
	}

	return Aeskey
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
	name = "datauser1"

	// fmt.Println("111")

	//生成RSA密钥
	privateKey, publicKey, err := generateRSAKeys()
	if err != nil {
		fmt.Println("Failed to create AES key:", err)
		return
	}
	// privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)

	//读取服务端消息
	go func() {
		for {
			var msg Message
			Aeskey := AESKey{}
			//读取消息
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			//处理发送给自己的消息
			if msg.Receiver != name {
				continue
			}

			// fmt.Printf("Received from %s: %s\n", msg.Sender, msg.Receiver)

			//私钥解密密文
			jsonAESkeylocal, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(msg.Content))

			//解密出的AES密钥
			Aeskey = json2AESkey(jsonAESkeylocal)

			fmt.Printf("Received from %s: %s\n", msg.Sender, Aeskey.KeyID)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dataid := scanner.Text()
		fmt.Printf("Requesting for data%s...\n", dataid)
		msg := Message{
			Sender:   name,
			Receiver: "patient1",
			Content:  publicKeyBytes,
			DataID:   dataid,
		}
		// fmt.Printf("%s", msg.Receiver)
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
}
