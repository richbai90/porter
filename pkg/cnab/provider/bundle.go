package cnabprovider

import (
	"context"

	"get.porter.sh/porter/pkg/cnab"
	configadapter "get.porter.sh/porter/pkg/cnab/config-adapter"
	"get.porter.sh/porter/pkg/config"
	"get.porter.sh/porter/pkg/manifest"
	"get.porter.sh/porter/pkg/runtime"
	"get.porter.sh/porter/pkg/yaml"
	"github.com/pkg/errors"
)

func (r *Runtime) LoadBundle(bundleFile string) (cnab.ExtendedBundle, error) {
	return cnab.LoadBundle(r.Context, bundleFile)
}

func (r *Runtime) ProcessBundleFromFile(bundleFile string) (cnab.ExtendedBundle, error) {
	b, err := r.LoadBundle(bundleFile)
	if err != nil {
		return cnab.ExtendedBundle{}, err
	}

	return r.ProcessBundle(b)
}

func (r *Runtime) ParseBundle(ctx context.Context, b *cnab.ExtendedBundle, args ActionArguments, cfg *config.Config) {

	stamp, err := configadapter.LoadStamp(*b)
	if err := errors.Wrap(err, "Failed to load stamp from bundle"); err != nil {
		return
	}

	mbytes, err := stamp.DecodeManifest()

	if err := errors.Wrap(err, "Failed to decode manifest"); err != nil {
		return
	}
	m := manifest.Manifest{}
	err = yaml.Unmarshal(mbytes, &m)

	if err := errors.Wrap(err, "Failed to unmarshall YAML"); err != nil {
		return
	}

	rm := runtime.NewRuntimeManifest(cfg.Context, "runtime", &m)

	step := manifest.Step{
		Data: b.Custom,
	}

	rm.ResolveStep(&step)

	b.Custom = step.Data

}

func (r *Runtime) ProcessBundle(b cnab.ExtendedBundle) (cnab.ExtendedBundle, error) {
	err := b.Validate()
	if err != nil {
		return b, errors.Wrap(err, "invalid bundle")
	}

	return b, r.ProcessRequiredExtensions(b)
}
