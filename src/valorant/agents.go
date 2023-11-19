package valorant

var agents map[string]string = map[string]string{
	"clay_pc_c":         "raze",
	"pandemic_pc_c":     "viper",
	"wraith_pc_c":       "omen",
	"hunter_pc_c":       "sova",
	"thorne_pc_c":       "sage",
	"phoenix_pc_c":      "phoenix",
	"wushu_pc_c":        "jett",
	"gumshoe_pc_c":      "cypher",
	"sarge_pc_c":        "brimstone",
	"breach_pc_c":       "breach",
	"vampire_pc_c":      "reyna",
	"killjoy_pc_c":      "killjoy",
	"guide_pc_c":        "skye",
	"stealth_pc_c":      "yoru",
	"rift_pc_c":         "astra",
	"grenadier_pc_c":    "kay/o",
	"deadeye_pc_c":      "chamber",
	"sprinter_pc_c":     "neon",
	"bountyhunter_pc_c": "fade",
	"mage_pc_c":         "harbor",
	"aggrobot_pc_c":     "gekko",
	"cable_pc_c":        "deadlock",
	"sequoia_pc_c":      "iso",
}

func GetAgentName(code string) string {
	return getNameFromCode(agents, code)
}
