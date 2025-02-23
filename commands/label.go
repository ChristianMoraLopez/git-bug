package commands

import (
	"github.com/spf13/cobra"

	_select "github.com/MichaelMure/git-bug/commands/select"
)

func newLabelCommand() *cobra.Command {
	env := newEnv()

	cmd := &cobra.Command{
		Use:     "label [ID]",
		Short:   "Display, add or remove labels to/from a bug.",
		PreRunE: loadBackend(env),
		RunE: closeBackend(env, func(cmd *cobra.Command, args []string) error {
			return runLabel(env, args)
		}),
		ValidArgsFunction: completeBug(env),
	}

	cmd.AddCommand(newLabelAddCommand())
	cmd.AddCommand(newLabelLsCommand())
	cmd.AddCommand(newLabelRmCommand())

	return cmd
}

func runLabel(env *Env, args []string) error {
	b, _, err := _select.ResolveBug(env.backend, args)
	if err != nil {
		return err
	}

	snap := b.Snapshot()

	for _, l := range snap.Labels {
		env.out.Println(l)
	}

	return nil
}
