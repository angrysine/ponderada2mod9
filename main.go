package main

func main() {
	go Broker()
	go Publisher()
	go Subscriber()
	select {}
}
