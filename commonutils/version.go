package commonutils

const (
	Version = "1.0.23"
	Name    = "commonutils"
)

func GetUtilsVersion() (string, string) {
	return Version, Name
}
