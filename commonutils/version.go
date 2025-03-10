package commonutils

const (
	Version = "1.0.20"
	Name    = "commonutils"
)

func GetUtilsVersion() (string, string) {
	return Version, Name
}
