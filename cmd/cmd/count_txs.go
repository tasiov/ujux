package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tasiov/ujux/cmd/utils"
)

var countTxsCmd = &cobra.Command{
	Use:   "count-txs",
	Short: "Count the number of transactions in a range of blocks",
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := viper.GetString("db-path")

		blockStore, err := utils.NewBlockStore(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		blockStoreState, err := blockStore.LoadBlockStoreState()

		start, err := cmd.Flags().GetInt64("start")
		if err != nil {
			fmt.Println(err)
			return
		}

		end, err := cmd.Flags().GetInt64("end")
		if err != nil {
			fmt.Println(err)
			return
		}

		if start < blockStoreState.Base {
			start = blockStoreState.Base
		}
		if end > blockStoreState.Height {
			end = blockStoreState.Height
		}
		if start > end {
			fmt.Println("start height is greater than end height")
			return
		}

		fmt.Println("start: ", start)
		fmt.Println("end: ", end)

		stateStore, err := utils.NewStateStore(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i := start; i <= end; i++ {
			_, err := stateStore.LoadABCIResponses(i)
			if err != nil {
				fmt.Println(err)
				return
			}

			block, err := blockStore.LoadBlock(i)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(block)
			break
		}
	},
}

func init() {
	indexCmd.AddCommand(countTxsCmd)
	countTxsCmd.Flags().Int64P("start", "s", 1, "Start height")
	countTxsCmd.Flags().Int64P("end", "e", math.MaxInt64, "End height")
}
