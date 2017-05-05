package filecli

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

const (
	inputFileName = "fileName"
	outputCreated = "created"
)

// MyActivity is a stub for your Activity implementation
type fileCreateActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &fileCreateActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *fileCreateActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *fileCreateActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	fileName := context.GetInput(inputFileName).(string)
	w, err := os.Create(fileName)
	if err != nil {
		context.SetOutput(outputCreated, "false")
		return false, err
	}
	w.Close()
	context.SetOutput(outputCreated, "true")
	return true, nil
}
