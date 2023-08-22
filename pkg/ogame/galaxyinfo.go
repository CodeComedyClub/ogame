package ogame

import (
	"encoding/json"
)

type GalaxyPageContent struct {
	FilterSettings struct {
		FilterEmpty    bool `json:"filter_empty"`
		FilterInactive bool `json:"filter_inactive"`
		FilterNewbie   bool `json:"filter_newbie"`
		FilterStrong   bool `json:"filter_strong"`
		FilterVacation bool `json:"filter_vacation"`
	} `json:"filterSettings"`
	LifeformEnabled bool   `json:"lifeformEnabled"`
	NewAjaxToken    string `json:"newAjaxToken"`
	Success         bool   `json:"success"`
	System          struct {
		AvailableMissiles    int64 `json:"availableMissiles"`
		AvailablePathfinders int64 `json:"availablePathfinders"`
		AvailableProbes      int64 `json:"availableProbes"`
		AvailableRecyclers   int64 `json:"availableRecyclers"`
		CanColonize          bool  `json:"canColonize"`
		CanExpedition        bool  `json:"canExpedition"`
		CanFly               bool  `json:"canFly"`
		CanSwitchGalaxy      bool  `json:"canSwitchGalaxy"`
		CanSystemEspionage   bool  `json:"canSystemEspionage"`
		CanSystemPhalanx     bool  `json:"canSystemPhalanx"`
		CurrentPlanetID      int64 `json:"currentPlanetId"`
		DeuteriumInDebris    bool  `json:"deuteriumInDebris"`
		Galaxy               int64 `json:"galaxy"`
		GalaxyContent        []struct {
			AvailableMissions []struct {
				CanSend              any    `json:"canSend"`
				DarkMatterCost       int64  `json:"darkMatterCost"`
				Description          string `json:"description"`
				DiscoveryCount       string `json:"discoveryCount"`
				GalaxyLink           string `json:"galaxyLink"`
				Link                 string `json:"link"`
				MissionType          int64  `json:"missionType"`
				MoveAction           string `json:"moveAction"`
				MoveLink             string `json:"moveLink"`
				PlanetMovePossible   bool   `json:"planetMovePossible"`
				SufficientDarkMatter bool   `json:"sufficientDarkMatter"`
				Title                string `json:"title"`
			} `json:"availableMissions"`
			Galaxy  int64                  `json:"galaxy"`
			Planets PlanetInfoArrayWrapper `json:"planets"`
			Player  struct {
				Actions struct {
					Alliance struct {
						AllianceClassCSS  string `json:"allianceClassCss"`
						AllianceClassName string `json:"allianceClassName"`
						Available         bool   `json:"available"`
						HighscoreLink     string `json:"highscoreLink"`
						HighscoreTitle    string `json:"highscoreTitle"`
						InfoPageLink      string `json:"infoPageLink"`
						InfoPageTitle     string `json:"infoPageTitle"`
						MemberCount       int64  `json:"memberCount"`
					} `json:"alliance"`
					Buddies struct {
						Available   bool   `json:"available"`
						Link        string `json:"link"`
						SupportIcon bool   `json:"supportIcon"`
						Title       string `json:"title"`
					} `json:"buddies"`
					Highscore struct {
						Available bool   `json:"available"`
						Link      string `json:"link"`
						Title     string `json:"title"`
					} `json:"highscore"`
					Ignore struct {
						Available bool   `json:"available"`
						Link      string `json:"link"`
						Title     string `json:"title"`
					} `json:"ignore"`
					Message struct {
						Available       bool   `json:"available"`
						DisabledChatBar bool   `json:"disabledChatBar"`
						Link            string `json:"link"`
						Title           string `json:"title"`
					} `json:"message"`
				} `json:"actions"`
				AllianceID                int64  `json:"allianceId"`
				AllianceName              string `json:"allianceName"`
				AllianceTag               string `json:"allianceTag"`
				HighscorePositionAlliance string `json:"highscorePositionAlliance"`
				HighscorePositionPlayer   int64  `json:"highscorePositionPlayer"`
				IsActive                  bool   `json:"isActive"`
				IsAdmin                   bool   `json:"isAdmin"`
				IsAllianceMember          bool   `json:"isAllianceMember"`
				IsBanned                  bool   `json:"isBanned"`
				IsBuddy                   bool   `json:"isBuddy"`
				IsHonorableTarget         bool   `json:"isHonorableTarget"`
				IsInactive                bool   `json:"isInactive"`
				IsLongInactive            bool   `json:"isLongInactive"`
				IsNewbie                  bool   `json:"isNewbie"`
				IsOnVacation              bool   `json:"isOnVacation"`
				IsOutlaw                  bool   `json:"isOutlaw"`
				IsStrong                  bool   `json:"isStrong"`
				Phalanx                   struct {
					Inactive bool   `json:"inactive"`
					Link     string `json:"link"`
					Title    string `json:"title"`
				} `json:"phalanx"`
				PlayerID   int64  `json:"playerId"`
				PlayerName string `json:"playerName"`
				Rank       struct {
					HasRank   bool   `json:"hasRank"`
					RankClass string `json:"rankClass"`
					RankTitle string `json:"rankTitle"`
				} `json:"rank"`
			} `json:"player"`
			Position        int64  `json:"position"`
			PositionFilters string `json:"positionFilters"`
			System          int64  `json:"system"`
		} `json:"galaxyContent"`
		HasAdmiral                 bool   `json:"hasAdmiral"`
		HasBirthdayPlanet          bool   `json:"hasBirthdayPlanet"`
		IsOutlaw                   bool   `json:"isOutlaw"`
		MaximumFleetSlots          int64  `json:"maximumFleetSlots"`
		PlayerID                   int64  `json:"playerId"`
		SettingsProbeCount         int64  `json:"settingsProbeCount"`
		ShowOutlawWarning          bool   `json:"showOutlawWarning"`
		SlotsColonized             int64  `json:"slotsColonized"`
		SwitchGalaxyDeuteriumCosts int64  `json:"switchGalaxyDeuteriumCosts"`
		System                     int64  `json:"system"`
		ToGalaxyLink               string `json:"toGalaxyLink"`
		UsedFleetSlots             int64  `json:"usedFleetSlots"`
	} `json:"system"`
	Token string `json:"token"`
}

type ExpeditionInfo struct {
	PlanetID          int64  `json:"planetId"`
	PlanetName        string `json:"planetName"`
	ImageInformation  string `json:"imageInformation"`
	PlanetType        int64  `json:"planetType"`
	RequiredShips     int64  `json:"requiredShips"`
	AvailableMissions []struct {
		MissionType int64  `json:"missionType"`
		Name        string `json:"name"`
	} `json:"availableMissions"`
	Resources struct {
		Metal struct {
			Name   string `json:"name"`
			Amount int64  `json:"amount"`
		} `json:"metal"`
		Crystal struct {
			Name   string `json:"name"`
			Amount int64  `json:"amount"`
		} `json:"crystal"`
		Deuterium struct {
			Name   string `json:"name"`
			Amount int64  `json:"amount"`
		} `json:"deuterium"`
	} `json:"resources"`
	RecyclePossible bool `json:"recyclePossible"`
}

type PlanetInfo struct {
	Activity struct {
		IdleTime int64 `json:"idleTime"`
	} `json:"activity"`
	AvailableMissions []struct {
		CanSpy      bool   `json:"canSpy"`
		Link        string `json:"link"`
		MissionType int64  `json:"missionType"`
		Name        string `json:"name"`
		ReportID    string `json:"reportId"`
		ReportLink  string `json:"reportLink"`
	} `json:"availableMissions"`
	ImageInformation string `json:"imageInformation"`
	IsDestroyed      bool   `json:"isDestroyed"`
	PlanetID         int64  `json:"planetId"`
	PlanetName       string `json:"planetName"`
	PlanetType       int64  `json:"planetType"`
	PlayerID         int64  `json:"playerId"`
	RecyclePossible  bool   `json:"recyclePossible"`
	RequiredShips    int64  `json:"requiredShips"`
	Resources        struct {
		Crystal struct {
			Amount string `json:"amount"`
			Name   string `json:"name"`
		} `json:"crystal"`
		Deuterium struct {
			Amount string `json:"amount"`
			Name   string `json:"name"`
		} `json:"deuterium"`
		Metal struct {
			Amount string `json:"amount"`
			Name   string `json:"name"`
		} `json:"metal"`
	} `json:"resources"`
	Size string `json:"size"`
}

type PlanetInfoArrayWrapper struct {
	Planets    []PlanetInfo
	Expedition ExpeditionInfo
}

/**
* Wrapper to handle multitype json attribute planets
 */
func (w *PlanetInfoArrayWrapper) UnmarshalJSON(data []byte) error {
	var res ExpeditionInfo
	if err := json.Unmarshal(data, &res); err == nil {
		w.Expedition = res
		return nil
	}
	return json.Unmarshal(data, &w.Planets)
}
