// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package concurrentbatchprocessor

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestUnmarshalDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	cm := confmap.New()
	assert.NoError(t, cm.Unmarshal(cfg))
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
}

func TestUnmarshalConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NoError(t, cm.Unmarshal(cfg))
	assert.Equal(t,
		&Config{
			SendBatchSize:            uint32(10000),
			SendBatchMaxSize:         uint32(11000),
			Timeout:                  time.Second * 10,
			MetadataCardinalityLimit: 1000,
			MaxInFlightSizeMiB:       12345,
		}, cfg)
}

func TestValidateConfig_DefaultBatchMaxSize(t *testing.T) {
	cfg := &Config{
		SendBatchSize:      100,
		SendBatchMaxSize:   0,
		MaxInFlightSizeMiB: 1,
	}
	assert.NoError(t, cfg.Validate())
}

func TestValidateConfig_ValidBatchSizes(t *testing.T) {
	cfg := &Config{
		SendBatchSize:      100,
		SendBatchMaxSize:   1000,
		MaxInFlightSizeMiB: 1,
	}
	assert.NoError(t, cfg.Validate())

}

func TestValidateConfig_InvalidBatchSize(t *testing.T) {
	cfg := &Config{
		SendBatchSize:      1000,
		SendBatchMaxSize:   100,
		MaxInFlightSizeMiB: 1,
	}
	assert.Error(t, cfg.Validate())
}

func TestValidateConfig_InvalidTimeout(t *testing.T) {
	cfg := &Config{
		Timeout:            -time.Second,
		MaxInFlightSizeMiB: 1,
	}
	assert.Error(t, cfg.Validate())
}

func TestValidateConfig_InvalidZero(t *testing.T) {
	cfg := &Config{}
	assert.Error(t, cfg.Validate())
}
