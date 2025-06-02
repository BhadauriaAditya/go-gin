package commands

import (
	"github.com/spf13/cobra"
)

var sendRemindersCmd = &cobra.Command{
	Use:   "task:send-reminders",
	Short: "Send reminders for tasks",
	Long: `Send reminders for tasks that are due or overdue.
This command will process all tasks and send appropriate reminders to users.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Sending reminders...")
	},
}

func init() {
	// Add flags laravel type cli arguments
	// sendRemindersCmd.Flags().StringP("type", "t", "", "Type of reminders to send")
}

// GetSendRemindersCmd returns the send reminders command
func GetSendRemindersCmd() *cobra.Command {
	return sendRemindersCmd
}
