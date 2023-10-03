package commands

import (
	"fmt"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/machinations"
	"github.com/therobertcrocker/ulsidor/internal/ui"
	"github.com/therobertcrocker/ulsidor/utils"
)

var (
	assetsFilePathFlag string
	assetIDFlag        string
)

func NewFactionsCmd() *cobra.Command {
	factionsCmd := &cobra.Command{
		Use:   "factions",
		Short: "Commands for managing factions",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Faction Management")
		},
	}

	AddCmd := NewAddFactionCmd()
	AssetsCmd := NewAssetsCmd()
	factionsCmd.AddCommand(AddCmd)
	factionsCmd.AddCommand(AssetsCmd)

	return factionsCmd

}

/*
===============================================================================

	FACTION COMMANDs

===============================================================================
*/
func NewAddFactionCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new faction",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Create Faction")

			factionInput, err := ui.CollectFactionInput()
			if err != nil {
				utils.Log.WithError(err).Error("failed to collect faction input")
			}

			if err := coreInstance.Factions.Codex.AddFaction(factionInput); err != nil {
				utils.Log.WithError(err).Error("failed to add faction")
			}

		},
	}

	return addCmd
}

/*
===============================================================================
                                 ASSET COMMANDs
===============================================================================
*/

// NewAssetsCmd manages the commands for the Asset sub-application.
// TODO: add a ui for selecting a sub-command if none is provided.

func NewAssetsCmd() *cobra.Command {
	assetsCmd := &cobra.Command{
		Use:   "assets",
		Short: "Commands for managing assets",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Asset Management")
		},
	}
	AddAssetCmd := NewAddAssetCmd()
	LoadAssetsCmd := NewLoadAssetsCmd()
	GetAssetCmd := NewGetAssetCmd()
	ClearAssetsCmd := NewClearAssetsCmd()

	assetsCmd.AddCommand(AddAssetCmd)
	assetsCmd.AddCommand(LoadAssetsCmd)
	assetsCmd.AddCommand(GetAssetCmd)
	assetsCmd.AddCommand(ClearAssetsCmd)
	return assetsCmd
}

// NewAddAssetCmd allows the user to add a new asset through an interactive UI.
func NewAddAssetCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new asset",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Create Asset")

			assetInput, err := ui.CollectAssetInput()
			if err != nil {
				utils.Log.WithError(err).Error("failed to collect asset input")
			}

			coreInstance.Factions.Assets.AddAsset(assetInput)

		},
	}

	return addCmd
}

// NewLoadAssetsCmd allows the user to load assets from a file.
func NewLoadAssetsCmd() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "Load assets from a file",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Load Assets")

			filetype := filepath.Ext(assetsFilePathFlag)

			var assets []*machinations.Asset
			var err error

			switch filetype {
			case ".json":
				assets, err = ui.LoadAssetsFromJSON(assetsFilePathFlag)
				if err != nil {
					utils.Log.WithError(err).Error("failed to load assets")
				}
			case ".csv":
				assets, err = ui.LoadAssetsFromCSV(assetsFilePathFlag)
				if err != nil {
					utils.Log.WithError(err).Error("failed to load assets")
				}
			default:
				utils.Log.WithField("filetype", filetype).Error("unsupported filetype")
			}

			if assets == nil {
				utils.Log.Info("user cancelled load")
				return
			}

			if err := coreInstance.Factions.Assets.LoadAssets(assets); err != nil {
				utils.Log.WithError(err).Error("failed to load assets")
				return
			}

			utils.Log.Info("assets loaded successfully")

		},
	}

	loadCmd.Flags().StringVarP(&assetsFilePathFlag, "file", "f", "", "Path to the file containing assets to load")

	return loadCmd
}

// NewGetAssetCmd allows the user to get an asset by ID.
func NewGetAssetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get an asset by ID",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Get Asset")

			asset, err := coreInstance.Factions.Assets.GetAsset(assetIDFlag)
			if err != nil {
				utils.Log.WithError(err).Error("failed to get asset")
				return
			}

			fmt.Println(asset.Display())

		},
	}

	getCmd.Flags().StringVarP(&assetIDFlag, "id", "i", "", "ID of the asset to get")

	return getCmd
}

// NewClearAssetsCmd allows the user to clear all assets from the codex.
func NewClearAssetsCmd() *cobra.Command {
	clearCmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear all assets from the codex",
		Run: func(cmd *cobra.Command, args []string) {
			printTitle("Machinations - Clear Assets")

			var confirm bool
			confirmPrompt := &survey.Confirm{
				Message: "Are you sure you want to clear all assets?",
			}
			if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
				utils.Log.WithError(err).Error("failed to ask user for confirmation")
				return
			}

			if !confirm {
				utils.Log.Info("user cancelled clear")
				return
			}

			coreInstance.Factions.Assets.Clear()

			utils.Log.Info("assets cleared successfully")

		},
	}

	return clearCmd
}
