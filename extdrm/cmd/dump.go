package cmd

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Ambloplites/extdrm"
	"github.com/Ambloplites/fsdump"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump SOURCE DESTINATION",
	Short: "Decrypt the contents of an encrypted extdrm filesystem",
	Args:  cobra.MinimumNArgs(2),
	Run:   runDump,
}

func init() {
	rootCmd.AddCommand(dumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dumpCmd.Flags().Int("workers", 0, "Number of workers. Specify a value less than one, and the number of logical CPUs available to the process will be used")
	dumpCmd.Flags().String("preset", "", "Name or path to the preset")
}

func runDump(cmd *cobra.Command, args []string) {
	var config *extdrm.DrmConfig

	workers, _ := cmd.Flags().GetInt("workers")
	src := args[0]
	dest := args[1]

	presetName, _ := cmd.Flags().GetString("preset")
	config, err := readPreset(presetName)
	if err != nil {
		log.Fatalln("failed to read preset:", err)
	}

	start := time.Now()

	ch, err := extdrm.ReadFS(src, config)
	if err != nil {
		log.Fatalln(err)
	}

	dumper := fsdump.Dumper{
		Src:        &fsdump.ChannelFileSource{Chan: ch},
		Dest:       dest,
		NumWorkers: workers,
	}
	dumper.Run()

	log.Println("time elapsed:", time.Since(start))
}

func readPreset(name string) (*extdrm.DrmConfig, error) {
	if name == "" {
		// bruteforce will be used
		return nil, nil
	}

	config := extdrm.LookupBuiltinPreset(name)
	if config != nil {
		return config, nil
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	config = &extdrm.DrmConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}
