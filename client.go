package go_smpp
import (
	"net"
	"fmt"
	"time"
)

type Client struct {
	host string
	username string
	password string
	conn *net.TCPConn
	OnBind OnPduCallback
	OnSubmit OnPduCallback
}

func NewClient(host, username, password string) (*Client) {
	var client Client;
	client.host = host
	client.username = username
	client.password = password
	return &client;
}

func listen(conn *net.TCPConn) {
	resp := make([]byte, 10240);
	for {
		fmt.Printf("Waiting to read\n")
		readLen,err := conn.Read(resp[0:])
		if err != nil {
			fmt.Printf("Something broke\n")
			return
		}
		if readLen == 0 {
			fmt.Printf("Server stopped\n")
			return
		}

		fmt.Printf("Read %d Bytes\n", readLen)
		pduResp := RawPdu(resp)

		switch pduResp.command_id {
		case PDU_COMMAND_SUBMIT_SM+PDU_COMMAND_RESP:
			fmt.Printf("Listener: Submit Message ID: %s\n", string(pduResp.Message_id))
		case PDU_COMMAND_BIND_RX+PDU_COMMAND_RESP,PDU_COMMAND_BIND_TX+PDU_COMMAND_RESP,PDU_COMMAND_BIND_TRX+PDU_COMMAND_RESP:
			fmt.Printf("Listener: BIND Response status: %d\n", pduResp.Command_status)
		case PDU_COMMAND_DELIVER_SM:
			fmt.Printf("Listener: DLR status: %d\n", pduResp.Command_status)
		}
	}
}

func keepAlive(conn *net.TCPConn) {
	for {
		time.Sleep(60 * time.Second)
		enq := EnquireLink(1)
		_, err := conn.Write(enq.Pack())
		if err != nil {
			fmt.Println("Error: Can't send message")
		}
	}
}

func (c *Client) Start() {
	addr,err := net.ResolveTCPAddr("tcp", c.host)
	if !HandleError("Failed to resolve", err) {
		return
	}

	laddr,err := net.ResolveTCPAddr("tcp", ":0")
	if !HandleError("Failed to resolve", err) {
		return
	}

	conn,err := net.DialTCP("tcp", laddr, addr)
	if !HandleError("Failed to connect", err) {
		return
	}
	go listen(conn)
	go keepAlive(conn)

	bind := Bind(1, PDU_COMMAND_BIND_TX, c.username, c.password, "GO-SMPP", 0, 1, 1, "")
	data := bind.Pack()
	fmt.Printf("Wrote %d bytes\n", len(data))

	if conn != nil {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error: Can't send message")
		}
		c.conn = conn
	} else {
		fmt.Println("Error: Can't establish connection")
	}
}

func (c *Client) Send(source,destination,message string) {
	sms := SubmitSM(1, "GO-SMPP", 1, 1, source, 1, 1, destination, PDU_DATA_CODING_LATIN_1, 0, message)
	data := sms.Pack()
	if c.conn != nil {
		_, err := c.conn.Write(data)
		if err != nil {
			fmt.Println("Error: Can't send message")
		}
	} else {
		fmt.Println("Error: Connection unestablished")
	}
}
