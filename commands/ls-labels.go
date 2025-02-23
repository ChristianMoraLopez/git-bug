package commands

import (
	"github.com/spf13/cobra"
)

func newLsLabelCommand() *cobra.Command {
	env := newEnv()

	cmd := &cobra.Command{
		Use:   "ls-label",
		Short: "List valid labels.",
		Long: `List valid labels.

Note: in the future, a proper label policy could be implemented where valid labels are defined in a configuration file. Until that, the default behavior is to return the list of labels already used.`,
		PreRunE: loadBackend(env),
		RunE: closeBackend(env, func(cmd *cobra.Command, args []string) error {
			return runLabelLs(env)
		}),
		Deprecated: ` and will be removed in v1.0.

The functionality provided by this command is now provided by
the following (equivalent) command:
git-bug label ls
`,
	}

	return cmd
}
