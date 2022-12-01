package config_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYAMLAnchors(t *testing.T) {
	b := load(t, "./yaml_anchors")
	assert.Len(t, b.Config.Resources.Workflows, 1)

	j := b.Config.Resources.Workflows["my_workflow"]
	require.Len(t, j.Tasks, 2)

	t0 := j.Tasks[0]
	t1 := j.Tasks[1]
	require.NotNil(t, t0)
	require.NotNil(t, t1)

	assert.Equal(t, "10.4.x-scala2.12", t0.NewCluster.SparkVersion)
	assert.Equal(t, "10.4.x-scala2.12", t1.NewCluster.SparkVersion)
}