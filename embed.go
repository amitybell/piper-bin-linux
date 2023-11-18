// GENERATED FILE

package linux

import (
	"embed"
	"github.com/amitybell/piper-asset"
)

var (
	//go:embed dist.tzst dist.json
	fs embed.FS

	Asset = asset.Asset{Name: "linux", FS: fs}
)
