package formulas

func errorCode(code int) string {
	switch code {
	case 1:
		return "#NULL!"
	case 2:
		return "#DIV/0!"
	case 3:
		return "#VALUE?"
	case 4:
		return "#REF!"
	case 5:
		return "#NAME?"
	case 6:
		return "#NUM!"
	case 7:
		return "#ERROR!"
	case 8:
		return "GETTING_DATA"
	default:
		return "#N/A"
	}
}
