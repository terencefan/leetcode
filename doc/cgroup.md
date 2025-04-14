CFS 使用 rbtree，原因是相比 heap 常数因子更小

每个 unit 1000 connections, 1G memory，1024 shares，0.5 cpu core

/sys/fs/cgroup

td-agent 1c 1G
shared-controller 1c 0.5G
mdsd 1c 2G
mdsd-customer 1c 2G
signalr 8c 14G 40K max upgraded connections

COMPlus_GCHeapCount
COMPlus_GCHighMemPercent 90% -> 80%
COMPlus_GCGen0MaxBudget 4MB -> 16MB

### CGroup V1 (tmpfs)
cfs_quota_us
cfs_period_us

cpu.shares [2, 262142]
cpuacct.usage

shares [2, 262142]

### CGroup v2 (cgroup2fs)
cpu.max
cpu.weight [1, 10000]
cpu.stat

```bash
cgexec -g cpu,memory:{cgroup} {executable} {args}
```

1. create new subtree
2. move current pids to subtree
3. enable cpu / memory controller

modify coreclr/cgroup.cpp to override current cgroup

controllers (available / non-leaf) / cgroup.subtree_control (actual / leaf)

1. Find root and path from `/proc/self/mountinfo`

   ```bash
   3360 3357 0:31 /kubepods/burstable/pod3b14c943-f7a6-4486-ab9f-aa2cfb8353ed/bdeedcb94010c5fa7462f7c046d31367090e3806b5faad1dc1b7dabfe3555ef4 /sys/fs/cgroup/cpu,cpuacct ro,nosuid,nodev,noexec,relatime master:7 - cgroup cgroup rw,cpu,cpuacct
   ```

   hierarchyRoot = `/kubepods/burstable/pod3b14c943-f7a6-4486-ab9f-aa2cfb8353ed/bdeedcb94010c5fa7462f7c046d31367090e3806b5faad1dc1b7dabfe3555ef4`
   hierarchyMount = `/sys/fs/cgroup/cpu`

2. Find cgroup info from `/proc/self/cgroup`

   ```bash
   3:cpu,cpuacct:/kubepods/burstable/pod3b14c943-f7a6-4486-ab9f-aa2cfb8353ed/bdeedcb94010c5fa7462f7c046d31367090e3806b5faad1dc1b7dabfe3555ef4
   ```

   cgroupPathRelativeToMount = `/kubepods/burstable/pod3b14c943-f7a6-4486-ab9f-aa2cfb8353ed/bdeedcb94010c5fa7462f7c046d31367090e3806b5faad1dc1b7dabfe3555ef4`

3. If host cgroup path `/host/sys/fs/cgroup` exists

   containerId = `bdeedcb94010c5fa7462f7c046d31367090e3806b5faad1dc1b7dabfe3555ef4`

   hierarchyRootExcludeContainerId = `/kubepods/burstable/pod3b14c943-f7a6-4486-ab9f-aa2cfb8353ed`

   `/host/sys/fs/cgroup/cpu`

   path = `/host/sys/fs/cgroup/cpu`
