package helper

func StatusDiscountProposalByLevel(level string) (string, error) {
	var status string
	switch level {
	case "MR":
		status = " TEMPORARY,INPUT,CONFIRM,APPROVE,REJECT "
	case "SPV":
		status = "CONFIRM,APPROVE,REJECT "
	case "ASM":
		status = "CONFIRM LV2,APPROVE,REJECT"
	case "FSM":
		status = "CONFIRM LV3,APPROVE,REJECT"
	case "GM":
		status = "CONFIRM LV4,APPROVE,REJECT"
	case "MD":
		status = "CONFIRM LV5,APPROVE,REJECT"
	case "NON MKT":
		status = "TEMPORARY,INPUT,CONFIRM,CONFIRM LV2,CONFIRM LV3,CONFIRM LV4,CONFIRM LV5,CONFIRM LV6,APPROVE,REJECT"
	case "ADM HO":
		status = "TEMPORARY,INPUT,CONFIRM,CONFIRM LV2,CONFIRM LV3,CONFIRM LV4,CONFIRM LV5,CONFIRM LV6,APPROVE,REJECT"
	}
	return status, nil
}
