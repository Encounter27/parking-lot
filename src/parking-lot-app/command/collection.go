package command

import (
	"fmt"
	"sync"
)

var once sync.Once
var lock = &sync.Mutex{}

type collection map[string]ICommand

var allCommands collection

func CreateCollection() collection {

	once.Do(func() { // atomic, does not allow repeating
		allCommands = make(collection) // thread safe
	})

	return allCommands
}

func (all collection) RegisterCommand(name string, cmd ICommand) error {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := all[name]; ok {
		return fmt.Errorf("Command %s already registered", name)
	}

	all[name] = cmd
	return nil
}

func (all collection) Get(name string) (ICommand, error) {
	lock.Lock()
	defer lock.Unlock()

	var cmd ICommand
	var ok bool
	if cmd, ok = all[name]; !ok {
		return nil, fmt.Errorf("Command %s not registered", name)
	}

	return cmd, nil
}
