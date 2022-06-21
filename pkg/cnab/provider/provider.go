package cnabprovider

import (
	"context"

	"get.porter.sh/porter/pkg/cnab"
	"get.porter.sh/porter/pkg/config"
)

// CNABProvider is the interface Porter uses to communicate with the CNAB runtime
type CNABProvider interface {
	LoadBundle(bundleFile string) (cnab.ExtendedBundle, error)
	Execute(ctx context.Context, arguments ActionArguments, cfg *config.Config) error
}
