package omcmd

type (
	// OptsGlobal contains options accepted by all actions
	OptsGlobal struct {
		Color          string
		Output         string
		Local          bool
		ObjectSelector string
		Quiet          bool
		Debug          bool
	}
)
