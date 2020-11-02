package chat

import (
	"log"
	"net"
)

func Start() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("ooops: %s", err.Error())
	}
	defer listener.Close()
	log.Println("starting at :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept conn: %s\n", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
