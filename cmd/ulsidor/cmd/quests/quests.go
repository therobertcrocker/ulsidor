package quests

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
)

var coreInstance *core.Core

func NewQuestsCmd(core *core.Core) *cobra.Command {
	coreInstance = core

	questsCmd := &cobra.Command{
		Use:   "quests",
		Short: "Commands for managing quests",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := coreInstance.InitQuestComponents()
			if err != nil {
				utils.Log.Error(err)
			}
		},
	}

	createCmd := NewCreateCmd()
	questsCmd.AddCommand(createCmd)

	return questsCmd
}
