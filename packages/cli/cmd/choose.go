// Package cmd ..
package cmd

import (
	"chess960-cli/clients"
	"chess960-cli/data"
	"chess960-cli/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// chooseCmd represents the choose command
var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "Choose a random position",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.LogProgramName()
		// Get URL
		fmt.Print("Position: ")
		var positionIndex int
		fmt.Scanln(&positionIndex)
		var position = data.Positions[positionIndex-1]
		fmt.Printf("Position %d: %s\n", positionIndex, position)

		// ---- convert to FEN ----
		// Assume rank 8 is the back rank, pawns on rank 7, empty elsewhere
		fen := fmt.Sprintf("%s/PPPPPPPP/8/8/8/8/pppppppp/%s w KQkq - 0 1",
			position, strings.ToLower(position))
		fmt.Println("FEN:", fen)

		// ---- lichess.org Cloud Eval ----
		eval, err := clients.CloudEvalCP(fen, "chess960")
		if err != nil {
			return fmt.Errorf("❌ failed to evaluate position: %w", err)
		}

		fmt.Println("Evaluation (centipawns):", eval)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(chooseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chooseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chooseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
