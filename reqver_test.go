package reqver

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var currentVersion = &Version{
	Major: 1,
	Minor: 20,
	Patch: 3,
}

func TestVersion_ParseVersion(t *testing.T) {
	asserter := assert.New(t)
	result, err := ParseVersion()

	asserter.NoError(err)
	asserter.Equal(currentVersion.Major, result.Major)
	asserter.Equal(currentVersion.Minor, result.Minor)
	asserter.Equal(currentVersion.Patch, result.Patch)
}

func TestVersion_ParseVersionFromString(t *testing.T) {
	asserter := assert.New(t)
	t.Run("Invalid string", func(t *testing.T) {
		result, err := ParseVersionFromString("Hello world")
		asserter.Error(err)
		asserter.Nil(result)
	})
	t.Run("Success", func(t *testing.T) {
		result, err := ParseVersion()

		asserter.NoError(err)
		asserter.Equal(currentVersion.Major, result.Major)
		asserter.Equal(currentVersion.Minor, result.Minor)
		asserter.Equal(currentVersion.Patch, result.Patch)
	})
}

func TestVersion_String(t *testing.T) {
	asserter := assert.New(t)

	expectedStr := fmt.Sprintf("go version go%d.%d.%d", currentVersion.Major, currentVersion.Minor, currentVersion.Patch)
	actualStr := currentVersion.String()

	asserter.Equal(expectedStr, actualStr)
}

// region Comparisons

func TestVersion_IsHigherOrEqual(t *testing.T) {
	asserter := assert.New(t)

	t.Run("empty object", func(t *testing.T) {
		versionToCheck := Version{}
		asserter.False(versionToCheck.IsHigher(currentVersion))
	})

	t.Run("major smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: -1,
			Minor: 1,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsHigherOrEqual(currentVersion))
	})
	t.Run("major bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: 1000,
			Minor: 1,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsHigherOrEqual(currentVersion))
	})
	t.Run("minor smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: -1,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsHigherOrEqual(currentVersion))
	})
	t.Run("minor bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: 1000,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsHigherOrEqual(currentVersion))
	})
	t.Run("patch smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: -1,
		}

		asserter.False(versionToCheck.IsHigherOrEqual(currentVersion))
	})
	t.Run("patch bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: 1000,
		}

		asserter.True(versionToCheck.IsHigherOrEqual(currentVersion))
	})
	t.Run("equal", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: currentVersion.Patch,
		}

		asserter.True(versionToCheck.IsHigherOrEqual(currentVersion))
	})
}

func TestVersion_IsHigher(t *testing.T) {
	asserter := assert.New(t)

	t.Run("major smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: -1,
			Minor: 1,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsHigher(currentVersion))
	})
	t.Run("major bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: 1000,
			Minor: 1,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsHigher(currentVersion))
	})
	t.Run("minor smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: -1,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsHigher(currentVersion))
	})
	t.Run("minor bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: 1000,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsHigher(currentVersion))
	})
	t.Run("patch smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: -1,
		}

		asserter.False(versionToCheck.IsHigher(currentVersion))
	})
	t.Run("patch bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: 1000,
		}

		asserter.True(versionToCheck.IsHigher(currentVersion))
	})
	t.Run("equal", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: currentVersion.Patch,
		}

		asserter.False(versionToCheck.IsHigher(currentVersion))
	})
}

func TestVersion_IsEqual(t *testing.T) {
	asserter := assert.New(t)

	t.Run("not equal", func(t *testing.T) {
		versionToCheck := Version{
			Major: -1,
			Minor: 0,
			Patch: 0,
		}

		asserter.False(versionToCheck.IsEqual(currentVersion))
	})
	t.Run("equal", func(t *testing.T) {
		versionToCheck := currentVersion
		asserter.True(versionToCheck.IsEqual(currentVersion))
	})
}

func TestVersion_IsSmallerOrEqual(t *testing.T) {
	asserter := assert.New(t)

	t.Run("major smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: 0,
			Minor: 1,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
	t.Run("major bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: 1000,
			Minor: 1,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
	t.Run("minor smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: 0,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
	t.Run("minor bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: 1000,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
	t.Run("patch smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: 0,
		}

		asserter.True(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
	t.Run("patch bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: 1000,
		}

		asserter.False(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
	t.Run("equal", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: currentVersion.Patch,
		}

		asserter.True(versionToCheck.IsSmallerOrEqual(currentVersion))
	})
}

func TestVersion_IsSmaller(t *testing.T) {
	asserter := assert.New(t)

	t.Run("major smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: 0,
			Minor: 1,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsSmaller(currentVersion))
	})
	t.Run("major bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: 1000,
			Minor: 1,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsSmaller(currentVersion))
	})
	t.Run("minor smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: 0,
			Patch: 2,
		}

		asserter.True(versionToCheck.IsSmaller(currentVersion))
	})
	t.Run("minor bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: 1000,
			Patch: 2,
		}

		asserter.False(versionToCheck.IsSmaller(currentVersion))
	})
	t.Run("patch smaller", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: 0,
		}

		asserter.True(versionToCheck.IsSmaller(currentVersion))
	})
	t.Run("patch bigger", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: 1000,
		}

		asserter.False(versionToCheck.IsSmaller(currentVersion))
	})
	t.Run("equal", func(t *testing.T) {
		versionToCheck := Version{
			Major: currentVersion.Major,
			Minor: currentVersion.Minor,
			Patch: currentVersion.Patch,
		}

		asserter.False(versionToCheck.IsSmaller(currentVersion))
	})
}

// endregion Comparisons
