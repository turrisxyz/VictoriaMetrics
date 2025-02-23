package gce

import (
	"flag"
	"fmt"
	"time"
)

// SDCheckInterval defines interval for targets refresh.
var SDCheckInterval = flag.Duration("promscrape.gceSDCheckInterval", time.Minute, "Interval for checking for changes in gce. "+
	"This works only if gce_sd_configs is configured in '-promscrape.config' file. "+
	"See https://prometheus.io/docs/prometheus/latest/configuration/configuration/#gce_sd_config for details")

// SDConfig represents service discovery config for gce.
//
// See https://prometheus.io/docs/prometheus/latest/configuration/configuration/#gce_sd_config
type SDConfig struct {
	Project string   `yaml:"project"`
	Zone    ZoneYAML `yaml:"zone"`
	Filter  string   `yaml:"filter,omitempty"`
	// RefreshInterval time.Duration `yaml:"refresh_interval"`
	// refresh_interval is obtained from `-promscrape.gceSDCheckInterval` command-line option.
	Port         *int    `yaml:"port,omitempty"`
	TagSeparator *string `yaml:"tag_separator,omitempty"`
}

// ZoneYAML holds info about zones.
type ZoneYAML struct {
	zones []string
}

// UnmarshalYAML implements yaml.Unmarshaler
func (z *ZoneYAML) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v interface{}
	if err := unmarshal(&v); err != nil {
		return err
	}
	var zones []string
	switch t := v.(type) {
	case string:
		zones = []string{t}
	case []interface{}:
		for _, vv := range t {
			zone, ok := vv.(string)
			if !ok {
				return fmt.Errorf("unexpected zone type detected: %T; contents: %#v", vv, vv)
			}
			zones = append(zones, zone)
		}
	default:
		return fmt.Errorf("unexpected type unmarshaled for ZoneYAML: %T; contents: %#v", v, v)
	}
	z.zones = zones
	return nil
}

// MarshalYAML implements yaml.Marshaler
func (z ZoneYAML) MarshalYAML() (interface{}, error) {
	return z.zones, nil
}

// GetLabels returns gce labels according to sdc.
func (sdc *SDConfig) GetLabels(baseDir string) ([]map[string]string, error) {
	cfg, err := getAPIConfig(sdc)
	if err != nil {
		return nil, fmt.Errorf("cannot get API config: %w", err)
	}
	ms := getInstancesLabels(cfg)
	return ms, nil
}

// MustStop stops further usage for sdc.
func (sdc *SDConfig) MustStop() {
	configMap.Delete(sdc)
}
