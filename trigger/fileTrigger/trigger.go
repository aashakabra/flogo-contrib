package fileTrigger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct {
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata: md}
}

//New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config: config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
}

// Init implements trigger.Trigger.Init
func (t *MyTrigger) Init(runner action.Runner) {
	if t.config.Settings == nil {
		panic(fmt.Sprintf("No Settings found for trigger '%s'", t.config.Id))
	}

	if _, ok := t.config.Settings["fileDir"]; !ok {
		panic(fmt.Sprintf("No FileDir found for trigger '%s' in settings", t.config.Id))
	}
	t.runner = runner
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	// start the trigger
	dir := ":" + t.config.GetSetting("fileDir")
	fmt.Println("Directory to scan is ", dir)
	scan(dir)

	handlers := t.config.Handlers
	for _, handler := range handlers {
		act := action.Get(handler.ActionId)
		_, _, err := t.runner.Run(context.Background(), act, handler.ActionId, nil)
		if err != nil {
			panic(err)
		}

	}
	return nil
}

func scan(dir string) {
	for {
		_, err := os.Stat(dir)
		if err == nil {
			fmt.Println("Success")
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("Asha")
	}
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}
