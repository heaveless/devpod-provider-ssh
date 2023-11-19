package cmd

import (
	"os"
	"os/exec"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

func NewRootCmd() *cobra.Command {
	sshCmd := &cobra.Command {
		Use: "devpod-provider-ssh",
		Short: "ssh Provider Commands",
		SilenceErrors: true,
		SilenceUsage: true,
		PersistentPreRunE: func(cobraCmd *cobra.Command, args []string) error {
			log.Default.MakeRaw()
			return nil
		},
	}

	return sshCmd
}

func BuildRoot() *cobra.Command {
	rootCmd := NewRootCmd()

	rootCmd.AddCommand(NewInitCmd())
	rootCmd.AddCommand(NewCommandCmd())

	return rootCmd
}

func Execute() {
	rootCmd := BuildRoot()

	err := rootCmd.Execute()
	if err != nil {
		if exitErr, ok := err.(*ssh.ExitError); ok {
			os.Exit(exitErr.ExitStatus())
		}

		if exitErr, ok := err.(*exec.ExitError); ok {
			if len(exitErr.Stderr) > 0 {
				log.Default.ErrorStreamOnly().Error(string(exitErr.Stderr))
			}
			os.Exit(exitErr.ExitCode())
		}

		log.Default.Fatal(err)
	}
}