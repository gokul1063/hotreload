package runner

import (
	"os"
	"os/exec"
	"syscall"

	"hotreload/internal/logging"
)

type Runner struct {
	cmdStr string
	cmd    *exec.Cmd
}

func New(execCmd string) *Runner {

	return &Runner{
		cmdStr: execCmd,
	}
}

func (r *Runner) Start() error {

	logging.LogWorkflow("runner", "Start", "started")

	cmd := exec.Command("sh", "-c", r.cmdStr)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	err := cmd.Start()
	if err != nil {

		logging.LogError("runner", "Start", err)
		return err
	}

	r.cmd = cmd

	logging.LogWorkflow("runner", "Start", "success")

	return nil
}

func (r *Runner) Stop() error {

	if r.cmd == nil || r.cmd.Process == nil {
		return nil
	}

	logging.LogWorkflow("runner", "Stop", "started")

	pgid, err := syscall.Getpgid(r.cmd.Process.Pid)
	if err != nil {

		logging.LogError("runner", "Getpgid", err)
		return err
	}

	err = syscall.Kill(-pgid, syscall.SIGKILL)
	if err != nil {

		logging.LogError("runner", "KillGroup", err)
		return err
	}

	_, _ = r.cmd.Process.Wait()

	logging.LogWorkflow("runner", "Stop", "success")

	r.cmd = nil

	return nil
}
