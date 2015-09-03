package go_smpp
import (
	"fmt"
)

var pduCommandLabel = map[int]string{
	PDU_COMMAND_BIND_RX: "BIND_RX",
	PDU_COMMAND_BIND_TX: "BIND_TX",
	PDU_COMMAND_QUERY_SM: "QUERY_SM",
	PDU_COMMAND_SUBMIT_SM: "SUBMIT_SM",
	PDU_COMMAND_DELIVER_SM: "DELIVER_SM",
	PDU_COMMAND_UNBIND: "UNBIND",
	PDU_COMMAND_REPLACE_SM: "REPLACE_SM",
	PDU_COMMAND_CANCEL_SM: "CANCEL_SM",
	PDU_COMMAND_BIND_TRX: "BIND_TRX",
	PDU_COMMAND_ENQUIRE: "ENQUIRE",
	PDU_COMMAND_SUBMIT_MULTI: "SUBMIT_MULTI",
	PDU_COMMAND_ALERT: "ALERT",
	PDU_COMMAND_DATA_SM: "DATA_SM",
	PDU_COMMAND_RESP: "RESP",
}

var pduCommandStatusLabel = map[int]string{
	PDU_COMMAND_STATUS_ESME_ROK: "No Error",
	PDU_COMMAND_STATUS_ESME_RINVMSGLEN: "Message Length is invalid",
	PDU_COMMAND_STATUS_ESME_RINVCMDLEN: "Command Length is invalid",
	PDU_COMMAND_STATUS_ESME_RINVCMDID: "Invalid Command ID",
	PDU_COMMAND_STATUS_ESME_RINVBNDSTS: "Incorrect BIND Status for given command",
	PDU_COMMAND_STATUS_ESME_RALYBND: "ESME Already in Bound State",
	PDU_COMMAND_STATUS_ESME_RINVPRTFLG: "Invalid Priority Flag",
	PDU_COMMAND_STATUS_ESME_RINVREGDLVFLG: "Invalid Registered Delivery Flag",
	PDU_COMMAND_STATUS_ESME_RSYSERR: "System Error",
	PDU_COMMAND_STATUS_ESME_RINVSRCADR: "Invalid Source Address",
	PDU_COMMAND_STATUS_ESME_RINVDSTADR: "Invalid Dest Addr",
	PDU_COMMAND_STATUS_ESME_RINVMSGID: "Message ID is invalid",
	PDU_COMMAND_STATUS_ESME_RBINDFAIL: "Bind Failed",
	PDU_COMMAND_STATUS_ESME_RINVPASWD: "Invalid Password",
	PDU_COMMAND_STATUS_ESME_RINVSYSID: "Invalid System ID",
	PDU_COMMAND_STATUS_ESME_RCANCELFAIL: "Cancel SM Failed",
	PDU_COMMAND_STATUS_ESME_RREPLACEFAIL: "Replace SM Failed",
	PDU_COMMAND_STATUS_ESME_RMSGQFUL: "Message Queue Full",
	PDU_COMMAND_STATUS_ESME_RINVSERTYP: "Invalid Service Type",
	PDU_COMMAND_STATUS_ESME_RINVNUMDESTS: "Invalid number of destinations",
	PDU_COMMAND_STATUS_ESME_RINVDLNAME: "Invalid Distribution List name",
	PDU_COMMAND_STATUS_ESME_RINVDESTFLAG: "Destination flag is invalid (submit_multi)",
	PDU_COMMAND_STATUS_ESME_RINVSUBREP: "Invalid ‘submit with replace’ request(i.e. submit_sm with replace_if_present_flag set)",
	PDU_COMMAND_STATUS_ESME_RINVESMCLASS: "Invalid esm_class field data",
	PDU_COMMAND_STATUS_ESME_RCNTSUBDL: "Cannot Submit to Distribution List",
	PDU_COMMAND_STATUS_ESME_RSUBMITFAIL: "submit_sm or submit_multi failed",
	PDU_COMMAND_STATUS_ESME_RINVSRCTON: "Invalid Source address TON",
	PDU_COMMAND_STATUS_ESME_RINVSRCNPI: "Invalid Source address NPI",
	PDU_COMMAND_STATUS_ESME_RINVDSTTON: "Invalid Destination address TON",
	PDU_COMMAND_STATUS_ESME_RINVDSTNPI: "Invalid Destination address NPI",
	PDU_COMMAND_STATUS_ESME_RINVSYSTYP: "Invalid system_type field",
	PDU_COMMAND_STATUS_ESME_RINVREPFLAG: "Invalid replace_if_present flag",
	PDU_COMMAND_STATUS_ESME_RINVNUMMSGS: "Invalid number of messages",
	PDU_COMMAND_STATUS_ESME_RTHROTTLED: "Throttling error (ESME has exceeded allowed message limits)",
}

var pduOptionalParameterTag  = map[int]string{
	PDU_OPT_PARAM_DEST_ADDR_SUBUNIT: "dest_addr_subunit",
	PDU_OPT_PARAM_DEST_NETWORK_TYPE: "dest_network_type",
	PDU_OPT_PARAM_DEST_BEARER_TYPE: "dest_bearer_type",
	PDU_OPT_PARAM_DEST_TELEMATICS_ID: "dest_telematics_id",
	PDU_OPT_PARAM_SOURCE_ADDR_SUBUNIT: "source_addr_subunit",
	PDU_OPT_PARAM_SOURCE_NETWORK_TYPE: "source_network_type",
	PDU_OPT_PARAM_SOURCE_BEARER_TYPE: "source_bearer_type",
	PDU_OPT_PARAM_SOURCE_TELEMATICS_ID: "source_telematics_id",
	PDU_OPT_PARAM_QOS_TIME_TO_LIVE: "qos_time_to_live",
	PDU_OPT_PARAM_PAYLOAD_TYPE: "payload_type",
	PDU_OPT_PARAM_ADDITIONAL_STATUS_INFO_TEXT: "additional_status_info_text",
	PDU_OPT_PARAM_RECEIPTED_MESSAGE_ID: "receipted_message_id",
	PDU_OPT_PARAM_MS_MSG_WAIT_FACILITIES: "ms_msg_wait_facilities",
	PDU_OPT_PARAM_PRIVACY_INDICATOR: "privacy_indicator",
	PDU_OPT_PARAM_SOURCE_SUBADDRESS: "source_subaddress",
	PDU_OPT_PARAM_DEST_SUBADDRESS: "dest_subaddress",
	PDU_OPT_PARAM_USER_MESSAGE_REFERENCE: "user_message_reference",
	PDU_OPT_PARAM_USER_RESPONSE_CODE: "user_response_code",
	PDU_OPT_PARAM_SOURCE_PORT: "source_port",
	PDU_OPT_PARAM_DESTINATION_PORT: "destination_port",
	PDU_OPT_PARAM_SAR_MSG_REF_NUM: "sar_msg_ref_num",
	PDU_OPT_PARAM_LANGUAGE_INDICATOR: "language_indicator",
	PDU_OPT_PARAM_SAR_TOTAL_SEGMENTS: "sar_total_segments",
	PDU_OPT_PARAM_SAR_SEGMENT_SEQNUM: "sar_segment_seqnum",
	PDU_OPT_PARAM_SC_INTERFACE_VERSION: "sc_interface_version",
	PDU_OPT_PARAM_CALLBACK_NUM_PRES_IND: "callback_num_pres_ind",
	PDU_OPT_PARAM_CALLBACK_NUM_ATAG: "callback_num_atag",
	PDU_OPT_PARAM_NUMBER_OF_MESSAGES: "number_of_messages",
	PDU_OPT_PARAM_CALLBACK_NUM: "callback_num",
	PDU_OPT_PARAM_DPF_RESULT: "dpf_result",
	PDU_OPT_PARAM_SET_DPF: "set_dpf",
	PDU_OPT_PARAM_MS_AVAILABILITY_STATUS: "ms_availability_status",
	PDU_OPT_PARAM_NETWORK_ERROR_CODE: "network_error_code",
	PDU_OPT_PARAM_MESSAGE_PAYLOAD: "message_payload",
	PDU_OPT_PARAM_DELIVERY_FAILURE_REASON: "delivery_failure_reason",
	PDU_OPT_PARAM_MORE_MESSAGES_TO_SEND: "more_messages_to_send",
	PDU_OPT_PARAM_MESSAGE_STATE: "message_state",
	PDU_OPT_PARAM_USSD_SERVICE_OP: "ussd_service_op",
	PDU_OPT_PARAM_DISPLAY_TIME: "display_time",
	PDU_OPT_PARAM_SMS_SIGNAL: "sms_signal",
	PDU_OPT_PARAM_MS_VALIDITY: "ms_validity",
	PDU_OPT_PARAM_ALERT_ON_MESSAGE_DELIVERY: "alert_on_message_delivery",
	PDU_OPT_PARAM_ITS_REPLY_TYPE: "its_reply_type",
	PDU_OPT_PARAM_ITS_SESSION_INFO: "its_session_info",
}

var pduTonLabel  = map[int]string{
	PDU_TON_UNKNOWN: "Unknown",
	PDU_TON_INTERNATIONAL: "International",
	PDU_TON_NATIONAL: "National",
	PDU_TON_NETWORK_SPECIFIC: "Network specific",
	PDU_TON_SUBSCRIBER_NUMBER: "Subscriber number",
	PDU_TON_ALPHANUMERIC: "Alphanumeric",
	PDU_TON_ABBREVIATED: "Abbreviated",
}

var pduNpiLabel = map[int]string{
	PDU_NPI_UNKNOWN: "Unknown",
	PDU_NPI_ISDN: "Isdn",
	PDU_NPI_DATA: "Data",
	PDU_NPI_TELEX: "Telex",
	PDU_NPI_LAND_MOBILE: "Land_mobile",
	PDU_NPI_NATIONAL: "National",
	PDU_NPI_PRIVATE: "Private",
	PDU_NPI_ERMES: "Ermes",
	PDU_NPI_INTERNET: "Internet",
	PDU_NPI_WAP: "Wap",
}

var pduDCSLabel = map[int]string{
	PDU_DATA_CODING_DEFAULT  : "SMSC Default Alphabet",
	PDU_DATA_CODING_IA5      : "IA5 (CCITT T.50)/ASCII (ANSI X3.4) b",
	PDU_DATA_CODING_OCTET_B  : "Octet unspecified (8-bit binary) b",
	PDU_DATA_CODING_LATIN_1  : "Latin 1 (ISO-8859-1) b",
	PDU_DATA_CODING_OCTET_A  : "Octet unspecified (8-bit binary) a",
	PDU_DATA_CODING_JIS      : "JIS (X 0208-1990) b",
	PDU_DATA_CODING_CYRLLIC  : "Cyrllic (ISO-8859-5) b",
	PDU_DATA_CODING_LATIN    : "Latin/Hebrew (ISO-8859-8) b",
	PDU_DATA_CODING_UCS2     : "UCS2 (ISO/IEC-10646) a",
	PDU_DATA_CODING_PICT     : "Pictogram Encoding b",
	PDU_DATA_CODING_MUSIC    : "ISO-2022-JP (Music Codes) b",
	PDU_DATA_CODING_KANJI    : "Extended Kanji JIS(X 0212-1990) b",
	PDU_DATA_CODING_KSC      : "KS C 5601 b",
}

func (p *Pdu) PrintDetailed () (string) {
	var display string

	display = fmt.Sprintf("Command Length: %d\n", p.command_length)
	display = fmt.Sprintf("%sCommand: %s\n", display, pduCommandLabel[int(p.command_id)])
	display = fmt.Sprintf("%sCommand Status: %s\n", display, pduCommandStatusLabel[int(p.Command_status)])
	display = fmt.Sprintf("%sSequence: %d\n", display, p.sequence_number)

	display = fmt.Sprintf("Header:\n%s\nBody:\n%s\n", display, p.printBody())

	return display
}

func (p *Pdu) printBody () (string) {
	switch p.command_id {
	case PDU_COMMAND_BIND_TX,PDU_COMMAND_BIND_RX,PDU_COMMAND_BIND_TRX:
		return p.printBindBody()
	case PDU_COMMAND_SUBMIT_SM, PDU_COMMAND_DATA_SM:
		return p.printSubmitBody()
	}
	return ""
}

func (p *Pdu) printBindBody() (string) {
	var display string
	display = fmt.Sprintf("System Id: %s\n", string(p.system_id))
	display = fmt.Sprintf("%sPassword: %s\n", display, string(p.password))
	display = fmt.Sprintf("%sSystem Type: %s\n", display, string(p.system_type))

	display = fmt.Sprintf("%sInterface Version: %d\n", display, p.interface_version)
	display = fmt.Sprintf("%sAddr TON: %d (%s)\n", display, p.addr_ton, pduTonLabel[int(p.addr_ton)])
	display = fmt.Sprintf("%sAddr TON: %d (%s)\n", display, p.addr_npi, pduNpiLabel[int(p.addr_npi)])

	display = fmt.Sprintf("%sAddress Range: %s\n", display, string(p.address_range))
	return display
}

func (p *Pdu) printSubmitBody() (string) {
	var display string
	display = fmt.Sprintf("System Type: %s\n", string(p.system_type))

	display = fmt.Sprintf("%sSource Addr TON: %d (%s)\n", display, p.source_addr_ton, pduTonLabel[int(p.source_addr_ton)])
	display = fmt.Sprintf("%sSource Addr NPI: %d (%s)\n", display, p.source_addr_npi, pduNpiLabel[int(p.source_addr_npi)])
	display = fmt.Sprintf("%sSource Addr: %s\n", display, string(p.source_addr))

	display = fmt.Sprintf("%sDestination Addr TON: %d (%s)\n", display, p.dest_addr_ton, pduTonLabel[int(p.dest_addr_ton)])
	display = fmt.Sprintf("%sDestination Addr NPI: %d (%s)\n", display, p.dest_addr_npi, pduNpiLabel[int(p.dest_addr_npi)])
	display = fmt.Sprintf("%sDestination Addr: %s\n", display, string(p.dest_addr))

	display = fmt.Sprintf("%sESM Class: %d (%s)\n", display, p.esm_class, "")
	display = fmt.Sprintf("%sProtocol ID: %d (%s)\n", display, p.protocol_id, "")
	display = fmt.Sprintf("%sPriority Flag: %d (%s)\n", display, p.priority_flag, "")

	display = fmt.Sprintf("%sScheduled: %d\n", display, p.schedule_delivery_time)
	display = fmt.Sprintf("%sValid Period: %d\n", display, p.validity_period)

	display = fmt.Sprintf("%sregistered_delivery: %d (%s)\n", display, p.registered_delivery, "")
	display = fmt.Sprintf("%sreplace_if_present_flag: %d (%s)\n", display, p.replace_if_present_flag, "")
	display = fmt.Sprintf("%sdata_coding: %d (%s)\n", display, p.data_coding, pduDCSLabel[p.data_coding])
	display = fmt.Sprintf("%ssm_default_msg_id: %d (%s)\n", display, p.sm_default_msg_id, "")

	display = fmt.Sprintf("%sMsg:\n%s", display, string(p.short_message))

	return display
}


func (p *Pdu) PrintOneLine () (string) {
	var display string

	display = fmt.Sprintf("%s (%d) \"%s\" #%d", pduCommandLabel[int(p.command_id)], p.command_length, pduCommandStatusLabel[int(p.Command_status)], p.sequence_number)

	display = fmt.Sprintf("%s: %s", display, p.printOneLineBody())

	return display
}

func (p *Pdu) printOneLineBody () (string) {
	switch p.command_id {
	case PDU_COMMAND_BIND_TX,PDU_COMMAND_BIND_RX,PDU_COMMAND_BIND_TRX:
		return p.printOneLineBindBody()
	case PDU_COMMAND_SUBMIT_SM, PDU_COMMAND_DATA_SM:
		return p.printOneLineSubmitBody()
	}

	return ""
}

func (p *Pdu) printOneLineBindBody() (string) {
	var display string
	display = fmt.Sprintf("Id: %s", string(p.system_id))
	display = fmt.Sprintf("%s System Type: %s", display, string(p.system_type))

	display = fmt.Sprintf("%s v:%d", display, p.interface_version)
	display = fmt.Sprintf("%s (TON/NPI) %d/%d", display, p.addr_ton, p.addr_npi)

	display = fmt.Sprintf("%s [%s]", display, string(p.address_range))
	return display
}

func (p *Pdu) printOneLineSubmitBody() (string) {
	var display string
	display = fmt.Sprintf("System Type: %s", string(p.system_type))
	display = fmt.Sprintf("%s src:%s (%d/%d)", display, string(p.source_addr),p.source_addr_ton, p.source_addr_npi)
	display = fmt.Sprintf("%s dst:%s (%d/%d)", display, string(p.dest_addr),p.dest_addr_ton, p.dest_addr_npi)

	display = fmt.Sprintf("%s esm/proto/pri:%d/%d/%d", display, p.esm_class, p.protocol_id, p.priority_flag)

	display = fmt.Sprintf("%s sch/valid:%s/%s", display, string(p.schedule_delivery_time), string(p.validity_period))

	display = fmt.Sprintf("%s reg/rep/dcs:%d/%d/%d", display, p.registered_delivery, p.replace_if_present_flag, p.data_coding)
	display = fmt.Sprintf("%s %s", display, p.short_message)

	return display
}


