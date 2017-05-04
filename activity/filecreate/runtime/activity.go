package filecreate

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-tibco-log")

const (
	ivMessage = "fileName"
	ovMessage = "created"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

//fileCreateActivity
type fileCreateActivity struct {
	metadata *activity.Metadata
}

func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&fileCreateActivity{metadata: md})
}

// Metadata returns the activity's metadata
func (a *fileCreateActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *fileCreateActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	fileName := context.GetInput(ivMessage).(string)
	w, err := os.Create(fileName)
	if err != nil {
		context.SetOutput(ovMessage, "false")
		return false, err
	}
	w.Close()
	context.SetOutput(ovMessage, "true")
	return true, nil
}
