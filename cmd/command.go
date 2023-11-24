package cmd

import (
	"context"
	"fmt"
	"os"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"

	"github.com/heaveless/devpod-provider-ssh/utils"
)

type CommandCmd struct {}

func NewCommandCmd() *cobra.Command {
	cmd := &CommandCmd{}
	commandCmd := &cobra.Command{
		Use:   "command",
		Short: "Command an instance",
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

	return commandCmd
}

func (cmd *CommandCmd) Run(
	ctx context.Context,
	providerSSH *utils.SSHProvider,
	logs log.Logger,
) error {
	command := os.Getenv("COMMAND")
	if command == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	return utils.Command(providerSSH, command)
}