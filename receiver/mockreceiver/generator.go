// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mockreceiver/internal/metadata"
)

type mockGenerator struct {
	settings component.TelemetrySettings
	cfg      *Config
	mb       *metadata.MetricsBuilder
	value    string
}

func newMockGenerator(
	settings receiver.CreateSettings,
	cfg *Config,
	value string,
) *mockGenerator {
	mg := &mockGenerator{
		settings: settings.TelemetrySettings,
		cfg:      cfg,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
		value:    value,
	}

	return mg
}

func (r *mockGenerator) start(_ context.Context, _ component.Host) error {
	// noop
	return nil
}

func (r *mockGenerator) scrape(context.Context) (pmetric.Metrics, error) {
	errs := &scrapererror.ScrapeErrors{}
	now := pcommon.NewTimestampFromTime(time.Now())
	err := r.mb.RecordValueDataPoint(now, r.value)
	if err != nil {
		errs.Add(err)
	}
	rb := r.mb.NewResourceBuilder()
	rb.SetMockValue(r.value)
	return r.mb.Emit(), errs.Combine()
}
