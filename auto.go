package badgeify

import (
	"encoding/json"
	"os"
)

type lastRun struct {
	Result struct {
		Line float32
	}
}

const lastRunFilePath = "./coverage/.last_run.json"

func LookupCoverage() (float32, error) {
	var r lastRun
	dat, err := os.ReadFile(lastRunFilePath)
	if err != nil {
		return 0.0, err
	}
	err = json.Unmarshal(dat, &r)
	if err != nil {
		return 0.0, err
	}
	return r.Result.Line, nil
}
