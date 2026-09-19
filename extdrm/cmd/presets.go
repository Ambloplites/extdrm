package cmd

import (
	"fmt"
	"slices"

	"github.com/Ambloplites/extdrm"
	"github.com/spf13/cobra"
)

// presetsCmd represents the presets command
var presetsCmd = &cobra.Command{
	Use:   "presets",
	Short: "List built-in presets",
	Run:   runPresets,
}

func init() {
	rootCmd.AddCommand(presetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// presetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// presetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runPresets(cmd *cobra.Command, args []string) {
	names := extdrm.BuiltinPresetNames()
	slices.Sort(names)
	for _, name := range names {
		fmt.Println(name)
	}
}
