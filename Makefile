BPF_CFLAGS = -O2 -g -Wall -target bpf
GO_CFLAGS = 

all: process_monitor.elf gadget-inspector

process_monitor.elf: process_monitor.c
	clang $(BPF_CFLAGS) -c $< -o $@

gadget-inspector: main.go
	go build -o gadget-inspector main.go

clean:
	rm -f process_monitor.elf gadget-inspector
