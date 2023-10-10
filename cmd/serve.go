package cmd

import (
	"fmt"
	"github.com/jkdv-systeme/kyasshu/internal/config"
	"github.com/jkdv-systeme/kyasshu/internal/http"
	"github.com/spf13/cobra"
)

// art generator: https://patorjk.com/software/taag/#p=display&f=Small&t=UPTIMEBUDDY
var banner = fmt.Sprintf(`
 _                   _        
| |___  _ __ _ _____| |_ _  _ 
| / / || / _' (_-<_-< ' \ || |
|_\_\\_, \__,_/__/__/_||_\_,_|
     |__/
v%s
`, config.Version)

var serveCommand = &cobra.Command{
	Use:              "serve",
	TraverseChildren: true,
	Short:            "Start the api server",
	Long:             `Starts the api server`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("\u001B[1;36m" + banner + "\u001B[0m")
		http.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCommand)
}
