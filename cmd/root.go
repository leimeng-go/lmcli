package cmd
import(
	"github.com/spf13/cobra"
)
func init() {

}

var (
	rootCmd=cobra.Command{
		Use:"",
	}
)

func Execute() error{
	return rootCmd.Execute()
}
