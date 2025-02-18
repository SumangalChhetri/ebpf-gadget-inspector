// process_monitor.c
#include <linux/bpf.h>
#include <linux/ptrace.h>
#include <linux/types.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

// Define the BPF map to store process information
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u32);      // Use __u32 instead of u32
    __type(value, char[16]);
} process_map SEC(".maps");

// BPF program to monitor execve system calls
SEC("tracepoint/syscalls/sys_enter_execve")
int trace_execve(struct pt_regs *ctx) {
    __u32 pid = bpf_get_current_pid_tgid() >> 32;
    char comm[16];
    bpf_get_current_comm(comm, sizeof(comm));

    bpf_map_update_elem(&process_map, &pid, comm, BPF_ANY);
    bpf_printk("Process executed: PID=%d, Command=%s\n", pid, comm);
    
    return 0;
}

char _license[] SEC("license") = "GPL";

