package ipc

import (
	"fmt"
	"sync"
)

type Signal int

const (
	SIGTERM Signal = 15
	SIGKILL Signal = 9
	SIGUSR1 Signal = 10
	SIGUSR2 Signal = 12
)

type Pipe struct {
	Name   string
	Buffer []byte
	Mu     sync.RWMutex
}

type MessageQueue struct {
	Name     string
	Messages []string
	Mu       sync.RWMutex
}

type IPCManager struct {
	pipes        map[string]*Pipe
	queues       map[string]*MessageQueue
	signalHandlers map[uint32]map[Signal]func()
	Mu           sync.RWMutex
}

func NewIPCManager() *IPCManager {
	return &IPCManager{
		pipes:        make(map[string]*Pipe),
		queues:       make(map[string]*MessageQueue),
		signalHandlers: make(map[uint32]map[Signal]func()),
	}
}

func (im *IPCManager) CreatePipe(name string) (*Pipe, error) {
	im.Mu.Lock()
	defer im.Mu.Unlock()

	if _, exists := im.pipes[name]; exists {
		return nil, fmt.Errorf("pipe already exists: %s", name)
	}

	pipe := &Pipe{
		Name:   name,
		Buffer: make([]byte, 0),
	}
	im.pipes[name] = pipe
	return pipe, nil
}

func (im *IPCManager) WriteToPipe(name string, data []byte) error {
	im.Mu.RLock()
	defer im.Mu.RUnlock()

	pipe, exists := im.pipes[name]
	if !exists {
		return fmt.Errorf("pipe not found: %s", name)
	}

	pipe.Mu.Lock()
	defer pipe.Mu.Unlock()

	pipe.Buffer = append(pipe.Buffer, data...)
	return nil
}

func (im *IPCManager) ReadFromPipe(name string) ([]byte, error) {
	im.Mu.RLock()
	defer im.Mu.RUnlock()

	pipe, exists := im.pipes[name]
	if !exists {
		return nil, fmt.Errorf("pipe not found: %s", name)
	}

	pipe.Mu.Lock()
	defer pipe.Mu.Unlock()

	data := make([]byte, len(pipe.Buffer))
	copy(data, pipe.Buffer)
	pipe.Buffer = make([]byte, 0)
	return data, nil
}

func (im *IPCManager) CreateQueue(name string) (*MessageQueue, error) {
	im.Mu.Lock()
	defer im.Mu.Unlock()

	if _, exists := im.queues[name]; exists {
		return nil, fmt.Errorf("queue already exists: %s", name)
	}

	queue := &MessageQueue{
		Name:     name,
		Messages: make([]string, 0),
	}
	im.queues[name] = queue
	return queue, nil
}

func (im *IPCManager) SendMessage(queueName string, message string) error {
	im.Mu.RLock()
	defer im.Mu.RUnlock()

	queue, exists := im.queues[queueName]
	if !exists {
		return fmt.Errorf("queue not found: %s", queueName)
	}

	queue.Mu.Lock()
	defer queue.Mu.Unlock()

	queue.Messages = append(queue.Messages, message)
	return nil
}

func (im *IPCManager) ReceiveMessage(queueName string) (string, error) {
	im.Mu.RLock()
	defer im.Mu.RUnlock()

	queue, exists := im.queues[queueName]
	if !exists {
		return "", fmt.Errorf("queue not found: %s", queueName)
	}

	queue.Mu.Lock()
	defer queue.Mu.Unlock()

	if len(queue.Messages) == 0 {
		return "", fmt.Errorf("queue is empty: %s", queueName)
	}

	message := queue.Messages[0]
	queue.Messages = queue.Messages[1:]
	return message, nil
}

func (im *IPCManager) RegisterSignalHandler(pid uint32, signal Signal, handler func()) {
	im.Mu.Lock()
	defer im.Mu.Unlock()

	if _, exists := im.signalHandlers[pid]; !exists {
		im.signalHandlers[pid] = make(map[Signal]func())
	}

	im.signalHandlers[pid][signal] = handler
}

func (im *IPCManager) SendSignal(pid uint32, signal Signal) error {
	im.Mu.RLock()
	defer im.Mu.RUnlock()

	handlers, exists := im.signalHandlers[pid]
	if !exists {
		return fmt.Errorf("no signal handlers for process %d", pid)
	}

	if handler, exists := handlers[signal]; exists {
		handler()
		return nil
	}

	return fmt.Errorf("no handler for signal %d", signal)
}
