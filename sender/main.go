package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing filename as command line argument.")
		os.Exit(1)
	}
	filename := os.Args[1]

	sender, err := net.Listen("tcp", fmt.Sprintf("%s:8000", os.Getenv("SENDER_IP")))
	if err != nil {
		panic(err)
	}
	defer sender.Close()
	fmt.Println("Listening to address:", os.Getenv("SENDER_IP"))

	connection, err := sender.Accept()
	if err != nil {
		panic(err)
	}

	result, _ := send(connection, filename)

	fmt.Println(result)
}

func send(sender net.Conn, filename string) (string, error) {
	fmt.Println("Sending file:", filename)

	encFilename := fmt.Sprintf("%s::filename::", filename)

	sender.Write([]byte(encFilename))

	file, err := os.Open(strings.TrimSpace(filename))
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(sender, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s successfully sent!", filename), nil
}
