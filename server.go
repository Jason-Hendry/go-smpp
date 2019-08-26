package go_smpp

import (
	"fmt"
	"github.com/pborman/uuid"
	"io"
	"net"
)

type SmppServer struct {
	NodeUuid    uuid.UUID
	Node        string
	bindAddr    string
	OnBind      OnPduCallback
	OnSubmit    OnPduCallback
	connections map[string]*SmppClientConn
}

type SmppClientConn struct {
	conn   *net.TCPConn
	server *SmppServer
	UserId string
}

func (c *SmppClientConn) Write(data []byte) {
	_, err := c.conn.Write(data)
	if err != nil {
		fmt.Println("Error in Write data")
	}
}
func (c *SmppClientConn) WritePdu(pdu Pdu) {
	_, err := c.conn.Write(pdu.Pack())
	if err != nil {
		fmt.Println("Error in Write Pdu")
	}
}

type OnPduCallback func(Pdu, *SmppClientConn)

func handleClient(conn *net.TCPConn, server SmppServer) {
	buf := make([]byte, 20480)
	read := make([]byte, 10240)
	more := make([]byte, 20480)
	more = more[:0]
	buf = buf[:0]
	client := SmppClientConn{conn, &server, ""}
	var smppPdu Pdu
	for {
		fmt.Printf("Node: %s Waiting to read\n", server.Node)
		readLen, err := conn.Read(read[0:])
		if readLen == 0 {
			fmt.Printf("Node: %s Read %d Bytes and probably quit\n", server.Node, readLen)
			return
		}
		fmt.Printf("Node: %s Add %d bytes to buffer\n", server.Node, readLen)
		smppPdu = readMore(&buf, &more, &read, readLen)
		if smppPdu.complete {
			fmt.Printf("Node: %s Got one\n", server.Node)
			processPdu(&client, smppPdu)
		}
		if err == io.EOF {
			fmt.Println("END OF FILE")
		}
	}
}

func readMore(bufP *[]byte, moreP *[]byte, readP *[]byte, readLen int) (smppPdu Pdu) {
	buf := *bufP
	more := *moreP
	read := *readP

	buf = more[:len(more)]               // Leftover from previous packet
	buf = append(buf, read[:readLen]...) // Newly ready data from socket

	smppPdu = RawPdu(buf)
	if smppPdu.complete == true {
		if uint32(len(read)) > smppPdu.commandLength {
			more = more[:0]
			more = append(more, buf[int(smppPdu.commandLength):len(buf)]...)
		}
	} else {
		more = buf[0:len(buf)]
	}
	*moreP = more
	return smppPdu
}

func GetSmppClient(smpp *SmppClientConn) *SmppClientConn {
	return smpp.server.connections[smpp.conn.RemoteAddr().String()]
}

func SetSmppClient(smppClient *SmppClientConn) {
	smppClient.server.connections[smppClient.conn.RemoteAddr().String()] = smppClient
}

func processPdu(client *SmppClientConn, pdu Pdu) {
	server := client.server
	fmt.Printf("Node: %s Got %d\n", server.Node, pdu.commandId)

	switch pdu.commandId {
	case PDU_COMMAND_BIND_TX, PDU_COMMAND_BIND_TRX, PDU_COMMAND_BIND_RX:
		fmt.Printf("Node: %s Bind\n", client.server.Node)
		if server.OnBind != nil {
			server.OnBind(pdu, client)
		} else {
			fmt.Println("Error: Undefined OnBind")
		}
	case PDU_COMMAND_SUBMIT_SM:
		fmt.Printf("Node: %s SUBMIT_SM\n", server.Node)
		if server.OnSubmit != nil {
			if GetSmppClient(client).UserId != "" {
				go server.OnSubmit(pdu, GetSmppClient(client))
			} else {
				var emptyByte []byte
				GetSmppClient(client).Write(emptyByte)
			}
		} else {
			fmt.Println("Error: Undefined OnBind")
		}
	case PDU_COMMAND_ENQUIRE:
		fmt.Printf("Node: %s ENQUIRE\n", server.Node)
		if server.OnSubmit != nil {
			if GetSmppClient(client).UserId != "" {
				go server.OnSubmit(pdu, GetSmppClient(client))
				resp := EnquireLinkResp(pdu)
				GetSmppClient(client).Write(resp.Pack())
			} else {
				fmt.Println("Forbidden")
				var emptyByte []byte
				GetSmppClient(client).Write(emptyByte)
			}
		} else {
			fmt.Println("Error: Undefined OnBind")
		}
	}
}

func (server SmppServer) Start() {
	bindAddr, err := net.ResolveTCPAddr("tcp", server.bindAddr)
	if !HandleError(fmt.Sprintf("Failed to resolve %s", server.bindAddr), err) {
		return
	}
	bind, err := net.ListenTCP("tcp", bindAddr)
	if !HandleError(fmt.Sprintf("Failed to bind to %s", server.bindAddr), err) {
		return
	}

	for {
		conn, err := bind.AcceptTCP()
		HandleError("Failed to accept client", err)
		fmt.Printf("Node: %s Connection: %s\n", server.Node, conn.RemoteAddr().String())
		server.connections[conn.RemoteAddr().String()] = &SmppClientConn{conn: conn, server: &server}
		go handleClient(conn, server)
	}
}

func Server(node string, bindAddr string) *SmppServer {
	var server SmppServer
	server.Node = node
	server.bindAddr = bindAddr
	server.connections = make(map[string]*SmppClientConn)
	return &server
}
