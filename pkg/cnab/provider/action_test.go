package cnabprovider

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"get.porter.sh/porter/pkg/cnab"
	"get.porter.sh/porter/pkg/config"
	"get.porter.sh/porter/pkg/storage"
	"github.com/cnabio/cnab-go/bundle"
	"github.com/cnabio/cnab-go/driver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddRelocation(t *testing.T) {
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/relocation-mapping.json")
	require.NoError(t, err)

	d := NewTestRuntime(t)
	defer d.Close()

	var args ActionArguments
	require.NoError(t, json.Unmarshal(data, &args.BundleReference.RelocationMap))

	opConf := d.AddRelocation(args)

	invoImage := bundle.InvocationImage{}
	invoImage.Image = "gabrtv/microservice@sha256:cca460afa270d4c527981ef9ca4989346c56cf9b20217dcea37df1ece8120687"

	op := &driver.Operation{
		Files: make(map[string]string),
		Image: invoImage,
	}
	err = opConf(op)
	assert.NoError(t, err)

	mapping, ok := op.Files["/cnab/app/relocation-mapping.json"]
	assert.True(t, ok)
	assert.Equal(t, string(data), mapping)
	assert.Equal(t, "my.registry/microservice@sha256:cca460afa270d4c527981ef9ca4989346c56cf9b20217dcea37df1ece8120687", op.Image.Image)

}

func TestRuntime_Execute(t *testing.T) {
	type args struct {
		ctx  context.Context
		args ActionArguments
		cfg  *config.Config
	}

	r := NewTestRuntime(t)
	b, _ := ioutil.ReadFile("testdata/template-bundle.json")
	bun := bundle.Bundle{}
	json.Unmarshal(b, &bun)
	xbun := cnab.NewBundle(bun)


	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Execute With Bundle Parameters",
			args: args{
				ctx: context.TODO(),
				args: ActionArguments{
					Action: "install",
					Installation: r.TestInstallations.CreateInstallation(storage.Installation{}),
					BundleReference: cnab.BundleReference{Reference: cnab.OCIReference{}, Definition: xbun},
					Driver: "docker",
					AllowDockerHostAccess: true,
					Params: make(map[string]interface{}),
					},
				cfg: r.Config,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Execute(tt.args.ctx, tt.args.args, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Runtime.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
