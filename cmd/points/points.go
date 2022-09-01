/*
Copyright © 2022 Muxable

*/
package points

import (
	"github.com/spf13/cobra"
)

// PointsCmd represents the points command
var PointsCmd = &cobra.Command{
	Use:   "points",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// rootCmd.AddCommand(pointsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pointsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pointsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
