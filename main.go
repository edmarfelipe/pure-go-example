package main

import "log"

const (
	addr = ":3000"
	url  = "https://api.github.com/users/"
)

func main() {
	svc := NewLoggingService(
		NewUserInfoService(url),
	)

	apiServer := NewAPIServer(svc)

	log.Println("serving http on " + addr)
	log.Fatal(apiServer.Start(addr))
}
