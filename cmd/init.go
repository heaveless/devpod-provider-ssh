package cmd

import (
	"context"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type InitCmd struct {}

func NewInitCmd() *cobra.Command {
	cmd := &InitCmd {}
	initCmd := &cobra.Command {
		Use: "init",
		Short: "Init account",
		RunE: func(_ *cobra.Command, args []string) error {
			sshProvider, err := ssh.NewProvider(log.Default)
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
	providerSSH *ssh.sshProvider,
	logs log.Logger,
) error {
	ssh.Init(providerSSH)
}