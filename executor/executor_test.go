package executor

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/scenario"
	"github.com/stretchr/testify/assert"
)

func generate(t *testing.T) (string, string) {
	dir, err := ioutil.TempDir("", "scenario-*")
	assert.Nil(t, err)
	name := filepath.Join(dir, "main.go")
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	assert.Nil(t, err)

	err = Generate(f, 123)
	assert.Nil(t, err)

	return dir, name
}

func TestGenerate(t *testing.T) {
	dir, _ := generate(t)
	os.RemoveAll(dir)
}

// a generated file should be compiled with a valid scenario
func TestCompile(t *testing.T) {
	scenario := `
package main

import (
	"context"

	"github.com/gobench-io/gobench/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   1,
			Rate: 100,
			Fu:   f1,
		},
	}
}

func f1(ctx context.Context, vui int) {
}`

	dir, _ := generate(t)

	defer os.RemoveAll(dir)

	// create scenario.go
	scenarioPath := filepath.Join(dir, "scenario.go")
	scenarioFile, _ := os.OpenFile(scenarioPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	_, _ = scenarioFile.Write([]byte(scenario))

	// create go.mod
	testDir, _ := os.Getwd()
	mainDir, _ := exec.Command("dirname", testDir).CombinedOutput()
	gomod := fmt.Sprintf(`
		module gobench.io/scenario
		replace github.com/gobench-io/gobench => %s
		`, string(mainDir))
	gomodPath := filepath.Join(dir, "go.mod")
	gomodFile, _ := os.OpenFile(gomodPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	_, _ = gomodFile.Write([]byte(gomod))

	out, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s; go build -o main.out", dir)).CombinedOutput()
	assert.Nil(t, err, string(out))
}

func TestStart(t *testing.T) {
	opts := &Options{
		AgentSock:    "/tmp/a1",
		ExecutorSock: "/tmp/e1",
		AppID:        1,
		Vus: scenario.Vus{
			scenario.Vu{
				Nu:   20,
				Rate: 100,
				Fu:   func(ctx context.Context, vui int) {},
			},
		},
	}
	logger := logger.NewNopLogger()

	e, err := NewExecutor(opts, logger)
	assert.Nil(t, err)

	// setup nop metric logger for the driver
	assert.Nil(t, e.driver.SetNopMetricLog())

	er, _ := newExecutorRPC(e)

	args := true
	reply := new(bool)

	err = er.Start(&args, reply)
	assert.Nil(t, err)
}
