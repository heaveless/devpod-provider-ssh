package cmd

import (
	"context"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
	"github.com/heaveless/devpod-provider-ssh/utils"
)

type InitCmd struct {}

func NewInitCmd() *cobra.Command {
	cmd := &InitCmd {}
	initCmd := &cobra.Command {
		Use: "init",
		Short: "Init account",
		RunE: func(_ *cobra.Command, args []string) error {
			sshProvider, err := utils.NewProvider(log.Default)
			if err != nil {
				return err
			}

			return cmd.Run(
				context.Background(),
				sshProvider,
				log.Default,
			)
		},
	}

	return initCmd
}

func (cmd *InitCmd) Run(
	ctx context.Context,
	providerSSH *utils.SSHProvider,
	logs log.Logger,
) error {
	err := utils.Init(providerSSH)
	if err != nil {
		return err
	}
	return nil
}