package v9

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/alaingilbert/ogame/pkg/extractor/v6"
	v7 "github.com/alaingilbert/ogame/pkg/extractor/v7"
	"github.com/alaingilbert/ogame/pkg/extractor/v71"
	"github.com/alaingilbert/ogame/pkg/ogame"
	"github.com/alaingilbert/ogame/pkg/utils"
	"regexp"
	"strings"
	"time"
)

func extractCancelLfBuildingInfos(pageHTML []byte) (token string, id, listID int64, err error) {
	return v7.ExtractCancelInfos(pageHTML, "cancelLinklfbuilding", "cancellfbuilding", 1)
}

func extractCancelResearchInfos(pageHTML []byte, lifeformEnabled bool) (token string, techID, listID int64, err error) {
	tableIdx := 1
	if lifeformEnabled {
		tableIdx = 2
	}
	return v7.ExtractCancelInfos(pageHTML, "cancelLinkresearch", "cancelresearch", tableIdx)
}

func extractEmpire(pageHTML []byte) ([]ogame.EmpireCelestial, error) {
	var out []ogame.EmpireCelestial
	raw, err := v6.ExtractEmpireJSON(pageHTML)
	if err != nil {
		return nil, err
	}
	j, ok := raw.(map[string]any)
	if !ok {
		return nil, errors.New("failed to parse json")
	}
	planetsRaw, ok := j["planets"].([]any)
	if !ok {
		return nil, errors.New("failed to parse json")
	}
	for _, planetRaw := range planetsRaw {
		planet, ok := planetRaw.(map[string]any)
		if !ok {
			return nil, errors.New("failed to parse json")
		}

		var tempMin, tempMax int64
		temperatureStr := utils.DoCastStr(planet["temperature"])
		m := v6.TemperatureRgx.FindStringSubmatch(temperatureStr)
		if len(m) == 3 {
			tempMin = utils.DoParseI64(m[1])
			tempMax = utils.DoParseI64(m[2])
		}
		mm := v6.DiameterRgx.FindStringSubmatch(utils.DoCastStr(planet["diameter"]))
		energyStr := utils.DoCastStr(planet["energy"])
		energyDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(energyStr))
		energy := utils.ParseInt(energyDoc.Find("div span").Text())
		celestialType := ogame.CelestialType(utils.DoCastF64(planet["type"]))
		out = append(out, ogame.EmpireCelestial{
			Name:     utils.DoCastStr(planet["name"]),
			ID:       ogame.CelestialID(utils.DoCastF64(planet["id"])),
			Diameter: utils.ParseInt(mm[1]),
			Img:      utils.DoCastStr(planet["image"]),
			Type:     celestialType,
			Fields: ogame.Fields{
				Built: utils.DoParseI64(utils.DoCastStr(planet["fieldUsed"])),
				Total: utils.DoParseI64(utils.DoCastStr(planet["fieldMax"])),
			},
			Temperature: ogame.Temperature{
				Min: tempMin,
				Max: tempMax,
			},
			Coordinate: ogame.Coordinate{
				Galaxy:   int64(utils.DoCastF64(planet["galaxy"])),
				System:   int64(utils.DoCastF64(planet["system"])),
				Position: int64(utils.DoCastF64(planet["position"])),
				Type:     celestialType,
			},
			Resources: ogame.Resources{
				Metal:     int64(utils.DoCastF64(planet["metal"])),
				Crystal:   int64(utils.DoCastF64(planet["crystal"])),
				Deuterium: int64(utils.DoCastF64(planet["deuterium"])),
				Energy:    energy,
			},
			Supplies: ogame.ResourcesBuildings{
				MetalMine:            int64(utils.DoCastF64(planet["1"])),
				CrystalMine:          int64(utils.DoCastF64(planet["2"])),
				DeuteriumSynthesizer: int64(utils.DoCastF64(planet["3"])),
				SolarPlant:           int64(utils.DoCastF64(planet["4"])),
				FusionReactor:        int64(utils.DoCastF64(planet["12"])),
				SolarSatellite:       int64(utils.DoCastF64(planet["212"])),
				MetalStorage:         int64(utils.DoCastF64(planet["22"])),
				CrystalStorage:       int64(utils.DoCastF64(planet["23"])),
				DeuteriumTank:        int64(utils.DoCastF64(planet["24"])),
			},
			Facilities: ogame.Facilities{
				RoboticsFactory: int64(utils.DoCastF64(planet["14"])),
				Shipyard:        int64(utils.DoCastF64(planet["21"])),
				ResearchLab:     int64(utils.DoCastF64(planet["31"])),
				AllianceDepot:   int64(utils.DoCastF64(planet["34"])),
				MissileSilo:     int64(utils.DoCastF64(planet["44"])),
				NaniteFactory:   int64(utils.DoCastF64(planet["15"])),
				Terraformer:     int64(utils.DoCastF64(planet["33"])),
				SpaceDock:       int64(utils.DoCastF64(planet["36"])),
				LunarBase:       int64(utils.DoCastF64(planet["41"])),
				SensorPhalanx:   int64(utils.DoCastF64(planet["42"])),
				JumpGate:        int64(utils.DoCastF64(planet["43"])),
			},
			Defenses: ogame.DefensesInfos{
				RocketLauncher:         int64(utils.DoCastF64(planet["401"])),
				LightLaser:             int64(utils.DoCastF64(planet["402"])),
				HeavyLaser:             int64(utils.DoCastF64(planet["403"])),
				GaussCannon:            int64(utils.DoCastF64(planet["404"])),
				IonCannon:              int64(utils.DoCastF64(planet["405"])),
				PlasmaTurret:           int64(utils.DoCastF64(planet["406"])),
				SmallShieldDome:        int64(utils.DoCastF64(planet["407"])),
				LargeShieldDome:        int64(utils.DoCastF64(planet["408"])),
				AntiBallisticMissiles:  int64(utils.DoCastF64(planet["502"])),
				InterplanetaryMissiles: int64(utils.DoCastF64(planet["503"])),
			},
			Researches: ogame.Researches{
				EnergyTechnology:             int64(utils.DoCastF64(planet["113"])),
				LaserTechnology:              int64(utils.DoCastF64(planet["120"])),
				IonTechnology:                int64(utils.DoCastF64(planet["121"])),
				HyperspaceTechnology:         int64(utils.DoCastF64(planet["114"])),
				PlasmaTechnology:             int64(utils.DoCastF64(planet["122"])),
				CombustionDrive:              int64(utils.DoCastF64(planet["115"])),
				ImpulseDrive:                 int64(utils.DoCastF64(planet["117"])),
				HyperspaceDrive:              int64(utils.DoCastF64(planet["118"])),
				EspionageTechnology:          int64(utils.DoCastF64(planet["106"])),
				ComputerTechnology:           int64(utils.DoCastF64(planet["108"])),
				Astrophysics:                 int64(utils.DoCastF64(planet["124"])),
				IntergalacticResearchNetwork: int64(utils.DoCastF64(planet["123"])),
				GravitonTechnology:           int64(utils.DoCastF64(planet["199"])),
				WeaponsTechnology:            int64(utils.DoCastF64(planet["109"])),
				ShieldingTechnology:          int64(utils.DoCastF64(planet["110"])),
				ArmourTechnology:             int64(utils.DoCastF64(planet["111"])),
			},
			Ships: ogame.ShipsInfos{
				LightFighter:   int64(utils.DoCastF64(planet["204"])),
				HeavyFighter:   int64(utils.DoCastF64(planet["205"])),
				Cruiser:        int64(utils.DoCastF64(planet["206"])),
				Battleship:     int64(utils.DoCastF64(planet["207"])),
				Battlecruiser:  int64(utils.DoCastF64(planet["215"])),
				Bomber:         int64(utils.DoCastF64(planet["211"])),
				Destroyer:      int64(utils.DoCastF64(planet["213"])),
				Deathstar:      int64(utils.DoCastF64(planet["214"])),
				SmallCargo:     int64(utils.DoCastF64(planet["202"])),
				LargeCargo:     int64(utils.DoCastF64(planet["203"])),
				ColonyShip:     int64(utils.DoCastF64(planet["208"])),
				Recycler:       int64(utils.DoCastF64(planet["209"])),
				EspionageProbe: int64(utils.DoCastF64(planet["210"])),
				SolarSatellite: int64(utils.DoCastF64(planet["212"])),
				Crawler:        int64(utils.DoCastF64(planet["217"])),
				Reaper:         int64(utils.DoCastF64(planet["218"])),
				Pathfinder:     int64(utils.DoCastF64(planet["219"])),
			},
		})
	}
	return out, nil
}

func extractOverviewProductionFromDoc(doc *goquery.Document) ([]ogame.Quantifiable, error) {
	res := make([]ogame.Quantifiable, 0)
	active := doc.Find("table.construction").Eq(4)
	href, _ := active.Find("td a").Attr("href")
	m := regexp.MustCompile(`openTech=(\d+)`).FindStringSubmatch(href)
	if len(m) == 0 {
		return []ogame.Quantifiable{}, nil
	}
	idInt := utils.DoParseI64(m[1])
	activeID := ogame.ID(idInt)
	activeNbr := utils.DoParseI64(active.Find("div.shipSumCount").Text())
	res = append(res, ogame.Quantifiable{ID: activeID, Nbr: activeNbr})
	active.Parent().Find("table.queue td").Each(func(i int, s *goquery.Selection) {
		img := s.Find("img")
		alt := img.AttrOr("alt", "")
		activeID := ogame.ShipName2ID(alt)
		if !activeID.IsSet() {
			activeID = ogame.DefenceName2ID(alt)
			if !activeID.IsSet() {
				return
			}
		}
		activeNbr := utils.ParseInt(s.Text())
		res = append(res, ogame.Quantifiable{ID: activeID, Nbr: activeNbr})
	})
	return res, nil
}

func extractResourcesFromDoc(doc *goquery.Document) ogame.Resources {
	return extractResourcesDetailsFromFullPageFromDoc(doc).Available()
}

func extractResourcesDetailsFromFullPageFromDoc(doc *goquery.Document) ogame.ResourcesDetails {
	out := ogame.ResourcesDetails{}
	metalDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(doc.Find("div#metal_box").AttrOr("title", "")))
	crystalDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(doc.Find("div#crystal_box").AttrOr("title", "")))
	deuteriumDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(doc.Find("div#deuterium_box").AttrOr("title", "")))
	energyDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(doc.Find("div#energy_box").AttrOr("title", "")))
	darkmatterDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(doc.Find("div#darkmatter_box").AttrOr("title", "")))
	out.Metal.Available = utils.ParseInt(metalDoc.Find("table tr").Eq(0).Find("td").Eq(0).Text())
	out.Metal.StorageCapacity = utils.ParseInt(metalDoc.Find("table tr").Eq(1).Find("td").Eq(0).Text())
	out.Metal.CurrentProduction = utils.ParseInt(metalDoc.Find("table tr").Eq(2).Find("td").Eq(0).Text())
	out.Crystal.Available = utils.ParseInt(crystalDoc.Find("table tr").Eq(0).Find("td").Eq(0).Text())
	out.Crystal.StorageCapacity = utils.ParseInt(crystalDoc.Find("table tr").Eq(1).Find("td").Eq(0).Text())
	out.Crystal.CurrentProduction = utils.ParseInt(crystalDoc.Find("table tr").Eq(2).Find("td").Eq(0).Text())
	out.Deuterium.Available = utils.ParseInt(deuteriumDoc.Find("table tr").Eq(0).Find("td").Eq(0).Text())
	out.Deuterium.StorageCapacity = utils.ParseInt(deuteriumDoc.Find("table tr").Eq(1).Find("td").Eq(0).Text())
	out.Deuterium.CurrentProduction = utils.ParseInt(deuteriumDoc.Find("table tr").Eq(2).Find("td").Eq(0).Text())
	out.Energy.Available = utils.ParseInt(energyDoc.Find("table tr").Eq(0).Find("td").Eq(0).Text())
	out.Energy.CurrentProduction = utils.ParseInt(energyDoc.Find("table tr").Eq(1).Find("td").Eq(0).Text())
	out.Energy.Consumption = utils.ParseInt(energyDoc.Find("table tr").Eq(2).Find("td").Eq(0).Text())
	out.Darkmatter.Available = utils.ParseInt(darkmatterDoc.Find("table tr").Eq(0).Find("td").Eq(0).Text())
	out.Darkmatter.Purchased = utils.ParseInt(darkmatterDoc.Find("table tr").Eq(1).Find("td").Eq(0).Text())
	out.Darkmatter.Found = utils.ParseInt(darkmatterDoc.Find("table tr").Eq(2).Find("td").Eq(0).Text())
	return out
}

func extractEspionageReportFromDoc(doc *goquery.Document, location *time.Location) (ogame.EspionageReport, error) {
	report := ogame.EspionageReport{}
	report.ID = utils.DoParseI64(doc.Find("div.detail_msg").AttrOr("data-msg-id", "0"))
	spanLink := doc.Find("span.msg_title a").First()
	txt := spanLink.Text()
	figure := spanLink.Find("figure").First()
	r := regexp.MustCompile(`([^\[]+) \[(\d+):(\d+):(\d+)]`)
	m := r.FindStringSubmatch(txt)
	if len(m) == 5 {
		report.Coordinate.Galaxy = utils.DoParseI64(m[2])
		report.Coordinate.System = utils.DoParseI64(m[3])
		report.Coordinate.Position = utils.DoParseI64(m[4])
	} else {
		return report, errors.New("failed to extract coordinate")
	}
	if figure.HasClass("planet") {
		report.Coordinate.Type = ogame.PlanetType
	} else if figure.HasClass("moon") {
		report.Coordinate.Type = ogame.MoonType
	}
	messageType := ogame.Report
	if doc.Find("span.espionageDefText").Size() > 0 {
		messageType = ogame.Action
	}
	report.Type = messageType
	msgDateRaw := doc.Find("span.msg_date").Text()
	msgDate, _ := time.ParseInLocation("02.01.2006 15:04:05", msgDateRaw, location)
	report.Date = msgDate.In(time.Local)

	username := doc.Find("div.detail_txt").First().Find("span span").First().Text()
	username = strings.TrimSpace(username)
	split := strings.Split(username, "(i")
	if len(split) > 0 {
		report.Username = strings.TrimSpace(split[0])
	}
	characterClassStr := doc.Find("div.detail_txt").Eq(1).Find("span span").First().Text()
	characterClassStr = strings.TrimSpace(characterClassStr)
	report.CharacterClass = v71.GetCharacterClass(characterClassStr)

	report.AllianceClass = ogame.NoAllianceClass
	allianceClassSpan := doc.Find("div.detail_txt").Eq(2).Find("span.alliance_class")
	if allianceClassSpan.HasClass("trader") {
		report.AllianceClass = ogame.Trader
	} else if allianceClassSpan.HasClass("warrior") {
		report.AllianceClass = ogame.Warrior
	} else if allianceClassSpan.HasClass("researcher") {
		report.AllianceClass = ogame.Researcher
	}

	// Bandit, Starlord
	banditstarlord := doc.Find("div.detail_txt").First().Find("span")
	if banditstarlord.HasClass("honorRank") {
		report.IsBandit = banditstarlord.HasClass("rank_bandit1") || banditstarlord.HasClass("rank_bandit2") || banditstarlord.HasClass("rank_bandit3")
		report.IsStarlord = banditstarlord.HasClass("rank_starlord1") || banditstarlord.HasClass("rank_starlord2") || banditstarlord.HasClass("rank_starlord3")
	}

	honorableFound := doc.Find("div.detail_txt").First().Find("span.status_abbr_honorableTarget")
	report.HonorableTarget = honorableFound.Length() > 0

	// IsInactive, IsLongInactive
	inactive := doc.Find("div.detail_txt").First().Find("span")
	if inactive.HasClass("status_abbr_longinactive") {
		report.IsInactive = true
		report.IsLongInactive = true
	} else if inactive.HasClass("status_abbr_inactive") {
		report.IsInactive = true
	}

	// APIKey
	apikey, _ := doc.Find("span.icon_apikey").Attr("title")
	apiDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(apikey))
	report.APIKey = apiDoc.Find("input").First().AttrOr("value", "")

	// Inactivity timer
	activity := doc.Find("div.detail_txt").Eq(3).Find("font")
	if len(activity.Text()) == 2 {
		report.LastActivity = utils.ParseInt(activity.Text())
	}

	// CounterEspionage
	ceTxt := doc.Find("div.detail_txt").Eq(2).Text()
	m1 := regexp.MustCompile(`(\d+)%`).FindStringSubmatch(ceTxt)
	if len(m1) == 2 {
		report.CounterEspionage = utils.DoParseI64(m1[1])
	}

	hasError := false
	resourcesFound := false
	buildingsFound := false
	doc.Find("ul.detail_list").Each(func(i int, s *goquery.Selection) {
		dataType := s.AttrOr("data-type", "")
		if dataType == "resources" && !resourcesFound {
			resourcesFound = true
			report.Metal = utils.ParseInt(s.Find("li").Eq(0).AttrOr("title", "0"))
			report.Crystal = utils.ParseInt(s.Find("li").Eq(1).AttrOr("title", "0"))
			report.Deuterium = utils.ParseInt(s.Find("li").Eq(2).AttrOr("title", "0"))
			report.Energy = utils.ParseInt(s.Find("li").Eq(3).AttrOr("title", "0"))
		} else if dataType == "buildings" && !buildingsFound {
			buildingsFound = true
			report.HasBuildingsInformation = s.Find("li.detail_list_fail").Size() == 0
			s.Find("li.detail_list_el").EachWithBreak(func(i int, s2 *goquery.Selection) bool {
				img := s2.Find("img")
				if img.Size() == 0 {
					hasError = true
					return false
				}
				imgClass := img.AttrOr("class", "")
				r := regexp.MustCompile(`building(\d+)`)
				buildingID := utils.DoParseI64(r.FindStringSubmatch(imgClass)[1])
				l := utils.ParseInt(s2.Find("span.fright").Text())
				level := &l
				switch ogame.ID(buildingID) {
				case ogame.MetalMine.ID:
					report.MetalMine = level
				case ogame.CrystalMine.ID:
					report.CrystalMine = level
				case ogame.DeuteriumSynthesizer.ID:
					report.DeuteriumSynthesizer = level
				case ogame.SolarPlant.ID:
					report.SolarPlant = level
				case ogame.FusionReactor.ID:
					report.FusionReactor = level
				case ogame.MetalStorage.ID:
					report.MetalStorage = level
				case ogame.CrystalStorage.ID:
					report.CrystalStorage = level
				case ogame.DeuteriumTank.ID:
					report.DeuteriumTank = level
				case ogame.AllianceDepot.ID:
					report.AllianceDepot = level
				case ogame.RoboticsFactory.ID:
					report.RoboticsFactory = level
				case ogame.Shipyard.ID:
					report.Shipyard = level
				case ogame.ResearchLab.ID:
					report.ResearchLab = level
				case ogame.MissileSilo.ID:
					report.MissileSilo = level
				case ogame.NaniteFactory.ID:
					report.NaniteFactory = level
				case ogame.Terraformer.ID:
					report.Terraformer = level
				case ogame.SpaceDock.ID:
					report.SpaceDock = level
				case ogame.LunarBase.ID:
					report.LunarBase = level
				case ogame.SensorPhalanx.ID:
					report.SensorPhalanx = level
				case ogame.JumpGate.ID:
					report.JumpGate = level
				}
				return true
			})
		} else if dataType == "research" {
			report.HasResearchesInformation = s.Find("li.detail_list_fail").Size() == 0
			s.Find("li.detail_list_el").EachWithBreak(func(i int, s2 *goquery.Selection) bool {
				img := s2.Find("img")
				if img.Size() == 0 {
					hasError = true
					return false
				}
				imgClass := img.AttrOr("class", "")
				r := regexp.MustCompile(`research(\d+)`)
				researchID := utils.DoParseI64(r.FindStringSubmatch(imgClass)[1])
				l := utils.ParseInt(s2.Find("span.fright").Text())
				level := &l
				switch ogame.ID(researchID) {
				case ogame.EspionageTechnology.ID:
					report.EspionageTechnology = level
				case ogame.ComputerTechnology.ID:
					report.ComputerTechnology = level
				case ogame.WeaponsTechnology.ID:
					report.WeaponsTechnology = level
				case ogame.ShieldingTechnology.ID:
					report.ShieldingTechnology = level
				case ogame.ArmourTechnology.ID:
					report.ArmourTechnology = level
				case ogame.EnergyTechnology.ID:
					report.EnergyTechnology = level
				case ogame.HyperspaceTechnology.ID:
					report.HyperspaceTechnology = level
				case ogame.CombustionDrive.ID:
					report.CombustionDrive = level
				case ogame.ImpulseDrive.ID:
					report.ImpulseDrive = level
				case ogame.HyperspaceDrive.ID:
					report.HyperspaceDrive = level
				case ogame.LaserTechnology.ID:
					report.LaserTechnology = level
				case ogame.IonTechnology.ID:
					report.IonTechnology = level
				case ogame.PlasmaTechnology.ID:
					report.PlasmaTechnology = level
				case ogame.IntergalacticResearchNetwork.ID:
					report.IntergalacticResearchNetwork = level
				case ogame.Astrophysics.ID:
					report.Astrophysics = level
				case ogame.GravitonTechnology.ID:
					report.GravitonTechnology = level
				}
				return true
			})
		} else if dataType == "ships" {
			report.HasFleetInformation = s.Find("li.detail_list_fail").Size() == 0
			s.Find("li.detail_list_el").EachWithBreak(func(i int, s2 *goquery.Selection) bool {
				img := s2.Find("img")
				if img.Size() == 0 {
					hasError = true
					return false
				}
				imgClass := img.AttrOr("class", "")
				r := regexp.MustCompile(`tech(\d+)`)
				shipID := utils.DoParseI64(r.FindStringSubmatch(imgClass)[1])
				l := utils.ParseInt(s2.Find("span.fright").Text())
				level := &l
				switch ogame.ID(shipID) {
				case ogame.SmallCargo.ID:
					report.SmallCargo = level
				case ogame.LargeCargo.ID:
					report.LargeCargo = level
				case ogame.LightFighter.ID:
					report.LightFighter = level
				case ogame.HeavyFighter.ID:
					report.HeavyFighter = level
				case ogame.Cruiser.ID:
					report.Cruiser = level
				case ogame.Battleship.ID:
					report.Battleship = level
				case ogame.ColonyShip.ID:
					report.ColonyShip = level
				case ogame.Recycler.ID:
					report.Recycler = level
				case ogame.EspionageProbe.ID:
					report.EspionageProbe = level
				case ogame.Bomber.ID:
					report.Bomber = level
				case ogame.SolarSatellite.ID:
					report.SolarSatellite = level
				case ogame.Destroyer.ID:
					report.Destroyer = level
				case ogame.Deathstar.ID:
					report.Deathstar = level
				case ogame.Battlecruiser.ID:
					report.Battlecruiser = level
				case ogame.Crawler.ID:
					report.Crawler = level
				case ogame.Reaper.ID:
					report.Reaper = level
				case ogame.Pathfinder.ID:
					report.Pathfinder = level
				}
				return true
			})
		} else if dataType == "defense" {
			report.HasDefensesInformation = s.Find("li.detail_list_fail").Size() == 0
			s.Find("li.detail_list_el").EachWithBreak(func(i int, s2 *goquery.Selection) bool {
				img := s2.Find("img")
				if img.Size() == 0 {
					hasError = true
					return false
				}
				imgClass := img.AttrOr("class", "")
				r := regexp.MustCompile(`defense(\d+)`)
				defenceID := utils.DoParseI64(r.FindStringSubmatch(imgClass)[1])
				l := utils.ParseInt(s2.Find("span.fright").Text())
				level := &l
				switch ogame.ID(defenceID) {
				case ogame.RocketLauncher.ID:
					report.RocketLauncher = level
				case ogame.LightLaser.ID:
					report.LightLaser = level
				case ogame.HeavyLaser.ID:
					report.HeavyLaser = level
				case ogame.GaussCannon.ID:
					report.GaussCannon = level
				case ogame.IonCannon.ID:
					report.IonCannon = level
				case ogame.PlasmaTurret.ID:
					report.PlasmaTurret = level
				case ogame.SmallShieldDome.ID:
					report.SmallShieldDome = level
				case ogame.LargeShieldDome.ID:
					report.LargeShieldDome = level
				case ogame.AntiBallisticMissiles.ID:
					report.AntiBallisticMissiles = level
				case ogame.InterplanetaryMissiles.ID:
					report.InterplanetaryMissiles = level
				}
				return true
			})
		}
	})
	if hasError {
		return report, ogame.ErrDeactivateHidePictures
	}
	return report, nil
}