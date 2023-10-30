// GENERATED FILE

package linux

import (
	"embed"
	"github.com/amitybell/piper-asset"
)

var (
	//go:embed dist.tar.zst hash.txt
	fs embed.FS

	Asset = asset.Asset{Name: "linux", FS: fs}
)
