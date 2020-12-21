package irc

import (
	"go-bot/app/message"

	"bufio"
	"log"
	"net"
	"net/textproto"
)

type IRC struct {
	Conn net.Conn
}

func Listen(conn net.Conn) {
	irc := IRC{conn}

	irc.Read()
}

func (i *IRC) Read() {
	tp := textproto.NewReader(bufio.NewReader(i.Conn))

	for {
		msg, err := tp.ReadLine()

		if err != nil {
			log.Panicln(err.Error())
		}

		go message.ReadAndParse(msg)
	}
}
