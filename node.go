package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
)

type Node struct {
	ID          uint32
	Port        int
	Successor   string
	Predecessor string
}

func NewNode() *Node {
	return &Node{}
}

func (node *Node) InitializeNode() {
	// Assign an ID based on the node's address
	address := node.Address()
	node.ID = hash(address)
	// Initialize data and finger table storage (write to files)
	// For example:
	// - Create a file to store data: data_<node.ID>.json
	// - Create a file to store finger table: fingers_<node.ID>.json
}

func (node *Node) Address() string {
	return fmt.Sprintf("localhost:%d", node.Port)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (node *Node) JoinNetwork(joinAddr string) error {
	// Implement logic to join the network
	// Send a request to joinAddr to find your successor and predecessor
	// Update your successor and predecessor
	return nil
}

// HTTP handler methods (to be implemented)

func (node *Node) FindSuccessorHandler(w http.ResponseWriter, r *http.Request) {
	// Handle /find_successor requests
}

func (node *Node) GetPredecessorHandler(w http.ResponseWriter, r *http.Request) {
	// Handle /get_predecessor requests
}

func (node *Node) NotifyHandler(w http.ResponseWriter, r *http.Request) {
	// Handle /notify requests
}

func (node *Node) PutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle /put requests for storing key-value pairs
}

func (node *Node) GetHandler(w http.ResponseWriter, r *http.Request) {
	// Handle /get requests for retrieving values by key
}
