package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tasiov/ujux/cmd/utils"
)

// heightInfoCmd represents the heightInfo command
var heightInfoCmd = &cobra.Command{
	Use:   "height-info",
	Short: "Find the first and last block height in the database",
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := viper.GetString("db-path")
		base, height, err := utils.GetBaseHeight(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("base: ", base)
		fmt.Println("height: ", height)
	},
}

func init() {
	indexCmd.AddCommand(heightInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// heightInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// heightInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
