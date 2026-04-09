package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-minimal/internal/kernel"
)

type Shell struct {
	kernel    *kernel.Kernel
	running   bool
	commands  map[string]CommandFunc
	bootTime  time.Time
}

type CommandFunc func(args []string) error

func NewShell(k *kernel.Kernel) *Shell {
	s := &Shell{
		kernel:   k,
		running:  true,
		commands: make(map[string]CommandFunc),
		bootTime: time.Now(),
	}
	s.registerCommands()
	return s
}

func (s *Shell) registerCommands() {
	s.commands["help"] = s.cmdHelp
	s.commands["echo"] = s.cmdEcho
	s.commands["info"] = s.cmdInfo
	s.commands["memory"] = s.cmdMemory
	s.commands["ps"] = s.cmdPS
	s.commands["exit"] = s.cmdExit
	s.commands["clear"] = s.cmdClear
}

func (s *Shell) Start() {
	fmt.Println("\n=== OS Minimal Shell ===")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for s.running {
		fmt.Print("kernel> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		cmd := parts[0]
		args := parts[1:]

		if handler, exists := s.commands[cmd]; exists {
			if err := handler(args); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", cmd)
		}
	}
}

func (s *Shell) cmdHelp(args []string) error {
	fmt.Println("Available commands:")
	fmt.Println("  help              - Display this help message")
	fmt.Println("  echo <text>       - Print text to console")
	fmt.Println("  info              - Display system information")
	fmt.Println("  memory            - Show memory statistics")
	fmt.Println("  ps                - List running processes")
	fmt.Println("  clear             - Clear the screen")
	fmt.Println("  exit              - Shutdown the OS")
	return nil
}

func (s *Shell) cmdEcho(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func (s *Shell) cmdInfo(args []string) error {
	fmt.Println("=== System Information ===")
	fmt.Println("OS: OS Minimal v1.0")
	fmt.Println("Architecture: x86-64 (simulated)")
	fmt.Printf("Uptime: %v\n", s.kernel.GetUptime())
	fmt.Printf("Boot time: %v\n", s.bootTime)
	return nil
}

func (s *Shell) cmdMemory(args []string) error {
	total, used := s.kernel.GetMemoryStats()
	free := total - used
	percentage := (float64(used) / float64(total)) * 100

	fmt.Println("=== Memory Statistics ===")
	fmt.Printf("Total: %d bytes (%d MB)\n", total, total/(1024*1024))
	fmt.Printf("Used:  %d bytes (%d MB)\n", used, used/(1024*1024))
	fmt.Printf("Free:  %d bytes (%d MB)\n", free, free/(1024*1024))
	fmt.Printf("Usage: %.2f%%\n", percentage)
	return nil
}

func (s *Shell) cmdPS(args []string) error {
	fmt.Println("=== Running Processes ===")
	fmt.Println("PID  Name")
	fmt.Println("1    kernel")
	fmt.Println("2    shell")
	return nil
}

func (s *Shell) cmdClear(args []string) error {
	fmt.Print("\033[2J\033[H")
	return nil
}

func (s *Shell) cmdExit(args []string) error {
	fmt.Println("Shutting down OS...")
	s.running = false
	return nil
}
