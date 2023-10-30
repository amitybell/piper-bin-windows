// GENERATED FILE

package windows

import (
	"embed"
	"github.com/amitybell/piper-asset"
)

var (
	//go:embed dist.tar.zst hash.txt
	fs embed.FS

	Asset = asset.Asset{Name: "windows", FS: fs}
)
