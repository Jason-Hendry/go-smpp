package go_smpp
import (
	"fmt"
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

const PDU_OPT_PARAM_DEST_ADDR_SUBUNIT 				= 0x0005 // GSM
const PDU_OPT_PARAM_DEST_NETWORK_TYPE 				= 0x0006 // Generic
const PDU_OPT_PARAM_DEST_BEARER_TYPE 				= 0x0007 // Generic
const PDU_OPT_PARAM_DEST_TELEMATICS_ID 				= 0x0008 // GSM
const PDU_OPT_PARAM_SOURCE_ADDR_SUBUNIT 			= 0x000D // GSM
const PDU_OPT_PARAM_SOURCE_NETWORK_TYPE 			= 0x000E // Generic
const PDU_OPT_PARAM_SOURCE_BEARER_TYPE 				= 0x000F // Generic
const PDU_OPT_PARAM_SOURCE_TELEMATICS_ID 			= 0x0010 // GSM
const PDU_OPT_PARAM_QOS_TIME_TO_LIVE 				= 0x0017 // Generic
const PDU_OPT_PARAM_PAYLOAD_TYPE 					= 0x0019 // Generic
const PDU_OPT_PARAM_ADDITIONAL_STATUS_INFO_TEXT 	= 0x001D // Generic
const PDU_OPT_PARAM_RECEIPTED_MESSAGE_ID 			= 0x001E // Generic
const PDU_OPT_PARAM_MS_MSG_WAIT_FACILITIES 			= 0x0030 // GSM
const PDU_OPT_PARAM_PRIVACY_INDICATOR 				= 0x0201 // CDMA, TDMA
const PDU_OPT_PARAM_SOURCE_SUBADDRESS 				= 0x0202 // CDMA, TDMA
const PDU_OPT_PARAM_DEST_SUBADDRESS 				= 0x0203 // CDMA, TDMA
const PDU_OPT_PARAM_USER_MESSAGE_REFERENCE 			= 0x0204 // Generic
const PDU_OPT_PARAM_USER_RESPONSE_CODE 				= 0x0205 // CDMA, TDMA
const PDU_OPT_PARAM_SOURCE_PORT 					= 0x020A // Generic
const PDU_OPT_PARAM_DESTINATION_PORT 				= 0x020B // Generic
const PDU_OPT_PARAM_SAR_MSG_REF_NUM 				= 0x020C // Generic
const PDU_OPT_PARAM_LANGUAGE_INDICATOR 				= 0x020D // CDMA, TDMA
const PDU_OPT_PARAM_SAR_TOTAL_SEGMENTS 				= 0x020E // Generic
const PDU_OPT_PARAM_SAR_SEGMENT_SEQNUM 				= 0x020F // Generic
const PDU_OPT_PARAM_SC_INTERFACE_VERSION 			= 0x0210 // Generic
const PDU_OPT_PARAM_CALLBACK_NUM_PRES_IND 			= 0x0302 // TDMA
const PDU_OPT_PARAM_CALLBACK_NUM_ATAG 				= 0x0303 // TDMA
const PDU_OPT_PARAM_NUMBER_OF_MESSAGES 				= 0x0304 // CDMA
const PDU_OPT_PARAM_CALLBACK_NUM 					= 0x0381 // CDMA, TDMA, GSM, iDEN
const PDU_OPT_PARAM_DPF_RESULT 						= 0x0420 // Generic
const PDU_OPT_PARAM_SET_DPF 						= 0x0421 // Generic
const PDU_OPT_PARAM_MS_AVAILABILITY_STATUS 			= 0x0422 // Generic
const PDU_OPT_PARAM_NETWORK_ERROR_CODE 				= 0x0423 // Generic
const PDU_OPT_PARAM_MESSAGE_PAYLOAD 				= 0x0424 // Generic
const PDU_OPT_PARAM_DELIVERY_FAILURE_REASON 		= 0x0425 // Generic
const PDU_OPT_PARAM_MORE_MESSAGES_TO_SEND 			= 0x0426 // GSM
const PDU_OPT_PARAM_MESSAGE_STATE 					= 0x0427 // Generic
const PDU_OPT_PARAM_USSD_SERVICE_OP 				= 0x0501 // GSM (USSD)
const PDU_OPT_PARAM_DISPLAY_TIME 					= 0x1201 // CDMA, TDMA
const PDU_OPT_PARAM_SMS_SIGNAL 						= 0x1203 // TDMA
const PDU_OPT_PARAM_MS_VALIDITY 					= 0x1204 // CDMA, TDMA
const PDU_OPT_PARAM_ALERT_ON_MESSAGE_DELIVERY 		= 0x130C // CDMA
const PDU_OPT_PARAM_ITS_REPLY_TYPE 					= 0x1380 // CDMA
const PDU_OPT_PARAM_ITS_SESSION_INFO 				= 0x1383 // CDMA

const PDU_TON_UNKNOWN           = 0
const PDU_TON_INTERNATIONAL     = 1
const PDU_TON_NATIONAL          = 2
const PDU_TON_NETWORK_SPECIFIC  = 3
const PDU_TON_SUBSCRIBER_NUMBER = 4
const PDU_TON_ALPHANUMERIC      = 5
const PDU_TON_ABBREVIATED       = 6

const PDU_NPI_UNKNOWN           = 0
const PDU_NPI_ISDN              = 1
const PDU_NPI_DATA              = 3
const PDU_NPI_TELEX             = 4
const PDU_NPI_LAND_MOBILE       = 6
const PDU_NPI_NATIONAL          = 8
const PDU_NPI_PRIVATE           = 9
const PDU_NPI_ERMES             = 10
const PDU_NPI_INTERNET          = 14
const PDU_NPI_WAP               = 18

const PDU_DELIVERY_RECEIPT_SMSC = 1 // SMSC delivery receipt

type parameter struct {
	tag uint16
	length uint16
	value []byte
}

type pdu struct {
	command_length uint32
	command_id uint32
	Command_status uint32
	sequence_number uint32

	service_type []byte
	source_addr_ton int
	source_addr_npi int
	source_addr []byte
	dest_addr_ton int
	dest_addr_npi int
	dest_addr []byte
	esm_class int
	protocol_id int
	priority_flag int
	schedule_delivery_time []byte
	validity_period []byte
	registered_delivery int
	replace_if_present_flag int
	data_coding int
	sm_default_msg_id int
	sm_length int
	short_message []byte

	complete bool

	system_id []byte
	password []byte
	system_type []byte

	interface_version int
	addr_ton int
	addr_npi int

	address_range []byte

	Message_id []byte

	optionalParameters []parameter
}

func (p *pdu) Pack() ([]byte) {
	var packet []byte
	var body []byte

	switch p.command_id {
	case PDU_COMMAND_BIND_TX,PDU_COMMAND_BIND_RX,PDU_COMMAND_BIND_TRX:
		body = packBindBody(*p)
	case PDU_COMMAND_SUBMIT_SM,PDU_COMMAND_DATA_SM,PDU_COMMAND_DELIVER_SM:
		body = packSubmitBody(*p)
	}

	p.command_length = uint32(len(body) + 16)

	appendInteger(&packet, p.command_length)
	appendInteger(&packet, p.command_id)
	appendInteger(&packet, p.Command_status)
	appendInteger(&packet, p.sequence_number)


	packet = append(packet, body...)

	return packet
}

func (p *pdu) getOptionalParameter(tag uint16) ([]byte) {
	for i := range p.optionalParameters {
		if p.optionalParameters[i].tag == tag {
			return p.optionalParameters[i].value
		}
	}
	return nil
}
func (p *pdu) hasOptionalParameter(tag uint16) (bool) {
	for i := range p.optionalParameters {
		if p.optionalParameters[i].tag == tag {
			return true
		}
	}
	return false
}
func (p *pdu) updateOptionalParameter(tag uint16, value []byte) {
	for i := range p.optionalParameters {
		if p.optionalParameters[i].tag == tag {
			p.optionalParameters[i].length = uint16(len(value))
			p.optionalParameters[i].value = value
		}
	}
}
func (p *pdu) addOptionalParameter(tag uint16, value []byte) {
	var param parameter
	param.tag = tag
	param.length = uint16(len(value))
	param.value = value
	p.optionalParameters = append(p.optionalParameters, param)
}

func (p *pdu) setOptionalParameter(tag uint16, value []byte) {
	if p.hasOptionalParameter(tag) {
		p.updateOptionalParameter(tag, value)
	} else {
		p.addOptionalParameter(tag, value)
	}
}


func unpackInteger(byte []byte) (uint32) {
	return uint32(byte[0]) << 24 + uint32(byte[1]) << 16 + uint32(byte[2]) << 8 + uint32(byte[3])
}
func packInteger(num uint32) ([]byte) {
	return []byte{byte(num >> 24),byte(num >> 16),byte(num >> 8),byte(num)}
}

func UnpackCOctetString(raw []byte, offset int, max int) ([]byte, int) {
	var result []byte
	for i := offset;i <= len(raw); i++ {
		if raw[i] == 0x00 || len(result) == max {
			return result,i+1
		}
		result = append(result, raw[i])
	}
	return []byte(""),offset+max
}
func UnpackOctetString(raw []byte, offset int, length int) ([]byte, int) {
	var result []byte
	for i := offset;i <= len(raw); i++ {
		if len(result) == length {
			return result,i
		}
		result = append(result, raw[i])
	}
	return []byte(""),offset+length
}

func Pdu(raw []byte) (pdu) {
	var output pdu
	var rawLen = len(raw)

	output.command_length = unpackInteger(raw[0:4])
	output.command_id = unpackInteger(raw[4:8])
	output.Command_status = unpackInteger(raw[8:12])
	output.sequence_number = unpackInteger(raw[12:16])

	output.complete = false
	if uint32(rawLen) >= output.command_length {
		output.complete = true

		switch output.command_id {
		case PDU_COMMAND_BIND_TX,PDU_COMMAND_BIND_RX,PDU_COMMAND_BIND_TRX:
			unpackBindBody(raw,&output)
		case PDU_COMMAND_SUBMIT_SM,PDU_COMMAND_DATA_SM,PDU_COMMAND_DELIVER_SM:
			unpackSubmitBody(raw,&output)
		case PDU_COMMAND_BIND_TX+PDU_COMMAND_RESP,PDU_COMMAND_BIND_RX+PDU_COMMAND_RESP,PDU_COMMAND_BIND_TRX+PDU_COMMAND_RESP:
			unpackBindBodyResp(raw,&output)
		case PDU_COMMAND_SUBMIT_SM+PDU_COMMAND_RESP,PDU_COMMAND_DATA_SM+PDU_COMMAND_RESP,PDU_COMMAND_DELIVER_SM+PDU_COMMAND_RESP:
			unpackSubmitBodyResp(raw,&output)
		}

	}
	return output
}

func unpackBindBody(raw []byte, output *pdu) {
	offset := 16
	output.system_id,offset = UnpackCOctetString(raw, offset, 16)
	output.password,offset = UnpackCOctetString(raw, offset, 9)
	output.system_type,offset = UnpackCOctetString(raw, offset, 13)

	output.interface_version = int(raw[offset])
	offset++
	output.addr_ton = int(raw[offset])
	offset++
	output.addr_npi = int(raw[offset])
	offset++

	output.address_range,_ = UnpackCOctetString(raw, offset, 41)
}


func unpackBindBodyResp(raw []byte, output *pdu) {
	offset := 16
	output.system_id,offset = UnpackCOctetString(raw, offset, 16)
	unpackTLVs(raw, output, offset)
}

func unpackSubmitBody(raw []byte, output *pdu) {
	offset := 16

	output.service_type,offset = UnpackCOctetString(raw, offset, 6)

	output.source_addr_ton = int(raw[offset])
	offset++
	output.source_addr_npi = int(raw[offset])
	offset++
	output.source_addr,offset = UnpackCOctetString(raw, offset, 21)

	output.dest_addr_ton = int(raw[offset])
	offset++
	output.dest_addr_npi = int(raw[offset])
	offset++
	output.dest_addr,offset = UnpackCOctetString(raw, offset, 21)

	output.esm_class = int(raw[offset])
	offset++
	output.protocol_id = int(raw[offset])
	offset++
	output.priority_flag = int(raw[offset])
	offset++

	output.schedule_delivery_time,offset = UnpackCOctetString(raw, offset, 17)
	output.validity_period,offset = UnpackCOctetString(raw, offset, 17)

	output.registered_delivery = int(raw[offset])
	offset++
	output.replace_if_present_flag = int(raw[offset])
	offset++
	output.data_coding = int(raw[offset])
	offset++
	output.sm_default_msg_id = int(raw[offset])
	offset++
	output.sm_length = int(raw[offset])
	offset++

	output.short_message,offset = UnpackOctetString(raw, offset, output.sm_length)

	unpackTLVs(raw, output, offset)
}

func unpackSubmitBodyResp(raw []byte, output *pdu) {
	offset := 16

	output.Message_id,offset = UnpackCOctetString(raw, offset, 65)
}

func unpackTLVs(raw []byte, output *pdu, offset int) {
	for uint32(offset) < output.command_length {
		offset = unpackTLV(raw, output, offset)
	}
}
func unpackTLV(raw []byte, output *pdu, offset int) (int) {
	var param parameter
	fmt.Println(raw[offset], raw[offset+1])

	param.tag = uint16(raw[offset] << 8 + raw[offset+1])
	param.length = uint16(raw[offset+2] << 8 + raw[offset+3])
	offset += 4
	param.value,offset = UnpackOctetString(raw, offset, int(param.length))
	output.optionalParameters = append(output.optionalParameters, param)
	return offset
}

func packTLVs(body *[]byte, p pdu) {
	var tag, length uint16
	var value,buf []byte

	for i := range p.optionalParameters {
		tag = p.optionalParameters[i].tag
		length = p.optionalParameters[i].length
		value = p.optionalParameters[i].value
		buf = []byte{byte(tag >> 8),byte(tag),byte(length >> 8),byte(length)}
		buf = append(buf, value...)
		*body = append(*body, buf...)
	}
}

func appendCOctetString(buf *[]byte, str []byte) {
	*buf = append(*buf, str...)
	*buf = append(*buf, 0x00) // Null Ternimate strings
}
func appendInteger(buf *[]byte, num uint32) {
	str := packInteger(num)
	*buf = append(*buf, str...)
}

func packBindBody(pdu pdu) ([]byte) {
	var body []byte

	appendCOctetString(&body, pdu.system_id)
	appendCOctetString(&body, pdu.password)
	appendCOctetString(&body, pdu.system_type)

	body = append(body, byte(pdu.interface_version))
	body = append(body, byte(pdu.addr_ton))
	body = append(body, byte(pdu.addr_npi))

	appendCOctetString(&body, pdu.address_range)

	return body
}

func packSubmitBody(pdu pdu) ([]byte) {
	var body []byte

	appendCOctetString(&body, pdu.service_type)

	body = append(body, byte(pdu.source_addr_ton))
	body = append(body, byte(pdu.source_addr_npi))
	appendCOctetString(&body, pdu.source_addr)

	body = append(body, byte(pdu.dest_addr_ton))
	body = append(body, byte(pdu.dest_addr_npi))
	appendCOctetString(&body, pdu.dest_addr)

	body = append(body, byte(pdu.esm_class))
	body = append(body, byte(pdu.protocol_id))
	body = append(body, byte(pdu.priority_flag))

	appendCOctetString(&body, pdu.schedule_delivery_time)
	appendCOctetString(&body, pdu.validity_period)

	body = append(body, byte(pdu.registered_delivery))
	body = append(body, byte(pdu.replace_if_present_flag))
	body = append(body, byte(pdu.data_coding))
	body = append(body, byte(pdu.sm_default_msg_id))
	body = append(body, byte(len(pdu.short_message)))

	body = append(body, pdu.short_message...)

	packTLVs(&body, pdu)
	return body
}

func Bind(sequence_number uint32, command uint32, system_id string, password string, system_type string, interface_version int, addr_ton int, addr_npi int, address_range string) (pdu) {
	var bind pdu;
	bind.command_id = command
	bind.Command_status = 0
	bind.sequence_number = sequence_number

	bind.system_id = []byte(system_id)
	bind.password = []byte(password)
	bind.system_type = []byte(system_type)
	bind.interface_version = interface_version
	bind.addr_ton = addr_ton
	bind.addr_npi = addr_npi
	bind.address_range = []byte(address_range)
	return bind;
}


func SubmitSM(sequence_number uint32, system_type string, source_addr_ton int, source_addr_npi int, source_addr string,
dest_addr_ton int,dest_addr_npi int,dest_addr string, data_coding int, sm_default_msg_id int, short_message string) (pdu) {
	var submit pdu;
	submit.command_id = PDU_COMMAND_SUBMIT_SM
	submit.Command_status = 0
	submit.sequence_number = sequence_number

	submit.system_type = []byte(system_type)
	submit.source_addr_ton = source_addr_ton
	submit.source_addr_npi = source_addr_npi
	submit.source_addr = []byte(source_addr)

	submit.dest_addr_ton = dest_addr_ton
	submit.dest_addr_npi = dest_addr_npi
	submit.dest_addr = []byte(dest_addr)

	submit.esm_class = 0
	submit.protocol_id = 0
	submit.priority_flag = 0

	submit.schedule_delivery_time = []byte("")
	submit.validity_period = []byte("")

	submit.registered_delivery = 0
	submit.replace_if_present_flag = 0
	submit.data_coding = data_coding
	submit.sm_default_msg_id = sm_default_msg_id

	submit.sm_length = 0
	submit.short_message = []byte(short_message)

	return submit
}


