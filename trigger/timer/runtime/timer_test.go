package timer

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

const testConfig3 string = `{
  "name": "tibco-timer",
  "settings": {
  },
  "endpoints": [
    {
      "flowURI": "local://testFlow2",
      "settings": {
        "repeating": "false"
      }
    }
  ]
}`

const testConfig string = `{
  "name": "tibco-timer",
  "settings": {
  },
  "endpoints": [
    {
      "flowURI": "local://testFlow2",
      "settings": {
        "repeating": "false",
        "startDate" : "2016-05-03T19:25:00Z-04:00"
      }
    }
  ]
}`

const testConfig2 string = `{
  "name": "tibco-timer",
  "settings": {
  },
  "endpoints": [
    {
      "flowURI": "local://testFlow2",
      "settings": {
      	"notImmediate": "false",
        "repeating": "true",
        "seconds": "5"
      }
    }
  ]
}`

const testConfig4 string = `{
  "name": "tibco-timer",
  "settings": {
  },
  "endpoints": [
    {
      "flowURI": "local://testFlow",
      "settings": {
        "repeating": "false",
        "startDate" : "05/01/2016, 12:25:01"
      }
    },
    {
      "flowURI": "local://testFlow2",
      "settings": {
        "repeating": "true",
        "startDate" : "05/01/2016, 12:25:01",
        "hours": "24"
      }
    },
    {
      "flowURI": "local://testFlow3",
      "settings": {
        "repeating": "true",
        "startDate" : "05/01/2016, 12:25:01",
        "minutes": "60"
      }
    },
    {
      "flowURI": "local://testFlow3",
      "settings": {
        "repeating": "true",
        "startDate" : "05/01/2016, 12:25:01",
        "seconds": "30"
      }
    }
  ]
}`

type TestRunner struct {
}

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
	log.Debugf("Ran Action: %v", uri)
	return 0, nil, nil
}

func TestRegistered(t *testing.T) {
	act := trigger.Get("tibco-timer")

	if act == nil {
		t.Error("Timer Trigger Not Registered")
		t.Fail()
		return
	}
}

func TestInit(t *testing.T) {
	tgr := trigger.Get("tibco-timer")

	runner := &TestRunner{}

	config := &trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)
	tgr.Init(config, runner)
}

func TestTimer(t *testing.T) {

	log.Debugf("TestTimer")
	tgr := trigger.Get("tibco-timer")

	tgr.Start()
	time.Sleep(time.Second * 2)
	defer tgr.Stop()

	log.Debug("Test timer done")
}
