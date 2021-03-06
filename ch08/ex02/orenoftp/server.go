package orenoftp

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/exp/errors/fmt"
)

// ディレクトリを変更する cd
// ディレクトリを列挙する ls
// ファイルの内容を送り出す get
// 接続を閉じる close
// put

type Server struct {
	Addr string
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fc := NewFtpConn(c, "/", wd)
	fc.Welcome()

	s := bufio.NewScanner(fc.Conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "USER":
			fc.User()
		case "PASS":
			fc.Pass()
		case "SYST":
			fc.Type()
		case "CWD": // cd
			fc.Cwd(args)
		case "LIST": // ls
			fc.List(args)
		case "PWD":
			fc.Pwd()
		case "PORT":
			fc.Port(args)
		case "LPRT":
			fc.Lprt(args)
		case "RETR":
			fc.Retr(args)
		case "STOR":
			fc.Stor(args)
		case "QUIT":
			fc.Quit()
		default:
			log.Println(fmt.Sprintf("unsupported command: %s", command))
			fmt.Fprint(fc.Conn, "502 Command not implemented.\n")
		}
	}
}
