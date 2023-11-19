package cmd

import (
	"context"
	"fmt"
	"os"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type CommandCmd struct {}

func NewCommandCmd() *cobra.Command {
	cmd := &CommandCmd{}
	commandCmd := &cobra.Command{
		Use:   "command",
		Short: "Command an instance",
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

	return commandCmd
}

func (cmd *CommandCmd) Run(
	ctx context.Context,
	providerSSH *ssh.SSHProvider,
	logs log.Logger,
) error {
	command := os.Getenv("COMMAND")
	if command == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	return ssh.Command(providerSSH, command)
}