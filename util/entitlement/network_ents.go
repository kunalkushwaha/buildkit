package entitlement

import (
	"syscall"

	"github.com/moby/buildkit/util/entitlement/osdefs"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	networkDomain = "network"
)

/* Implements "network.none" entitlement
 * - No access to /proc/pid/net, /proc/sys/net, /sys/class/net
 * - No caps: CAP_NET_ADMIN, CAP_NET_BIND_SERVICE, CAP_NET_RAW, CAP_NET_BROADCAST
 * - Blocked syscalls:
 *     socket, socketpair, setsockopt, getsockopt, getsockname, getpeername, bind, listen, accept,
 *     accept4, connect, shutdown,recvfrom, recvmsg, recvmmsg, sendto, sendmsg, sendmmsg, sethostname,
 *     setdomainname, socket for non AF_LOCAL/AF_UNIX domain
 * - Enable network namespacing
 */
func networkNoneEntitlementEnforce(profile Profile) (Profile, error) {
	ociProfile, err := ociProfileConversionCheck(profile, "network.none")
	if err != nil {
		return nil, err
	}

	capsToRemove := []osdefs.Capability{osdefs.CapNetAdmin, osdefs.CapNetBindService, osdefs.CapNetRaw, osdefs.CapNetBroadcast}
	ociProfile.RemoveCaps(capsToRemove...)

	pathsToMask := []string{"/proc/pid/net", "/proc/sys/net", "/sys/class/net"}
	ociProfile.AddMaskedPaths(pathsToMask...)

	nsToAdd := []specs.LinuxNamespaceType{specs.NetworkNamespace}
	ociProfile.AddNamespaces(nsToAdd...)

	syscallsToBlock := []osdefs.Syscall{osdefs.SysSocket, osdefs.SysSocketpair, osdefs.SysSetsockopt, osdefs.SysGetsockopt, osdefs.SysGetsockname, osdefs.SysGetpeername,
		osdefs.SysBind, osdefs.SysListen, osdefs.SysAccept, osdefs.SysAccept4, osdefs.SysConnect, osdefs.SysShutdown, osdefs.SysRecvfrom, osdefs.SysRecvmsg, osdefs.SysRecvmmsg, osdefs.SysSendto,
		osdefs.SysSendmsg, osdefs.SysSendmmsg, osdefs.SysSethostname, osdefs.SysSetdomainname,
	}
	ociProfile.BlockSyscalls(syscallsToBlock...)

	syscallsWithArgsToAllow := map[osdefs.Syscall][]specs.LinuxSeccompArg{
		osdefs.SysSocket: {
			{
				Index: 0,
				Op:    specs.OpEqualTo,
				Value: syscall.AF_UNIX,
			},
			{
				Index: 0,
				Op:    specs.OpEqualTo,
				Value: syscall.AF_LOCAL,
			},
		},
	}
	ociProfile.AllowSyscallsWithArgs(syscallsWithArgsToAllow)

	// FIXME: build an Apparmor Profile if necessary + add `deny network`

	return profile, nil
}

/* Implements "network.admin" entitlement
 * - Authorized caps: CAP_NET_ADMIN, CAP_NET_BROADCAST, CAP_NET_RAW, CAP_NET_BIND_SERVICE
 */
func networkAdminEntitlementEnforce(profile Profile) (Profile, error) {
	ociProfile, err := ociProfileConversionCheck(profile, "network.admin")
	if err != nil {
		return nil, err
	}

	capsToAdd := []osdefs.Capability{osdefs.CapNetAdmin, osdefs.CapNetRaw, osdefs.CapNetBindService, osdefs.CapNetBroadcast}
	ociProfile.AddCaps(capsToAdd...)

	return profile, nil
}
