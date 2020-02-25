package legiondata

type Data struct {
	Units        []Unit        `json:"units"`
	Keywords     []Keyword     `json:"keywords"`
	Upgrades     []Upgrade     `json:"upgrades"`
	CommandCards []CommandCard `json:"commandCards"`
	Sources      []Source      `json:"sources"`
}

type AttackDice struct {
	White int `json:"white,omitempty"`
	Black int `json:"black,omitempty"`
	Red   int `json:"red,omitempty"`
}

type Keyword struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      bool   `json:"action,omitempty"`
	Affects     string `json:"affects,omitempty"`
}

type Surge struct {
	Attack  string `json:"attack,omitempty"`
	Defense string `json:"defense,omitempty"`
}

type Unit struct {
	UID                string    `json:"uid,omitempty"`
	LDF                string    `json:"ldf"`
	Factions           []string  `json:"factions"`
	Unique             *bool     `json:"unique,omitempty"`
	Name               string    `json:"name"`
	Subtitle           string    `json:"subtitle,omitempty"`
	Type               string    `json:"type"`
	Points             int       `json:"points"`
	PrintedPoints      int       `json:"printed_points,omitempty"`
	PointsWithUpgrades int       `json:"pointsWithUpgrades,omitempty"`
	Rank               string    `json:"rank"`
	Minis              int       `json:"minis"`
	Wounds             int       `json:"wounds"`
	Courage            *int      `json:"courage,omitempty"`
	Resilience         *int      `json:"resilience,omitempty"`
	Defense            string    `json:"defense"`
	Surge              Surge     `json:"surge"`
	Speed              int       `json:"speed"`
	Slots              []string  `json:"slots"`
	Keywords           []string  `json:"keywords"`
	Weapons            []Weapon  `json:"weapons"`
	Upgrades           []Upgrade `json:"upgrades,omitempty"`
	CommandCards       []string  `json:"commandCards,omitempty"`
	Entourage          []string  `json:"entourage,omitempty"`
}

type Upgrade struct {
	LDF          string `json:"ldf"`
	Name         string `json:"name"`
	Side         string `json:"side,omitempty"`
	Unique       *bool  `json:"unique,omitempty"`
	Description  string `json:"description,omitempty"`
	Exhaust      *bool  `json:"exhaust,omitempty"`
	Restrictions *struct {
		LDF  string `json:"ldf,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"restrictions,omitempty"`
	UnitTypeExclusions []string `json:"unitTypeExclusions,omitempty"`
	Slot               string   `json:"slot"`
	Points             int      `json:"points"`
	PrintedPoints      int      `json:"printed_points,omitempty"`
	Keywords           []string `json:"keywords,omitempty"`
	Weapon             *Weapon  `json:"weapon,omitempty"`
}

type Weapon struct {
	Name  string `json:"name,omitempty"`
	Range struct {
		From int `json:"from"`
		To   int `json:"to,omitempty"`
	} `json:"range"`
	Keywords []string   `json:"keywords,omitempty"`
	Dice     AttackDice `json:"dice,omitempty"`
	Surge    *Surge     `json:"surge,omitempty"`
}

type CommandCard struct {
	LDF          string  `json:"ldf"`
	Name         string  `json:"name"`
	Pips         int     `json:"pips"`
	Orders       string  `json:"orders"`
	Description  string  `json:"description,omitempty"`
	Weapon       *Weapon `json:"weapon,omitempty"`
	Restrictions *struct {
		LDF  string `json:"ldf,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"restrictions,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

type Source struct {
	LDF      string `json:"ldf"`
	Name     string `json:"name"`
	Wave     int    `json:"wave"`
	Released bool   `json:"released"`
	Contents struct {
		Units    map[string]int `json:"units"`
		Upgrades map[string]int `json:"Upgrades"`
	} `json:"contents"`
}
