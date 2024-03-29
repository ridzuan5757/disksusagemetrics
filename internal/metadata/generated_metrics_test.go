// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiskFreeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiskTotalDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiskUsedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiskUtilizationDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMountFreeDataPoint(ts, 1, "mount.point-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMountTotalDataPoint(ts, 1, "mount.point-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMountUsedDataPoint(ts, 1, "mount.point-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMountUtilizationDataPoint(ts, 1, "mount.point-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "disk.free":
					assert.False(t, validatedMetrics["disk.free"], "Found a duplicate in the metrics slice: disk.free")
					validatedMetrics["disk.free"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Amount of space available on the file system.", ms.At(i).Description())
					assert.Equal(t, "byte", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "disk.total":
					assert.False(t, validatedMetrics["disk.total"], "Found a duplicate in the metrics slice: disk.total")
					validatedMetrics["disk.total"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Total size of the file system.", ms.At(i).Description())
					assert.Equal(t, "byte", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "disk.used":
					assert.False(t, validatedMetrics["disk.used"], "Found a duplicate in the metrics slice: disk.used")
					validatedMetrics["disk.used"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Amount of space used on the file system.", ms.At(i).Description())
					assert.Equal(t, "byte", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "disk.utilization":
					assert.False(t, validatedMetrics["disk.utilization"], "Found a duplicate in the metrics slice: disk.utilization")
					validatedMetrics["disk.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of disk usage for the file system.", ms.At(i).Description())
					assert.Equal(t, "utilization", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "mount.free":
					assert.False(t, validatedMetrics["mount.free"], "Found a duplicate in the metrics slice: mount.free")
					validatedMetrics["mount.free"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Amount of space available on the file system.", ms.At(i).Description())
					assert.Equal(t, "byte", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("mount.point")
					assert.True(t, ok)
					assert.EqualValues(t, "mount.point-val", attrVal.Str())
				case "mount.total":
					assert.False(t, validatedMetrics["mount.total"], "Found a duplicate in the metrics slice: mount.total")
					validatedMetrics["mount.total"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Total size of the file system.", ms.At(i).Description())
					assert.Equal(t, "byte", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("mount.point")
					assert.True(t, ok)
					assert.EqualValues(t, "mount.point-val", attrVal.Str())
				case "mount.used":
					assert.False(t, validatedMetrics["mount.used"], "Found a duplicate in the metrics slice: mount.used")
					validatedMetrics["mount.used"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Amount of space used on the file system.", ms.At(i).Description())
					assert.Equal(t, "byte", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("mount.point")
					assert.True(t, ok)
					assert.EqualValues(t, "mount.point-val", attrVal.Str())
				case "mount.utilization":
					assert.False(t, validatedMetrics["mount.utilization"], "Found a duplicate in the metrics slice: mount.utilization")
					validatedMetrics["mount.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of disk usage for the file system.", ms.At(i).Description())
					assert.Equal(t, "utilization", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("mount.point")
					assert.True(t, ok)
					assert.EqualValues(t, "mount.point-val", attrVal.Str())
				}
			}
		})
	}
}
