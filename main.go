package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func confirm(message string, count, maxAttempts int, verbose bool) bool {
	r := bufio.NewReader(os.Stdin)
	n := 0
	for {
		if n >= count {
			// reached required confirm count
			return true
		}
		attempt := 0
		for {
			if attempt >= maxAttempts {
				// reached maximum number of attempts
				return false
			}
			if verbose {
				fmt.Printf(
					"%s (%d/%d, attempt %d/%d) [y/n]: ",
					message,
					n+1,
					count,
					attempt+1,
					maxAttempts,
				)
			} else {
				fmt.Printf("%s [y/n]: ", message)
			}
			attempt++

			res, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			// skip empty input (i.e. "\n")
			if len(res) < 2 {
				continue
			}
			confirmed := strings.ToLower(strings.TrimSpace(res))[0] == 'y'
			if verbose {
				fmt.Printf("confirmed=%v\n", confirmed)
			}
			if !confirmed {
				return false
			} else {
				n++
				break
			}
		}
	}
}

func main() {
	message := "confirm"
	tries := 1
	count := 1
	verbose := false
	rootCmd := &cobra.Command{
		Use:   "confirm",
		Short: "confirm an action with yes/no",
		Long:  "A dead simple CLI app that asks the user for a yes/no confirmation.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				message = args[0]
			}
			if confirm(message, count, tries, verbose) {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		},
	}
	rootCmd.PersistentFlags().
		IntVarP(&tries, "tries", "t", 1, "maximum number of tries")
	rootCmd.PersistentFlags().
		IntVarP(&count, "count", "n", 1, "number of required confirmations")
	rootCmd.PersistentFlags().
		BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
