package reqver

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Version is used to work with go versions.
type Version struct {
	Major int
	Minor int
	Patch int
}

// ParseVersion returns the Version object with major, minor and patch fields.
// Returns error if there are problems parsing version.
func ParseVersion() (*Version, error) {
	return parseVersion(runtime.Version())
}

// ParseVersionFromString returns the Version object with major, minor and patch fields from string passed.
// Returns error if there are problems parsing version.
func ParseVersionFromString(version string) (*Version, error) {
	return parseVersion(version)
}

// parseVersion internal implementation to parse go version from string
// Returns error if there are problems parsing version.
func parseVersion(stringVersion string) (*Version, error) {
	versionParts := strings.Split(stringVersion, ".")
	if len(versionParts) < 2 {
		return nil, fmt.Errorf("error parsing go version")
	}

	major, err := strconv.Atoi(strings.TrimPrefix(versionParts[0], "go"))
	if err != nil {
		return nil, err
	}

	minor, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return nil, err
	}

	if len(versionParts) == 2 {
		version := Version{
			Major: major,
			Minor: minor,
			Patch: 0,
		}

		return &version, nil
	}

	patch, err := strconv.Atoi(versionParts[2])
	if err != nil {
		return nil, err
	}

	version := Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}

	return &version, nil
}

// String returns Go version in a string format
func (v *Version) String() string {
	return fmt.Sprintf("go version go%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// region Comparisons

func (v *Version) IsValid() bool {
	return v.Major >= 0 && v.Minor >= 0 && v.Patch >= 0
}

// IsHigherOrEqual compares go versions.
// Returns true if given version is higher or equal to the current one, otherwise - false.
func (v *Version) IsHigherOrEqual(version *Version) bool {
	if !version.IsValid() || !v.IsValid() {
		return false
	}

	if v.Major > version.Major {
		return true
	} else if v.Major < version.Major {
		return false
	} else {
		if v.Minor > version.Minor {
			return true
		} else if v.Minor < version.Minor {
			return false
		} else {
			if v.Patch > version.Patch {
				return true
			} else if v.Patch < version.Patch {
				return false
			} else {
				return true
			}
		}
	}
}

// IsHigher compares go versions.
// Returns true if given version is higher than the current one, otherwise - false.
func (v *Version) IsHigher(version *Version) bool {
	if !version.IsValid() || !v.IsValid() {
		return false
	}

	if v.Major > version.Major {
		return true
	} else if v.Major < version.Major {
		return false
	} else {
		if v.Minor > version.Minor {
			return true
		} else if v.Minor < version.Minor {
			return false
		} else {
			if v.Patch > version.Patch {
				return true
			} else if v.Patch < version.Patch {
				return false
			} else {
				return false
			}
		}
	}
}

// IsEqual compares go versions.
// Returns true if given version is equal to the current one, otherwise - false.
func (v *Version) IsEqual(version *Version) bool {
	if !version.IsValid() || !v.IsValid() {
		return false
	}

	return v.Major == version.Major && v.Minor == version.Minor && v.Patch == version.Patch
}

// IsSmallerOrEqual compares go versions.
// Returns true if given version is smaller or equal to the current one, otherwise - false.
func (v *Version) IsSmallerOrEqual(version *Version) bool {
	if !version.IsValid() || !v.IsValid() {
		return false
	}

	if v.Major > version.Major {
		return false
	} else if v.Major < version.Major {
		return true
	} else {
		if v.Minor > version.Minor {
			return false
		} else if v.Minor < version.Minor {
			return true
		} else {
			if v.Patch > version.Patch {
				return false
			} else if v.Patch < version.Patch {
				return true
			} else {
				return true
			}
		}
	}
}

// IsSmaller compares go versions.
// Returns true if given version is smaller than the current one, otherwise - false.
func (v *Version) IsSmaller(version *Version) bool {
	if !version.IsValid() || !v.IsValid() {
		return false
	}

	if v.Major > version.Major {
		return false
	} else if v.Major < version.Major {
		return true
	} else {
		if v.Minor > version.Minor {
			return false
		} else if v.Minor < version.Minor {
			return true
		} else {
			if v.Patch > version.Patch {
				return false
			} else if v.Patch < version.Patch {
				return true
			} else {
				return false
			}
		}
	}
}

// endregion Comparisons
