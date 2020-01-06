package archive

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/argoproj/argo/cmd/argo/commands/client"
	"github.com/argoproj/argo/cmd/server/workflowarchive"
)

func NewResubmitCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "resubmit NAMESPACE UID",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				cmd.HelpFunc()(cmd, args)
				os.Exit(1)
			}
			namespace := args[0]
			uid := args[1]
			conn := client.GetClientConn()
			ctx := client.ContextWithAuthorization()
			client := workflowarchive.NewArchivedWorkflowServiceClient(conn)
			wf, err := client.ResubmitArchivedWorkflow(ctx, &workflowarchive.ResubmitArchivedWorkflowRequest{
				Namespace: namespace,
				Uid:       uid,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Archived workflow '%s' resubmitted\n", wf.Name)
		},
	}
	return command
}