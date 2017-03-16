package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/pilosa/pilosa/ctl"
)

func NewCheckCommand(stdin io.Reader, stdout, stderr io.Writer) *cobra.Command {
	checker := ctl.NewCheckCommand(os.Stdin, os.Stdout, os.Stderr)
	checkCmd := &cobra.Command{
		Use:   "check <path> [path2]...",
		Short: "Do a consistency check on a pilosa data file.",
		Long: `
Performs a consistency check on data files.
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("path required")
				return
			}
			checker.Paths = args
			if err := checker.Run(context.Background()); err != nil {
				fmt.Println(err)
			}
		},
	}
	return checkCmd
}

func init() {
	subcommandFns["check"] = NewCheckCommand
}