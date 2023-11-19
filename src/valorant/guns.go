package valorant

var guns map[string]string = map[string]string{
	"tx_hud_pistol_classic":     "classic",
	"tx_hud_pistol_slim":        "shorty",
	"tx_hud_pistol_autopistol":  "frenzy",
	"tx_hud_pistol_luger":       "ghost",
	"tx_hud_pistol_sheriff":     "sheriff",
	"tx_hud_shotguns_pump":      "bucky",
	"tx_hud_shotguns_persuader": "judge",
	"tx_hud_smgs_vector":        "stinger",
	"tx_hud_smgs_ninja":         "spectre",
	"tx_hud_rifles_burst":       "bulldog",
	"tx_hud_rifles_dmr":         "guardian",
	"tx_hud_rifles_ghost":       "phantom",
	"tx_hud_rifles_volcano":     "vandal",
	"tx_hud_sniper_bolt":        "marshal",
	"tx_hud_sniper_operater":    "operator",
	"tx_hud_lmg":                "ares",
	"tx_hud_hmg":                "odin",
	"knife":                     "knife",
}

func GetGunName(code string) string {
	return getNameFromCode(guns, code)
}
