package main

import (
	"os/exec"
)

type Service string

func (s Service) Start() error {
	_, err := exec.Command("launchctl", "load", string(s)).Output()
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Stop() error {
	_, err := exec.Command("launchctl", "unload", string(s)).Output()
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Restart() error {
	err := s.Stop()
	if err != nil {
		return err
	}

	return s.Start()
}
