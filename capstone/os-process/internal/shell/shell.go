package shell

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-process/internal/filesystem"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/memory"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/process"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/scheduler"
)

type Shell struct {
	pm         *process.ProcessManager
	mm         *memory.MemoryManager
	fs         *filesystem.FileSystem
	sched      *scheduler.Scheduler
	running    bool
	commands   map[string]CommandFunc
	currentDir string
}

type CommandFunc func(args []string) error

func NewShell(pm *process.ProcessManager, mm *memory.MemoryManager, fs *filesystem.FileSystem, sched *scheduler.Scheduler) *Shell {
	s := &Shell{
		pm:         pm,
		mm:         mm,
		fs:         fs,
		sched:      sched,
		running:    true,
		commands:   make(map[string]CommandFunc),
		currentDir: "/",
	}
	s.registerCommands()
	return s
}

func (s *Shell) registerCommands() {
	s.commands["help"] = s.cmdHelp
	s.commands["ps"] = s.cmdPS
	s.commands["run"] = s.cmdRun
	s.commands["kill"] = s.cmdKill
	s.commands["info"] = s.cmdInfo
	s.commands["memory"] = s.cmdMemory
	s.commands["ls"] = s.cmdLS
	s.commands["mkdir"] = s.cmdMkdir
	s.commands["touch"] = s.cmdTouch
	s.commands["rm"] = s.cmdRM
	s.commands["cat"] = s.cmdCat
	s.commands["echo"] = s.cmdEcho
	s.commands["exit"] = s.cmdExit
	s.commands["clear"] = s.cmdClear
}

func (s *Shell) Start() {
	fmt.Println("\n=== OS Process Shell ===")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for s.running {
		fmt.Print("os-process> ")
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
	fmt.Println("  ps                - List all processes")
	fmt.Println("  run <name> <sec>  - Create and run a process for N seconds")
	fmt.Println("  kill <pid>        - Terminate a process")
	fmt.Println("  info              - Display system information")
	fmt.Println("  memory            - Show memory statistics")
	fmt.Println("  ls [path]         - List directory contents")
	fmt.Println("  mkdir <path>      - Create directory")
	fmt.Println("  touch <path>      - Create file")
	fmt.Println("  rm <path>         - Delete file or directory")
	fmt.Println("  cat <path>        - Display file contents")
	fmt.Println("  echo <text>       - Print text")
	fmt.Println("  clear             - Clear the screen")
	fmt.Println("  exit              - Shutdown the OS")
	return nil
}

func (s *Shell) cmdPS(args []string) error {
	procs := s.pm.GetAllProcesses()
	fmt.Println("=== Running Processes ===")
	fmt.Println("PID  Name        State       CPU Time    Memory")
	fmt.Println("---  ----        -----       --------    ------")

	for _, p := range procs {
		p.Mu.RLock()
		state := "Ready"
		if p.State == process.Running {
			state = "Running"
		} else if p.State == process.Blocked {
			state = "Blocked"
		} else if p.State == process.Terminated {
			state = "Terminated"
		}
		fmt.Printf("%-4d %-11s %-11s %-11v %d KB\n", p.PID, p.Name, state, p.CPUTime, p.MemoryUsed/1024)
		p.Mu.RUnlock()
	}

	return nil
}

func (s *Shell) cmdRun(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: run <name> <duration_seconds>")
	}

	name := args[0]
	seconds, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	p := s.pm.CreateProcess(name, time.Duration(seconds)*time.Second)
	p.Start()

	addr, err := s.mm.Allocate(p.PID, 1024*1024)
	if err != nil {
		return fmt.Errorf("failed to allocate memory: %v", err)
	}

	p.MemoryUsed = 1024 * 1024
	fmt.Printf("Process created: PID=%d, Name=%s, Memory=0x%x\n", p.PID, name, addr)

	return nil
}

func (s *Shell) cmdKill(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: kill <pid>")
	}

	pid, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid PID: %v", err)
	}

	return s.pm.TerminateProcess(uint32(pid))
}

func (s *Shell) cmdInfo(args []string) error {
	fmt.Println("=== System Information ===")
	fmt.Println("OS: OS Process v1.0")
	fmt.Println("Architecture: x86-64 (simulated)")
	fmt.Printf("Processes: %d\n", len(s.pm.GetAllProcesses()))
	return nil
}

func (s *Shell) cmdMemory(args []string) error {
	total, used, free := s.mm.GetMemoryStats()
	percentage := (float64(used) / float64(total)) * 100

	fmt.Println("=== Memory Statistics ===")
	fmt.Printf("Total: %d MB\n", total/(1024*1024))
	fmt.Printf("Used:  %d MB\n", used/(1024*1024))
	fmt.Printf("Free:  %d MB\n", free/(1024*1024))
	fmt.Printf("Usage: %.2f%%\n", percentage)
	return nil
}

func (s *Shell) cmdLS(args []string) error {
	path := s.currentDir
	if len(args) > 0 {
		path = args[0]
	}

	entries, err := s.fs.ListDirectory(path)
	if err != nil {
		return err
	}

	fmt.Printf("Contents of %s:\n", path)
	for _, entry := range entries {
		fmt.Printf("  %s\n", entry)
	}
	return nil
}

func (s *Shell) cmdMkdir(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: mkdir <path>")
	}
	return s.fs.CreateDirectory(args[0])
}

func (s *Shell) cmdTouch(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: touch <path>")
	}
	return s.fs.CreateFile(args[0], []byte{})
}

func (s *Shell) cmdRM(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: rm <path>")
	}
	return s.fs.DeleteFile(args[0])
}

func (s *Shell) cmdCat(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: cat <path>")
	}

	content, err := s.fs.ReadFile(args[0])
	if err != nil {
		return err
	}

	fmt.Println(string(content))
	return nil
}

func (s *Shell) cmdEcho(args []string) error {
	fmt.Println(strings.Join(args, " "))
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
