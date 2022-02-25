package main

const (
	tableMangle = "mangle"
	tableNat    = "nat"

	chainIP4        = "TP_CLASH_V4"
	chainIP6        = "TP_CLASH_V6"
	chainIP4Local   = "TP_CLASH_LOCAL_V4"
	chainIP6Local   = "TP_CLASH_LOCAL_V6"
	chainIP4DNS     = "TP_CLASH_DNS_V4"
	chainIP6DNS     = "TP_CLASH_DNS_V6"
	chainOutput     = "OUTPUT"
	chainPreRouting = "PREROUTING"

	actionReturn   = "RETURN"
	actionTProxy   = "TPROXY"
	actionRedirect = "REDIRECT"
	actionDNat     = "DNAT"
	actionMark     = "MARK"

	tproxyMark = "666"
	clashUser  = "tpclash"
)
