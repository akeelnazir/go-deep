package shell

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/filesystem"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/memory"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/process"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/scheduler"
)

type Shell struct {
	pm       *process.ProcessManager
	mm       *memory.MemoryManager
	fs       *filesystem.FileSystem
	sched    *scheduler.Scheduler
	running  bool
	commands map[string]CommandFunc
	bootTime time.Time
}

type CommandFunc func(args []string) error

func NewShell(pm *process.ProcessManager, mm *memory.MemoryManager, fs *filesystem.FileSystem, sched *scheduler.Scheduler) *Shell {
	s := &Shell{
		pm:       pm,
		mm:       mm,
		fs:       fs,
		sched:    sched,
		running:  true,
		commands: make(map[string]CommandFunc),
		bootTime: time.Now(),
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
	s.commands["write"] = s.cmdWrite
	s.commands["cpu"] = s.cmdCPU
	s.commands["cpu-info"] = s.cmdCPUInfo
	s.commands["registers"] = s.cmdRegisters
	s.commands["time"] = s.cmdTime
	s.commands["exit"] = s.cmdExit
	s.commands["clear"] = s.cmdClear
}

func (s *Shell) Start() {
	fmt.Println("\n=== OS Emulator Shell ===")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for s.running {
		fmt.Print("emulator> ")
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
	fmt.Println("Process Management:")
	fmt.Println("  ps                - List all processes")
	fmt.Println("  run <name>        - Create and run a process")
	fmt.Println("  kill <pid>        - Terminate process")
	fmt.Println("Memory Management:")
	fmt.Println("  memory            - Show memory statistics")
	fmt.Println("Filesystem:")
	fmt.Println("  ls [path]         - List directory")
	fmt.Println("  mkdir <path>      - Create directory")
	fmt.Println("  touch <path>      - Create file")
	fmt.Println("  rm <path>         - Delete file")
	fmt.Println("  cat <path>        - Display file contents")
	fmt.Println("  write <path> <text> - Write to file")
	fmt.Println("CPU & Execution:")
	fmt.Println("  cpu               - Show CPU statistics")
	fmt.Println("  cpu-info          - Display CPU information")
	fmt.Println("  registers         - Show CPU registers")
	fmt.Println("System:")
	fmt.Println("  info              - Display system information")
	fmt.Println("  time              - Show system time")
	fmt.Println("  clear             - Clear screen")
	fmt.Println("  exit              - Shutdown emulator")
	return nil
}

func (s *Shell) cmdPS(args []string) error {
	procs := s.pm.GetAllProcesses()
	fmt.Println("=== Running Processes ===")
	fmt.Println("PID  Name        State       CPU Time    Memory")
	fmt.Println("---  ----        -----       --------    ------")

	for _, p := range procs {
		p.Mu.RLock()
		state := "New"
		if p.State == process.Ready {
			state = "Ready"
		} else if p.State == process.Running {
			state = "Running"
		} else if p.State == process.Blocked {
			state = "Blocked"
		} else if p.State == process.Terminated {
			state = "Terminated"
		}
		fmt.Printf("%-4d %-11s %-11s %-11v %d KB\n", p.PID, p.Name, state, p.CPUTime, p.MemorySize/1024)
		p.Mu.RUnlock()
	}

	return nil
}

func (s *Shell) cmdRun(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: run <name>")
	}

	p := s.pm.CreateProcess(args[0])
	p.MemorySize = 1024 * 1024
	p.Start()

	fmt.Printf("Process created: PID=%d, Name=%s\n", p.PID, p.Name)
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
	fmt.Println("OS: OS Emulator v1.0")
	fmt.Println("Architecture: x86-64 (emulated)")
	fmt.Printf("Processes: %d\n", len(s.pm.GetAllProcesses()))
	fmt.Printf("Uptime: %v\n", time.Since(s.bootTime))
	return nil
}

func (s *Shell) cmdMemory(args []string) error {
	used, free := s.mm.GetMemoryStats()
	total := used + free
	percentage := (float64(used) / float64(total)) * 100

	fmt.Println("=== Memory Statistics ===")
	fmt.Printf("Total: %d MB\n", total/(1024*1024))
	fmt.Printf("Used:  %d MB\n", used/(1024*1024))
	fmt.Printf("Free:  %d MB\n", free/(1024*1024))
	fmt.Printf("Usage: %.2f%%\n", percentage)
	return nil
}

func (s *Shell) cmdLS(args []string) error {
	path := "/"
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

func (s *Shell) cmdWrite(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: write <path> <content>")
	}

	path := args[0]
	content := strings.Join(args[1:], " ")

	return s.fs.WriteFile(path, []byte(content))
}

func (s *Shell) cmdCPU(args []string) error {
	current := s.sched.GetCurrentProcess()
	if current == nil {
		fmt.Println("No process currently running")
		return nil
	}

	fmt.Printf("Current Process CPU: %s\n", current.CPU.String())
	return nil
}

func (s *Shell) cmdCPUInfo(args []string) error {
	fmt.Println("=== CPU Information ===")
	fmt.Println("Architecture: x86-64 (simulated)")
	fmt.Println("Registers: EAX, EBX, ECX, EDX, ESP, EBP, ESI, EDI")
	fmt.Println("Memory: 256 MB")
	fmt.Println("Page Size: 4096 bytes")
	return nil
}

func (s *Shell) cmdRegisters(args []string) error {
	current := s.sched.GetCurrentProcess()
	if current == nil {
		fmt.Println("No process currently running")
		return nil
	}

	current.CPU.Mu.RLock()
	defer current.CPU.Mu.RUnlock()

	fmt.Println("=== CPU Registers ===")
	for name, value := range current.CPU.Registers {
		fmt.Printf("%s: 0x%08x\n", name, value)
	}
	fmt.Printf("PC: 0x%08x\n", current.CPU.PC)
	fmt.Printf("SP: 0x%08x\n", current.CPU.SP)
	return nil
}

func (s *Shell) cmdTime(args []string) error {
	fmt.Printf("System Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Uptime: %v\n", time.Since(s.bootTime))
	return nil
}

func (s *Shell) cmdClear(args []string) error {
	fmt.Print("\033[2J\033[H")
	return nil
}

func (s *Shell) cmdExit(args []string) error {
	fmt.Println("Shutting down emulator...")
	s.running = false
	return nil
}
