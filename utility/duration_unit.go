package utility

var durationUnit = map[int8]string{
	1: "month",
	2: "year",
}

var contractStatus = map[int8]string{
	1: "active",
	2: "expired",
	3: "cancelled",
}

func GetDurationUnitText(unitId int8) string {
	return durationUnit[unitId]
}

func GetContractStatus(statusId int8) string {
	return contractStatus[statusId]
}
