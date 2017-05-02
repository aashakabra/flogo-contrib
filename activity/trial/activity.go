package trial

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-tibco-log")

const (
	ivMessage = "message"

	ovMessage = "message"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// LogActivity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type LogActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LogActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *LogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *LogActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)

	file, err := os.Create(message)
	if err != nil {
		return false, err
	}
	file.Close()

	context.SetOutput(ovMessage, "true")
	return true, nil
}
