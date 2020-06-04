package cmd

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan all keys matching pattern",
	Long:  `Scan all keys matching pattern`,
	Run: func(cmd *cobra.Command, args []string) {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		ctx := context.Background()
		pattern := ""
		if len(args) > 0 {
			pattern = args[0]
		}
		iter := client.Scan(ctx, 0, pattern, 0).Iterator()
		for iter.Next(ctx) {
			//		val, _ := client.Get(ctx, iter.Val()).Result()

			//		fmt.Println(iter.Val(), "=>", val)
			fmt.Println(iter.Val())
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
