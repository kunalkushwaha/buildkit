package entitlement

import "fmt"

// ProfileType is a string identifying a specific and unique security profile type
type ProfileType string

// Profile is an abstract interface which represents security profiles.
// Each security profile has its own type and its own API as needs
// may vary across different profile formats.
type Profile interface {
	GetType() ProfileType
}

func ociProfileConversionCheck(profile Profile, entitlementID string) (*OCIProfile, error) {
	if profile == nil {
		return nil, fmt.Errorf("profile is nil for %s", entitlementID)
	}

	if profile.GetType() != OCIProfileType {
		return nil, fmt.Errorf("%s not implemented for non-OCI profiles", entitlementID)
	}

	ociProfile, ok := profile.(*OCIProfile)
	if !ok {
		return nil, fmt.Errorf("%s: error converting to OCI profile", entitlementID)
	}

	return ociProfile, nil
}

/* Add capability if not present to capability set */
func addCapToList(capList []string, capToAdd string) []string {
	for _, cap := range capList {
		if cap == capToAdd {
			return capList
		}
	}

	return append(capList, capToAdd)
}

/* Remove capability if present from capability set */
func removeCapFromList(capList []string, capToRemove string) []string {
	for index, cap := range capList {
		if cap == capToRemove {
			capList = append(capList[:index], capList[index+1:]...)
			break
		}
	}

	return capList
}
