package sources

import (
	"context"
	"sort"
	"sync"
	"testing"

	"github.com/nginx/agent/sdk/v2/proto"
	tutils "github.com/nginx/agent/v2/test/utils"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/stretchr/testify/assert"
)

func TestNewDiskIOSource(t *testing.T) {
	namespace := "test"
	env := tutils.GetMockEnv()
	actual := NewDiskIOSource(namespace, env)

	assert.Equal(t, "io", actual.group)
	assert.Equal(t, namespace, actual.namespace)
}

func TestDiskIOCollect(t *testing.T) {
	namespace := "test"
	env := tutils.GetMockEnv()
	env.Mock.On("DiskDevices").Return([]string{"disk1", "disk2"}, nil)
	diskio := NewDiskIOSource(namespace, env)
	diskio.diskIOStatsFunc = func(ctx context.Context, names ...string) (map[string]disk.IOCountersStat, error) {
		return map[string]disk.IOCountersStat{"disk1": {}, "unknownDisk": {}}, nil
	}

	ctx := context.TODO()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	channel := make(chan *proto.StatsEntity, 100)
	diskio.Collect(ctx, wg, channel)
	wg.Wait()

	actual := <-channel

	actualMetricNames := []string{}
	for _, simpleMetric := range actual.Simplemetrics {
		actualMetricNames = append(actualMetricNames, simpleMetric.Name)
	}
	sort.Strings(actualMetricNames)
	expected := []string{"test.io.iops_r", "test.io.iops_w", "test.io.kbs_r", "test.io.kbs_w", "test.io.wait_r", "test.io.wait_w"}

	assert.Equal(t, expected, actualMetricNames)
}
