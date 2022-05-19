package backend

import (
	"os"

	"github.com/BitsOfAByte/marketxiv/build"
)

func IsPreview() bool {
	return build.Version == "v0.0.1-next" && os.Getenv("MXIV_PREVIEW_CONSENT") != "true"
}
