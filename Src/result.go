package main

import "errors"

type Result struct {
	Id int
	PlayerName string
	Result int
	Time int
}

func (r Result) Validate() (error){
	if r.Result <= 0 {
		return errors.New("Result is required")
	}
	if r.Time <= 0 {
		return errors.New("Time is required")
	}
	if r.PlayerName == "" {
		return errors.New("Player name is required")
	}

	return nil
}