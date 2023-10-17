package cmd

import (
	"github.com/spf13/cobra"
)

var rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "Index a range of blocks",
	Run: func(cmd *cobra.Command, args []string) {
		// dbPath := viper.GetString("db-path")

		// blockstoreDb, err := utils.NewLevelDb(dbPath, "blockstore")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// blockstoreState, err := utils.GetBlockStoreState(blockstoreDb)

		// fmt.Println("base: ", blockstoreState.Base)
		// fmt.Println("height: ", blockstoreState.Height)
	},
}

func init() {
	indexCmd.AddCommand(rangeCmd)
}
