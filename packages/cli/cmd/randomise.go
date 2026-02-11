// Package cmd ..
package cmd

import (
	"chess960-cli/clients"
	"chess960-cli/data"
	"chess960-cli/utils"
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

// randomiseCmd represents the randomise command
var randomiseCmd = &cobra.Command{
	Use:   "randomise",
	Short: "Select a random Chess960 position and evaluate it",
	Long: `Selects a random Chess960 (Chess60) position from the built-in list
and evaluates the initial position using Lichess Cloud Eval.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.LogProgramName()

		n := len(data.Positions)

		// ---- pick random position from data.Positions ----
		positionIndex := rand.Intn(n)
		position := data.Positions[positionIndex]
		fmt.Printf("Position %d: %s\n", positionIndex+1, position)

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

/* ----------------------------- Helpers ----------------------------- */

func init() {
	rootCmd.AddCommand(randomiseCmd)
}
