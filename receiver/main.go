package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func main() {
	receiver, err := net.Dial("tcp", fmt.Sprintf("%s:8000", os.Getenv("SENDER_IP")))
	if err != nil {
		panic(err)
	}
	defer receiver.Close()

	fmt.Println("Successfully dialed to IP:", os.Getenv("SENDER_IP"))

	result, err := receive(receiver)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func receive(receiver net.Conn) (string, error) {
	result, _ := ioutil.ReadAll(receiver)

	metadata := strings.SplitN(string(result), "::filename::", 2)

	file, err := os.Create(metadata[0])
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := strings.Join(metadata[1:], "")

	_, err = io.Copy(file, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("File %s received!", metadata[0]), nil
}
