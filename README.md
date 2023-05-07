# reqver 1.0.0

Reqver is a test-covered, lightweight and easy to use library for parsing, printing and comparing go versions. It can be useful if you need to be sure that the certain go version is installed at the user device.

## Usage

### Get library with `go get` command

>`go get "github.com/Dmitry-Grizlov/reqver"`

### Import it in your code

```
    import "github.com/Dmitry-Grizlov/reqver"
```

### Parse current go version

```
    currentVersion, err := reqver.ParseVersion()
```

### Create go version object with version you want to compare with

```
    desiredVersion := Version{
        Major: 1,
        Minor: 2,
        Patch: 3,
    }
```

### You can use these objects for comparison then

```
    if desiredVersion.IsSmaller(currentVersion) {
        fmt.Errorf("Please upgrade your go version.")
    }
```