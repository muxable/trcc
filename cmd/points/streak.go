/*
Copyright © 2022 Muxable

*/
package points

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	streak int
	// tier    int
	// minutes int
	// raid    int
)

// streakCmd represents the streak command
var streakCmd = &cobra.Command{
	Use:   "streak",
	Short: "Number of consecutive streams watched",
	Long:  `Number of consecutive streams watched for at least five minutes. Each stream must be at least 10 minutes long and it must have been at least 30 minutes since the last stream ended.`,
	Run: func(cmd *cobra.Command, args []string) {

		if streak <= 1 {
			fmt.Println("0")
		} else if streak >= 5 {
			fmt.Println("450")
		} else {
			switch streak {
			case 2:
				fmt.Println("300")
			case 3:
				fmt.Println("350")
			case 4:
				fmt.Println("400")
			}
		}
	},
}

func init() {

	streakCmd.Flags().IntVarP(&streak, "streak", "s", 0, "Days consecutively watching")

	if err := streakCmd.MarkFlagRequired("streak"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	PointsCmd.AddCommand(streakCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streakCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streakCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
