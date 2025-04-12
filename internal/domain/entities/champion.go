package entities

type Champion struct {
	ID           string  `json:"id"`
	Key          string  `json:"key"`
	Name         string  `json:"name"`
	Title        string  `json:"title"`
	MovementSpeed float64 `json:"movespeed"`
}

type ChampionData struct {
	Type    string                  `json:"type"`
	Format  string                  `json:"format"`
	Version string                  `json:"version"`
	Data    map[string]ChampionInfo `json:"data"`
}

type ChampionInfo struct {
	ID           string  `json:"id"`
	Key          string  `json:"key"`
	Name         string  `json:"name"`
	Title        string  `json:"title"`
	MovementSpeed float64 `json:"movespeed"`
}
