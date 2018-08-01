package osdefs

const (
	// CapAuditControl - Enable and disable kernel auditing; change auditing filter
	// rules; retrieve auditing status and filtering rules.
	CapAuditControl Capability = "CAP_AUDIT_CONTROL"

	// CapAuditRead - Allow reading the audit log via a multicast netlink socket.
	CapAuditRead Capability = "CAP_AUDIT_READ"

	// CapAuditWrite - Write records to kernel auditing log.
	CapAuditWrite Capability = "CAP_AUDIT_WRITE"

	// CapBlockSuspend - Employ features that can block system suspend.
	CapBlockSuspend Capability = "CAP_BLOCK_SUSPEND"

	// CapChown - Make arbitrary changes to file UIDs and GIDs.
	CapChown Capability = "CAP_CHOWN"

	// CapDacOverride - Bypass file read, write, and execute permission checks.
	CapDacOverride Capability = "CAP_DAC_OVERRIDE"

	// CapDacReadSearch -
	// * Bypass file read permission checks and directory read and execute permission checks
	// * invoke open_by_handle_at(2)
	// * use the linkat(2) AT_EMPTY_PATH flag to create a link to a file referred to by a file descriptor
	CapDacReadSearch Capability = "CAP_DAC_READ_SEARCH"

	// CapFowner -
	// * Bypass permission checks on operations that normally require  the filesystem UID of the process
	// 	 to match the UID of the file, excluding the operations covered by CAP_DAC_OVERRIDE and CAP_DAC_READ_SEARCH
	// * set inode flags on arbitrary files
	// * set Access Control Lists (ACLs) on arbitrary files
	// * ignore directory sticky bit on file deletion
	// * specify O_NOATIME for arbitrary files in open and fcntl
	CapFowner Capability = "CAP_FOWNER"

	// CapFsetid -
	// * Don't clear set-user-ID and set-group-ID mode bits when a file is modified
	// * set the set-group-ID bit for a file whose GID does not match  the filesystem or any of the supplementary
	// GIDs of the  calling process.
	CapFsetid Capability = "CAP_FSETID"

	// CapIpcLock - Lock memory
	CapIpcLock Capability = "CAP_IPC_LOCK"

	// CapIpcOwner - Bypass permission checks for operations on System V IPC objects.
	CapIpcOwner Capability = "CAP_IPC_OWNER"

	// CapKill - Bypass permission checks for sending signals.
	CapKill Capability = "CAP_KILL"

	// CapLease - Establish leases on arbitrary files.
	CapLease Capability = "CAP_LEASE"

	// CapLinuxImmutable - Set the FS_APPEND_FL and FS_IMMUTABLE_FL inode flags.
	CapLinuxImmutable Capability = "CAP_LINUX_IMMUTABLE"

	// CapMacAdmin - Override Mandatory Access Control (MAC).  Implemented for the Smack Linux Security Module (LSM).
	CapMacAdmin Capability = "CAP_MAC_ADMIN"

	// CapMacOverride - Allow MAC configuration or state changes.  Implemented for the Smack LSM.
	CapMacOverride Capability = "CAP_MAC_OVERRIDE"

	// CapMknod - Create special files using mknod
	CapMknod Capability = "CAP_MKNOD"

	// CapNetAdmin - Perform various network-related operations:
	// * interface configuration;
	// * administration of IP firewall, masquerading, and accounting;
	// * modify routing tables;
	// * bind to any address for transparent proxying;
	// * set type-of-service (TOS)
	// * clear driver statistics;
	// * set promiscuous mode;
	// * enabling multicasting;
	// * use setsockopt(2) to set the following socket options:  SO_DEBUG, SO_MARK, SO_PRIORITY (for a priority
	// 	 outside the range 0 to 6), SO_RCVBUFFORCE, and SO_SNDBUFFORCE.
	CapNetAdmin Capability = "CAP_NET_ADMIN"

	// CapNetBindService - Bind a socket to Internet domain privileged ports (port numbers less than 1024).
	CapNetBindService Capability = "CAP_NET_BIND_SERVICE"

	// CapNetBroadcast - Make socket broadcasts, and listen to multicasts.
	CapNetBroadcast Capability = "CAP_NET_BROADCAST"

	// CapNetRaw -
	// * Use RAW and PACKET sockets
	// * bind to any address for transparent proxying.
	CapNetRaw Capability = "CAP_NET_RAW"

	// CapSetgid -
	// * Make arbitrary manipulations of process GIDs and supplementary GID list
	// * forge GID when passing socket credentials via UNIX domain sockets
	// * write a group ID mapping in a user namespace
	CapSetgid Capability = "CAP_SETGID"

	// CapSetfcap - Set file capabilities.
	CapSetfcap Capability = "CAP_SETFCAP"

	// CapSetpcap -
	// * If file capabilities are not supported: grant or remove any  capability in the caller's permitted
	// 	   capability set to or from any other process.
	// *  If file capabilities are supported: add any capability from  the calling thread's bounding set to its
	//     inheritable set; drop capabilities from the bounding set; make changes to the securebits flags.
	CapSetpcap Capability = "CAP_SETPCAP"

	// CapSetuid -
	// * Make arbitrary manipulations of process UIDs
	// * forge UID when passing socket credentials via UNIX domain sockets
	// * write a user ID mapping in a user namespace
	CapSetuid Capability = "CAP_SETUID"

	// CapSysAdmin - Perform administrative operations on the system (see man capabilities(7))
	CapSysAdmin Capability = "CAP_SYS_ADMIN"

	// CapSysBoot - Use reboot and kexec_load.
	CapSysBoot Capability = "CAP_SYS_BOOT"

	// CapSysChroot - Use chroot.
	CapSysChroot Capability = "CAP_SYS_CHROOT"

	// CapSysModule - Load and unload kernel modules.
	CapSysModule Capability = "CAP_SYS_MODULE"

	// CapSysNice -
	// * Raise processes nice value
	// * set real-time scheduling policies for processes
	// * set CPU affinity for arbitrary processes
	// * set I/O scheduling class and priority for arbitrary processes
	CapSysNice Capability = "CAP_SYS_NICE"

	// CapSysPacct - Use acct.
	CapSysPacct Capability = "CAP_SYS_PACCT"

	// CapSysPtrace - Trace, inspect and modify the state of arbitrary processes.
	CapSysPtrace Capability = "CAP_SYS_PTRACE"

	// CapSysRawio - Perform various privileged IO operations (see man capabilities(7))
	CapSysRawio Capability = "CAP_SYS_RAWIO"

	// CapSysResource - Perform various privileged resource configuration operations (see man capabilities (7))
	CapSysResource Capability = "CAP_SYS_RESOURCE"

	// CapSysTime - Set system and hardware clocks.
	CapSysTime Capability = "CAP_SYS_TIME"

	// CapSysTtyConfig -
	// * Use vhangupl
	// * Perform various privileged ioctl operations on TTYs
	CapSysTtyConfig Capability = "CAP_SYS_TTY_CONFIG"

	// CapSyslog -
	// * Perform privileged syslog operations
	// * View kernel addresses exposed via /proc under certain conditions
	CapSyslog Capability = "CAP_SYSLOG"

	// CapWakeAlarm - trigger something that will wake up the system
	CapWakeAlarm Capability = "CAP_WAKE_ALARM"
)
