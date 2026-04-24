package shell

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akeelnazir/go-deep/capstone/os-full/internal/device"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/filesystem"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/ipc"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/memory"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/process"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/scheduler"
)

type Shell struct {
	pm       *process.ProcessManager
	mm       *memory.MemoryManager
	fs       *filesystem.FileSystem
	dm       *device.DeviceManager
	im       *ipc.IPCManager
	sched    *scheduler.Scheduler
	running  bool
	commands map[string]CommandFunc
}

type CommandFunc func(args []string) error

func NewShell(pm *process.ProcessManager, mm *memory.MemoryManager, fs *filesystem.FileSystem,
	dm *device.DeviceManager, im *ipc.IPCManager, sched *scheduler.Scheduler) *Shell {
	s := &Shell{
		pm:       pm,
		mm:       mm,
		fs:       fs,
		dm:       dm,
		im:       im,
		sched:    sched,
		running:  true,
		commands: make(map[string]CommandFunc),
	}
	s.registerCommands()
	return s
}

func (s *Shell) registerCommands() {
	s.commands["help"] = s.cmdHelp
	s.commands["ps"] = s.cmdPS
	s.commands["run"] = s.cmdRun
	s.commands["kill"] = s.cmdKill
	s.commands["nice"] = s.cmdNice
	s.commands["info"] = s.cmdInfo
	s.commands["memory"] = s.cmdMemory
	s.commands["pages"] = s.cmdPages
	s.commands["ls"] = s.cmdLS
	s.commands["mkdir"] = s.cmdMkdir
	s.commands["touch"] = s.cmdTouch
	s.commands["rm"] = s.cmdRM
	s.commands["cat"] = s.cmdCat
	s.commands["chmod"] = s.cmdChmod
	s.commands["chown"] = s.cmdChown
	s.commands["devices"] = s.cmdDevices
	s.commands["mount"] = s.cmdMount
	s.commands["exit"] = s.cmdExit
	s.commands["clear"] = s.cmdClear
}

func (s *Shell) Start() {
	fmt.Println("\n=== OS Full Shell ===")
	fmt.Println("Type 'help' for available commands")
	fmt.Println()

	reader := bufio.NewReader(nil)

	for s.running {
		fmt.Print("os-full> ")
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
	fmt.Println("  ps                    - List all processes")
	fmt.Println("  run <name> <pri> <sec> - Create process with priority")
	fmt.Println("  kill <pid>            - Terminate process")
	fmt.Println("  nice <pid> <priority> - Change process priority")
	fmt.Println("Memory Management:")
	fmt.Println("  memory                - Show memory statistics")
	fmt.Println("  pages                 - Show page table information")
	fmt.Println("Filesystem:")
	fmt.Println("  ls [path]             - List directory")
	fmt.Println("  mkdir <path>          - Create directory")
	fmt.Println("  touch <path>          - Create file")
	fmt.Println("  rm <path>             - Delete file")
	fmt.Println("  cat <path>            - Display file contents")
	fmt.Println("  chmod <mode> <path>   - Change permissions")
	fmt.Println("  chown <user> <path>   - Change owner")
	fmt.Println("Devices:")
	fmt.Println("  devices               - List devices")
	fmt.Println("  mount <dev> <path>    - Mount device")
	fmt.Println("System:")
	fmt.Println("  info                  - Display system information")
	fmt.Println("  clear                 - Clear screen")
	fmt.Println("  exit                  - Shutdown OS")
	return nil
}

func (s *Shell) cmdPS(args []string) error {
	procs := s.pm.GetAllProcesses()
	fmt.Println("=== Running Processes ===")
	fmt.Println("PID  PPID Name        State       Pri CPU Time    Memory")
	fmt.Println("---  ---- ----        -----       --- --------    ------")

	for _, p := range procs {
		p.Mu.RLock()
		state := "Ready"
		if p.State == process.Running {
			state = "Running"
		} else if p.State == process.Terminated {
			state = "Terminated"
		}
		fmt.Printf("%-4d %-4d %-11s %-11s %-3d %-11v %d KB\n",
			p.PID, p.PPID, p.Name, state, p.Priority, p.CPUTime, p.MemoryUsed/1024)
		p.Mu.RUnlock()
	}

	return nil
}

func (s *Shell) cmdRun(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("usage: run <name> <priority> <duration_seconds>")
	}

	priority, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid priority: %v", err)
	}

	if priority < 0 || priority > 15 {
		return fmt.Errorf("priority must be 0-15")
	}

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

func (s *Shell) cmdNice(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: nice <pid> <priority>")
	}

	pid, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid PID: %v", err)
	}

	priority, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid priority: %v", err)
	}

	return s.pm.ChangePriority(uint32(pid), priority)
}

func (s *Shell) cmdInfo(args []string) error {
	fmt.Println("=== System Information ===")
	fmt.Println("OS: OS Full v1.0")
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

func (s *Shell) cmdPages(args []string) error {
	total, used := s.mm.GetPageCount()
	fmt.Println("=== Page Table Information ===")
	fmt.Printf("Total Pages: %d\n", total)
	fmt.Printf("Used Pages:  %d\n", used)
	fmt.Printf("Free Pages:  %d\n", total-used)
	fmt.Printf("Page Size:   %d bytes\n", memory.PageSize)
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

func (s *Shell) cmdChmod(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: chmod <mode> <path>")
	}

	mode, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid mode: %v", err)
	}

	perms := filesystem.Permissions{
		Owner: (mode / 100) % 10,
		Group: (mode / 10) % 10,
		Other: mode % 10,
	}

	return s.fs.ChangePermissions(args[1], perms)
}

func (s *Shell) cmdChown(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: chown <user> <path>")
	}

	return s.fs.ChangeOwner(args[1], args[0])
}

func (s *Shell) cmdDevices(args []string) error {
	devices := s.dm.ListDevices()
	fmt.Println("=== Registered Devices ===")
	for _, d := range devices {
		fmt.Println(d.String())
	}
	return nil
}

func (s *Shell) cmdMount(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: mount <device> <path>")
	}

	return s.dm.MountDevice(args[0], args[1])
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
