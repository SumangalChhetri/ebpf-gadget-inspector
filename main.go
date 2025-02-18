package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: gadget-inspector <path-to-ebpf-program>")
	}

	// Path to the eBPF ELF file
	ebpfFile := os.Args[1]

	// Allow unlimited locked memory for eBPF
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("failed to remove memlock: %v", err)
	}

	// Load the eBPF program from file
	spec, err := ebpf.LoadCollectionSpec(ebpfFile)
	if err != nil {
		log.Fatalf("failed to load eBPF program: %v", err)
	}

	// Print success message with emojis
	fmt.Println("✅ eBPF program loaded successfully!\n")

	// Print eBPF Programs with formatted details
	fmt.Println("🔍 eBPF Programs:")
	for name, obj := range spec.Programs {
		fmt.Printf("- Program: %s\n  Type: %v\n", name, obj.Type)
	}

	// Print eBPF Maps (if any) with formatted details
	fmt.Println("\n📌 eBPF Maps:")
	for name, obj := range spec.Maps {
		fmt.Printf("- Map: %s\n  Type: %v\n  Key Size: %d bytes\n  Value Size: %d bytes\n",
			name, obj.Type, obj.KeySize, obj.ValueSize)
	}
}

