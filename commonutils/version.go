package commonutils

const (
	Version = "1.0.0"
	Name    = "commonutils"
)

func GetUtilsVersion() (string, string) {
	return Version, Name
}
