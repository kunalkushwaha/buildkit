package entitlement

import (
	"reflect"
	"testing"

	"github.com/moby/buildkit/util/entitlement/osdefs"
	"github.com/moby/buildkit/util/entitlement/testutils"
	"github.com/stretchr/testify/require"
)

func TestProfileSecurityConfined(t *testing.T) {
	testSyscall := osdefs.SysExit
	ociProfile := NewOCIProfile(testutils.TestSpec(), "test-profile")

	require.NotNil(t, ociProfile.OCI)
	require.NotNil(t, ociProfile.OCI.Linux)
	require.NotNil(t, ociProfile.OCI.Linux.Seccomp)

	syscalls := []osdefs.Syscall{testSyscall}
	ociProfile.AllowSyscalls(syscalls...)

	seccompProfileWithTestSys := *ociProfile.OCI.Linux.Seccomp

	ociProfile.AllowSyscalls(syscalls...)
	didAdd := reflect.DeepEqual(seccompProfileWithTestSys, *ociProfile.OCI.Linux.Seccomp)
	require.True(t, didAdd, "Syscall was not already added to the seccomp profile")
}
