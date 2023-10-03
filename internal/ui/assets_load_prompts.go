package ui

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gocarina/gocsv"
	"github.com/therobertcrocker/ulsidor/internal/machinations"
	"github.com/therobertcrocker/ulsidor/utils"
)

// LoadAssetsFromJSON loads assets from a JSON file.
func LoadAssetsFromJSON(filePath string) ([]*machinations.Asset, error) {

	//display filepath to user
	utils.Log.WithField("filepath", filePath).Info("loading assets from file")

	//open file and read contents into []*Asset
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//unmarshal json file into []*Asset
	var assets []*machinations.Asset
	err = json.NewDecoder(file).Decode(&assets)
	if err != nil {
		utils.Log.WithError(err).Error("failed to unmarshal json file")
		return nil, err
	}

	utils.Log.WithField("count", len(assets)).Info("loaded assets from file")

	// display assets to user, one line per asset
	for i, asset := range assets {
		row := fmt.Sprintf("%d: %s", i, asset.Display())
		fmt.Println(row)
	}

	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: "Are you sure you want to load these assets?",
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		return nil, err
	}

	if !confirm {
		utils.Log.Info("user cancelled load")
		return nil, nil
	}

	return assets, nil
}

type AssetCSV struct {
	Name         string `csv:"name"`
	Description  string `csv:"description"`
	Type         string `csv:"type"`
	Cost         int    `csv:"cost"`
	Upkeep       int    `csv:"upkeep"`
	Threshold    int    `csv:"threshold"`
	HitPoints    int    `csv:"hit_points"`
	Arcane       string `csv:"arcane"`
	AttackTrait  string `csv:"attack_trait"`
	DefendTrait  string `csv:"defend_trait"`
	AttackCount  int    `csv:"attack_count"`
	AttackDice   int    `csv:"attack_dice"`
	AttackMod    int    `csv:"attack_mod"`
	CounterCount int    `csv:"counter_count"`
	CounterDice  int    `csv:"counter_dice"`
	CounterMod   int    `csv:"counter_mod"`
	QualityOne   string `csv:"quality_1"`
	QualityTwo   string `csv:"quality_2"`
}

// LoadAssetsFromCSV loads assets from a CSV file.
func LoadAssetsFromCSV(filepath string) ([]*machinations.Asset, error) {

	// display filepath to user
	utils.Log.WithField("filepath", filepath).Info("loading assets from file")

	// open file and read contents into []*Asset
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// unmarshal csv file into []*AssetCSV
	var assetsRaw []*AssetCSV
	if err := gocsv.UnmarshalFile(file, &assetsRaw); err != nil {
		return nil, err
	}

	// convert []*AssetCSV to []*Asset
	var assets []*machinations.Asset
	for _, assetRaw := range assetsRaw {
		asset := &machinations.Asset{
			Name:        assetRaw.Name,
			Description: assetRaw.Description,
			Type:        assetRaw.Type,
			Cost:        assetRaw.Cost,
			Upkeep:      assetRaw.Upkeep,
			Threshold:   assetRaw.Threshold,
			HitPoints:   assetRaw.HitPoints,
			Arcane:      assetRaw.Arcane,
		}

		if assetRaw.AttackTrait != "None" {
			asset.HasAttack = true
			asset.Attack = &machinations.Attack{
				AttackerTrait: assetRaw.AttackTrait,
				DefenderTrait: assetRaw.DefendTrait,
				Damage: machinations.Damage{
					Count:    assetRaw.AttackCount,
					Dice:     assetRaw.AttackDice,
					Modifier: assetRaw.AttackMod,
				},
			}
		}

		if assetRaw.CounterCount != 0 {
			asset.HasCounter = true
			asset.Counter = &machinations.Damage{
				Count:    assetRaw.CounterCount,
				Dice:     assetRaw.CounterDice,
				Modifier: assetRaw.CounterMod,
			}
		}

		var qualities []string
		if assetRaw.QualityOne != "None" {
			qualities = append(qualities, assetRaw.QualityOne)
		}
		if assetRaw.QualityTwo != "None" {
			qualities = append(qualities, assetRaw.QualityTwo)
		}
		asset.Qualities = qualities

		assets = append(assets, asset)

	}

	// display assets to user, one line per asset
	for i, asset := range assets {
		row := fmt.Sprintf("%d: %s", i, asset.Display())
		fmt.Println(row)
	}

	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: "Are you sure you want to load these assets?",
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		return nil, err
	}

	if !confirm {
		utils.Log.Info("user cancelled load")
		return nil, nil
	}

	return assets, nil

}
