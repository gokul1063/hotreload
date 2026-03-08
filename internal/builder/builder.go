package builder

import (
	"os"
	"os/exec"

	"hotreload/internal/logging"
)

type Builder struct {
	cmd string
}

func New(buildCmd string) *Builder {

	return &Builder{
		cmd: buildCmd,
	}
}

func (b *Builder) Build() error {

	logging.LogWorkflow("builder", "Build", "started")

	cmd := exec.Command("sh", "-c", b.cmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {

		logging.LogError("builder", "Build", err)
		return err
	}

	logging.LogWorkflow("builder", "Build", "success")

	return nil
}
