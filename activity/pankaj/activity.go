package pankaj

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// MyActivity is a stub for your Activity implementation
type fileCreateActivity struct {
	metadata *activity.Metadata
}

const (
	ivFilename = "fileName"
	ovCreated  = "isCreated"
)

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

	inputFilename := context.GetInput(ivFilename).(string)
	w, err := os.Create(inputFilename)
	if err != nil {
		context.SetOutput(ovCreated, "false")
		return false, err
	}
	w.Close()
	context.SetOutput(ovCreated, "true")
	return true, nil
}
