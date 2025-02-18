# eBPF Gadget Inspector

eBPF Gadget Inspector is a tool designed to load eBPF programs (ELF files), inspect them, and display detailed information about their programs and maps. It leverages the Cilium eBPF Go library to provide insights such as program type, map type, key sizes, and value sizes.

---

## Project Overview

This project allows you to verify and inspect eBPF programs loaded into the kernel. It prints out information about the eBPF programs (e.g., tracepoints, XDP) and any associated maps (e.g., Hash maps) that store runtime data. This tool is useful for developers working with eBPF and kernel instrumentation.

---

## 📦 Project Structure

```

ebpf-gadget-inspector/
├── Makefile               # Build configuration for process_monitor.elf
├── README.md              # This complete file
├── gadget-inspector       # Go binary that loads and inspects eBPF programs (generated)
├── go.mod                 # Go module definition
├── go.sum                 # Go checksum file
├── main.go                # Main Go file containing logic for loading and inspecting eBPF programs
└── process_monitor.c      # C source file for a process-monitoring eBPF program

```

## 🔧 Setup & Installation

Ensure you have the following installed on your system:

### 1. Go (version 1.16 or higher)
```
Download from: [https://golang.org/dl/](https://golang.org/dl/)
```
### 2. Clang with BPF target support
Install with:

```bash
sudo apt-get install clang llvm
```

## 3.Linux Headers
Install with: 
```bash
sudo apt-get install linux-headers-$(uname -r)
```

## Cilium eBPF Go Library
Install with: 
```bash
go get github.com/cilium/ebpf
```
## Setting Up the Project
Clone the Repository
```bash
git clone https://github.com/SumangalChhetri/ebpf-gadget-inspector.git
cd ebpf-gadget-inspector
```

## Build the eBPF Program
Compile the C eBPF program:

```bash
clang -O2 -g -target bpf -c process_monitor.c -o process_monitor.elf
```

## Build the Go Binary
Compile the Gadget Inspector tool:
```bash
go build -o gadget-inspector main.go
```

## Install Go Dependencies
Ensure all dependencies are installed:

```bash
go mod tidy
```

## (Optional) Use the Makefile
Build the eBPF program using the provided Makefile:

```bash
make
```
## Usage
To run the gadget-inspector tool and inspect an eBPF ELF file, use the following command:
```bash
sudo ./gadget-inspector ./process_monitor.elf
```

## Expected Output
When the tool runs successfully, the output should look like this:

```bash
✅ eBPF program loaded successfully!

🔍 eBPF Programs:
- Program: trace_execve
  Type: TracePoint

📌 eBPF Maps:
- Map: process_map
  Type: Hash
  Key Size: 4 bytes
  Value Size: 16 bytes
```

## Troubleshooting
Clang & Linux Headers
Ensure you have the correct versions installed:

```bash
sudo apt-get install linux-headers-$(uname -r)
```

## File Paths
Verify that the correct path to process_monitor.elf is provided when running the tool.

## Permissions
Run the tool with sudo since loading eBPF programs requires elevated privileges.

## Contribution

1.Contributions are welcome! If you'd like to contribute:

2.Fork this repository.

3.Make your changes.

4.Submit a pull request describing your improvements or bug fixes.

## 📜 License

This project is licensed under the MIT License. See the LICENSE file for details.

---
