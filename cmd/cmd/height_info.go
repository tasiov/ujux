package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tasiov/ujux/cmd/utils"
)

var heightInfoCmd = &cobra.Command{
	Use:   "height-info",
	Short: "Find the first and last block height in the database",
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := viper.GetString("db-path")

		blockStore, err := utils.NewBlockStore(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		blockStoreState, err := blockStore.LoadBlockStoreState()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("base: ", blockStoreState.Base)
		fmt.Println("height: ", blockStoreState.Height)
	},
}

func init() {
	indexCmd.AddCommand(heightInfoCmd)
}
