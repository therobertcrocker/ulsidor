package quests

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/core"
)

var coreInstance *core.Core

func NewQuestsCmd(core *core.Core) *cobra.Command {
	coreInstance = core

	questsCmd := &cobra.Command{
		Use:   "quests",
		Short: "Commands for managing quests",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			coreInstance.InitQuestComponents()
			// ... other code ...
		},
	}

	createCmd := NewCreateCmd()
	expCreateCmd := NewTestCreateCmd()
	questsCmd.AddCommand(createCmd)
	questsCmd.AddCommand(expCreateCmd)

	return questsCmd
}
