package offense

type Batter struct {
	PlayerID             string
	YearID               int
	Stint                int
	TeamID               string
	LeagueID             string
	Games                int
	AtBats               int
	Runs                 int
	Hits                 int
	Doubles              int
	Triples              int
	Homeruns             int
	RunsBattedIn         int
	StolenBases          int
	CaughtStealing       int
	Walks                int
	Strikeouts           int
	IntentionalWalks     int
	HitByPitch           int
	SacrificeBunt        int
	SacrificeFly         int
	GroundIntoDoublePlay int
	BattingAverage       float64
	OnBasePercentage     float64
	SluggingPercentage   float64
	OnBasePlusSlugging   float64
}
