// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver"

import (
	"fmt"

	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver/internal/metadata"
)

type MockReceiverConfig struct {
	MockValue string `mapstructure:"mock_value"`
}

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	MetricsBuilderConfig                    metadata.MetricsBuilderConfig `mapstructure:",squash"`
	MockReceiverConfig                      `mapstructure:",squash"`
}

func (cfg *Config) Validate() error {
	mockValue := cfg.MockValue
	if mockValue == "" {
		return fmt.Errorf("missing mock.value: '%s'", cfg.MockValue)
	}
	return nil
}
