// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockreceiver

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver/internal/metadata"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		desc        string
		value       string
		errExpected bool
		errText     string
	}{
		{
			desc:        "default_endpoint",
			value:       "1",
			errExpected: false,
		},
		{
			desc:        "custom_value",
			value:       "5",
			errExpected: false,
		},
		{
			desc:        "invalid_value",
			value:       "",
			errExpected: true,
			errText:     "missing mock.value: ''",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cfg := NewFactory().CreateDefaultConfig().(*Config)
			cfg.MockValue = tc.value
			err := component.ValidateConfig(cfg)
			if tc.errExpected {
				require.EqualError(t, err, tc.errText)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	sub, err := cm.Sub(component.NewIDWithName(metadata.Type, "").String())
	require.NoError(t, err)
	require.NoError(t, component.UnmarshalConfig(sub, cfg))

	expected := factory.CreateDefaultConfig().(*Config)
	expected.MockValue = "5"
	expected.CollectionInterval = 10 * time.Second

	require.Equal(t, expected, cfg)
}
