package valorant

var maps map[string]string = map[string]string{
	"triad":      "haven",
	"duality":    "bind",
	"bonsai":     "split",
	"ascent":     "ascent",
	"port":       "icebox",
	"foxtrot":    "breeze",
	"canyon":     "fracture",
	"pitt":       "pearl",
	"jam":        "lotus",
	"range":      "practice",
	"hurm_alley": "district",
	"hurm_yard":  "piazza",
	"hurm_bowl":  "kasbah",
	"juliett":    "sunset",
}

func GetMapName(code string) string {
	return getNameFromCode(maps, code)
}
