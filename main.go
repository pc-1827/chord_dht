package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a new node
	node := NewNode()

	// Read port number from user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the port number to run the server on: ")
	portStr, _ := reader.ReadString('\n')
	portStr = strings.TrimSpace(portStr)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid port number")
	}
	node.Port = port

	// Read address of existing node to join
	fmt.Print("Enter address of existing node to join (or leave blank to start a new network): ")
	joinAddr, _ := reader.ReadString('\n')
	joinAddr = strings.TrimSpace(joinAddr)

	// Initialize node
	node.InitializeNode()

	if joinAddr != "" {
		// Join existing network
		err := node.JoinNetwork(joinAddr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Start new network (set successor and predecessor to self)
		node.Successor = node.Address()
		node.Predecessor = node.Address()
	}

	// Set up HTTP handlers
	http.HandleFunc("/find_successor", node.FindSuccessorHandler)
	http.HandleFunc("/get_predecessor", node.GetPredecessorHandler)
	http.HandleFunc("/notify", node.NotifyHandler)
	http.HandleFunc("/put", node.PutHandler)
	http.HandleFunc("/get", node.GetHandler)

	// Start HTTP server
	address := fmt.Sprintf(":%d", node.Port)
	fmt.Printf("Node running on port %d\n", node.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
