package main

import (
	"fmt"
	"net/http"
)

// Adapter Design pattern
type Computer interface {
	InsertIntoPort() string
}
type Mac struct{}

func (m *Mac) InsertIntoPort() string {
	return "Lightning connector is plugged into Mac machine."
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() string {
	return "USB connector is plugged into Windows machine."
}

// adapter
type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoPort() string {
	return "Adapter converts USB signal to Lightning.\n" + w.windowsMachine.InsertIntoUSBPort()
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	client := &Client{}

	var com Computer

	switch r.URL.Path {
	case "/mac":
		com = &Mac{}
	case "/windows":
		windowsMachine := &Windows{}
		com = &WindowsAdapter{windowsMachine: windowsMachine}
	default:
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response := client.InsertConnectorIntoComputer(com)
	fmt.Fprintf(w, "%s", response)
}

type Client struct{}

func (c *Client) InsertConnectorIntoComputer(com Computer) string {
	return "Client inserts connector into computer.\n" + com.InsertIntoPort()
}

func main() {
	http.HandleFunc("/mac", ClientHandler)
	http.HandleFunc("/windows", ClientHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
