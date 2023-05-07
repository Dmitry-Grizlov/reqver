package reqver

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

// ParseVersion returns the Version object with major, minor and patch fields.
// Returns error if there are problems parsing version.
func ParseVersion() (*Version, error) {
	myVersion := runtime.Version()
	versionParts := strings.Split(myVersion, ".")
	if len(versionParts) < 2 {
		return nil, fmt.Errorf("error parsing go version")
	}

	majorParts := strings.Split(versionParts[0], "go")
	if len(majorParts) != 2 {
		return nil, fmt.Errorf("error parsing go version")
	}

	major, err := strconv.Atoi(majorParts[1])
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

// region Getters

// GetMajor returns major number from version
func (v *Version) GetMajor() int {
	return v.Major
}

// GetMinor returns minor number from version
func (v *Version) GetMinor() int {
	return v.Minor
}

// GetPatch returns patch number from version
func (v *Version) GetPatch() int {
	return v.Patch
}

// endregion Getters

// region Comparisons

// IsHigherOrEqual compares go versions.
// Returns true if given version is higher or equal to the current one, otherwise - false.
func (v *Version) IsHigherOrEqual(version Version) bool {
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
func (v *Version) IsHigher(version Version) bool {
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
func (v *Version) IsEqual(version Version) bool {
	return v.Major == version.Major && v.Minor == version.Minor && v.Patch == version.Patch
}

// IsSmallerOrEqual compares go versions.
// Returns true if given version is smaller or equal to the current one, otherwise - false.
func (v *Version) IsSmallerOrEqual(version Version) bool {
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
func (v *Version) IsSmaller(version Version) bool {
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
