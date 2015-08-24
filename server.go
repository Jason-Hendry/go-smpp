package go_smpp
import (
	"net"
	"fmt"
	"io"
)

func handleClient(conn *net.TCPConn) {
	buf := make([]byte, 20480)
	read := make([]byte, 10240)
	more := make([]byte, 20480)
	more = more[:0]
	buf = buf[:0]
	var smppPdu pdu
	for {
		_,err := conn.Read(read)
		smppPdu = readMore(&buf, &more, &read)
		if smppPdu.complete {
			processPdu(conn, smppPdu)
		}
		if err == io.EOF {

		}

	}
}

func readMore(bufP *[]byte, moreP *[]byte, readP *[]byte) (smppPdu pdu) {
	buf := *bufP
	more := *moreP
	read := *readP

	buf = more[:len(more)] // Leftover from previous packet
	buf = append(buf, read[:len(read)]...) // Newly ready data from socket

	smppPdu = Pdu(buf)
	if smppPdu.complete == true {
		if uint32(len(read)) > smppPdu.command_length {
			more = more[:0]
			more = append(more, buf[len(buf)-int(smppPdu.command_length):]...)
		}
	} else {
		more = append(more, read...)
	}
	*moreP = more;
	return smppPdu
}

func processPdu(conn *net.TCPConn, pdu pdu) {
	switch pdu.command_id {
		case PDU_COMMAND_BIND_TX:
		case PDU_COMMAND_BIND_TRX:
		case PDU_COMMAND_BIND_RX:
			bind(conn, pdu)
	}
}

func bind(conn *net.TCPConn, pdu pdu) {
	fmt.Println("Bind User:%s Pass:%s", string(pdu.system_id), string(pdu.password))

}

func start(bindAddress string) {
	bindAddr,err := net.ResolveTCPAddr("tcp", bindAddress)
	HandleError(fmt.Sprintf("Failed to resolve %s",bindAddress), err)
	bind,err := net.ListenTCP("tcp", bindAddr)
	HandleError(fmt.Sprintf("Failed to bind to %s",bindAddress), err)
	for {
		conn,err := bind.AcceptTCP()
		HandleError("Failed to accept client", err)
		go handleClient(conn)
	}
}


func Server(){


}