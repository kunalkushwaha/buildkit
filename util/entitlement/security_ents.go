package entitlement

import (
	"github.com/moby/buildkit/util/entitlement/osdefs"
	"github.com/opencontainers/runtime-spec/specs-go"
)

/* Implements "security.confined" entitlement:
 * - Blocked caps: CAP_SYS_ADMIN, CAP_SYS_PTRACE, CAP_SETUID, CAP_SETGID, CAP_SETPCAP, CAP_SETFCAP, CAP_MAC_ADMIN,
 *					CAP_MAC_OVERRIDE, CAP_DAC_OVERRIDE, CAP_DAC_READ_SEARCH, CAP_FSETID, CAP_SYS_MODULE, CAP_SYSLOG,
 * 					CAP_SYS_RAWIO, CAP_LINUX_IMMUTABLE, CAP_SYS_RESOURCE
 * - Blocked syscalls: ptrace, arch_prctl, personality, madvise, prctl with PR_CAPBSET_DROP and PR_CAPBSET_READ
 */
func securityConfinedEntitlementEnforce(profile Profile) (Profile, error) {
	ociProfile, err := ociProfileConversionCheck(profile, "security.confined")
	if err != nil {
		return nil, err
	}

	capsToRemove := []osdefs.Capability{
		osdefs.CapMacAdmin, osdefs.CapMacOverride, osdefs.CapDacOverride, osdefs.CapDacReadSearch, osdefs.CapSetpcap, osdefs.CapSetfcap, osdefs.CapSetuid, osdefs.CapSetgid,
		osdefs.CapSysPtrace, osdefs.CapFsetid, osdefs.CapSysModule, osdefs.CapSyslog, osdefs.CapSysRawio, osdefs.CapSysAdmin, osdefs.CapLinuxImmutable,
		osdefs.CapSysResource,
	}
	ociProfile.RemoveCaps(capsToRemove...)

	syscallsToBlock := []osdefs.Syscall{
		osdefs.SysPtrace, osdefs.SysArchPrctl, osdefs.SysPersonality, osdefs.SysMadvise,
	}
	ociProfile.BlockSyscalls(syscallsToBlock...)

	syscallsWithArgsToAllow := map[osdefs.Syscall][]specs.LinuxSeccompArg{
		osdefs.SysPrctl: {
			{
				Index: 0,
				Value: osdefs.PrCapbsetDrop,
				Op:    specs.OpNotEqual,
			},
			{
				Index: 0,
				Value: osdefs.PrCapbsetRead,
				Op:    specs.OpNotEqual,
			},
		},
	}
	ociProfile.AllowSyscallsWithArgs(syscallsWithArgsToAllow)

	/* FIXME: Add AppArmor rules to deny RW on sensitive FS directories */

	return ociProfile, nil
}

/* Implements "security.unconfied" entitlement:
 * - Authorized caps: CAP_MAC_ADMIN, CAP_MAC_OVERRIDE, CAP_DAC_OVERRIDE, CAP_DAC_READ_SEARCH, CAP_SETPCAP, CAP_SETFCAP,
 * 						CAP_SETUID, CAP_SETGID, CAP_SYS_PTRACE, CAP_FSETID, CAP_SYS_MODULE, CAP_SYSLOG, CAP_SYS_RAWIO,
 *						CAP_SYS_ADMIN, CAP_LINUX_IMMUTABLE, CAP_SYS_BOOT, CAP_SYS_NICE, CAP_SYS_PACCT,
 *						CAP_SYS_TTY_CONFIG, CAP_SYS_TIME, CAP_WAKE_ALARM, CAP_AUDIT_READ, CAP_AUDIT_WRITE,
 *						CAP_AUDIT_CONTROL,
 * 						CAP_SYS_RESOURCE
 * - Allowed syscalls: ptrace, arch_prctl, personality, setuid, setgid, prctl, madvise, mount, init_module,
 *						finit_module, setns, clone, unshare
 * - No read-only paths
 */
func securityUnconfinedEntitlementEnforce(profile Profile) (Profile, error) {
	ociProfile, err := ociProfileConversionCheck(profile, "security.unconfined")
	if err != nil {
		return nil, err
	}

	capsToAdd := []osdefs.Capability{
		osdefs.CapMacAdmin, osdefs.CapMacOverride, osdefs.CapDacOverride, osdefs.CapDacReadSearch, osdefs.CapSetpcap, osdefs.CapSetfcap, osdefs.CapSetuid, osdefs.CapSetgid,
		osdefs.CapSysPtrace, osdefs.CapFsetid, osdefs.CapSysModule, osdefs.CapSyslog, osdefs.CapSysRawio, osdefs.CapSysAdmin, osdefs.CapLinuxImmutable, osdefs.CapSysBoot,
		osdefs.CapSysNice, osdefs.CapSysPacct, osdefs.CapSysTtyConfig, osdefs.CapSysTime, osdefs.CapWakeAlarm, osdefs.CapAuditRead, osdefs.CapAuditWrite, osdefs.CapAuditControl,
		// FIXME: osdefs.CapSysResource should probably part of a limit_resource entitlement..
		osdefs.CapSysResource,
	}
	ociProfile.AddCaps(capsToAdd...)

	syscallsToAllow := []osdefs.Syscall{
		osdefs.SysPtrace, osdefs.SysArchPrctl, osdefs.SysPersonality, osdefs.SysSetuid, osdefs.SysSetgid, osdefs.SysPrctl, osdefs.SysMadvise, osdefs.SysMount, osdefs.SysUmount2,
		osdefs.SysInitModule, osdefs.SysFinitModule, osdefs.SysSetns, osdefs.SysClone, osdefs.SysUnshare, osdefs.SysKeyctl, osdefs.SysPivotRoot,
		osdefs.SysSethostname,
		osdefs.SysSetdomainname,
		osdefs.SysIopl,
		osdefs.SysIoperm,
		osdefs.SysCreateModule,
		osdefs.SysInitModule,
		osdefs.SysDeleteModule,
		osdefs.SysGetKernelSyms,
		osdefs.SysQueryModule,
		osdefs.SysQuotactl,
		osdefs.SysGetpmsg,
		osdefs.SysPutpmsg,
	}
	ociProfile.AllowSyscalls(syscallsToAllow...)

	// Just in case some default configuration does add read-only paths, we remove them
	ociProfile.OCI.Linux.ReadonlyPaths = []string{}

	return ociProfile, nil
}
