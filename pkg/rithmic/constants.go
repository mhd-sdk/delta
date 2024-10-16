package rithmic

import "delta/pkg/generated/rti"

const DEFAULT_RITHMIC_SYSTEM_NAME = "Rithmic Test"

var AVAILABLE_RITHMIC_INFRA_TYPES = []rti.RequestLogin_SysInfraType{
	rti.RequestLogin_TICKER_PLANT,
	rti.RequestLogin_ORDER_PLANT,
	rti.RequestLogin_HISTORY_PLANT,
	rti.RequestLogin_PNL_PLANT,
	rti.RequestLogin_REPOSITORY_PLANT,
}
