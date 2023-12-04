// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hierarchicalresourcequota // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/hierarchicalresourcequota"

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func RecordMetrics(mb *metadata.MetricsBuilder, hrq *unstructured.Unstructured, ts pcommon.Timestamp) {
	name, _, _ := unstructured.NestedString(hrq.Object, "metadata", "name")
	uid, _, _ := unstructured.NestedString(hrq.Object, "metadata", "uid")
	namespace, _, _ := unstructured.NestedString(hrq.Object, "metadata", "namespace")
	statusHard, _, _ := unstructured.NestedMap(hrq.Object, "status", "hard")
	for k, v := range statusHard {
		q := resource.MustParse(v.(string))
		val := q.Value()
		if strings.HasSuffix(string(k), ".cpu") {
			val = q.MilliValue()
		}
		mb.RecordK8sHierarchicalResourceQuotaHardLimitDataPoint(ts, val, string(k))
	}
	statusUsed, _, _ := unstructured.NestedMap(hrq.Object, "status", "used")
	for k, v := range statusUsed {
		q := resource.MustParse(v.(string))
		val := q.Value()
		if strings.HasSuffix(string(k), ".cpu") {
			val = q.MilliValue()
		}
		mb.RecordK8sHierarchicalResourceQuotaHardLimitDataPoint(ts, val, string(k))
	}

	rb := mb.NewResourceBuilder()
	rb.SetK8sHierarchicalresourcequotaUID(string(uid))
	rb.SetK8sHierarchicalresourcequotaName(name)
	rb.SetK8sNamespaceName(namespace)
	mb.EmitForResource(metadata.WithResource(rb.Emit()))
}
