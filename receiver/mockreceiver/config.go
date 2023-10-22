// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver"

import (
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver/internal/metadata"
)

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	MetricsBuilderConfig                    metadata.MetricsBuilderConfig `mapstructure:",squash"`
}

func (cfg *Config) Validate() error {
	return nil
}
