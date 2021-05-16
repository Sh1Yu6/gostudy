package main

func main() {
	server := NewServer("192.168.0.96", 8888)
	server.Start()
}
