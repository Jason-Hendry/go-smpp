package go_smpp
import (
)

const PDU_COMMAND_BIND_RX    = 1
const PDU_COMMAND_BIND_TX    = 2
const PDU_COMMAND_QUERY_SM   = 3
const PDU_COMMAND_SUBMIT_SM  = 4
const PDU_COMMAND_DELIVER_SM = 5
const PDU_COMMAND_UNBIND     = 6
const PDU_COMMAND_REPLACE_SM = 7
const PDU_COMMAND_CANCEL_SM  = 8
const PDU_COMMAND_BIND_TRX   = 9

const PDU_COMMAND_ENQUIRE    = 0x15

const PDU_COMMAND_SUBMIT_MULTI = 0x21
const PDU_COMMAND_ALERT        = 0x0102
const PDU_COMMAND_DATA_SM      = 0x0103

const PDU_COMMAND_RESP = 0x80000000



type pdu struct {
	command_length uint32
	command_id uint32
	command_status uint32
	sequence_number uint32

	service_type int
	source_addr_ton int
	source_addr_npi int
	source_addr int
	dest_addr_ton int
	dest_addr_npi int
	dest_addr int
	esm_class int
	protocol_id int
	priority_flag int
	schedule_delivery_time int
	validity_period int
	registered_delivery int
	replace_if_present_flag int
	data_coding int
	sm_default_msg_id int
	sm_length int
	short_message int

	complete bool

	system_id []byte
	password []byte
	system_type []byte

	interface_version int
	addr_ton int
	addr_npi int

	address_range []byte
}

func atoi(byte []byte) (uint32) {
	return uint32(byte[0]) << 24 + uint32(byte[1]) << 16 + uint32(byte[2]) << 8 + uint32(byte[3])
}
func cOctetString(raw []byte, offset int, max int) ([]byte, int) {
	var result []byte
	for i := offset;i < len(raw); i++ {
		if raw[i] == 0x00 || len(result) == max {
			return result,i+1
		}
		result = append(result, raw[i])
	}
	return []byte(""),offset+max
}

func Pdu(raw []byte) (pdu) {
	var output pdu

	var rawLen = len(raw)


	output.command_length = atoi(raw[0:4])
	output.command_id = atoi(raw[4:8])
	output.command_status = atoi(raw[8:12])
	output.sequence_number = atoi(raw[12:16])



	output.complete = false
	if uint32(rawLen) >= output.command_length {
		output.complete = true

		switch output.command_id {
		case PDU_COMMAND_BIND_TX,PDU_COMMAND_BIND_RX,PDU_COMMAND_BIND_TRX:
			bindBody(raw,&output)

		}

	}
	return output
}

func bindBody(raw []byte, output *pdu) {
	offset := 16
	output.system_id,offset = cOctetString(raw, offset, 16)
	output.password,offset = cOctetString(raw, offset, 9)
	output.system_type,offset = cOctetString(raw, offset, 13)

	output.interface_version = int(raw[offset])
	offset++
	output.addr_ton = int(raw[offset])
	offset++
	output.addr_npi = int(raw[offset])
	offset++

	output.address_range,_ = cOctetString(raw, offset, 41)
}