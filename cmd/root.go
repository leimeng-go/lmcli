package cmd
import(
	"github.com/spf13/cobra"
)
func init() {
   rootCmd.AddCommand(wordCmd)
   rootCmd.AddCommand(timeCmd)
}

var (
	rootCmd=cobra.Command{
		Use:"",
	}
)

func Execute() error{
	return rootCmd.Execute()
}
