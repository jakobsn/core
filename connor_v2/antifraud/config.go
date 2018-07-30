package antifraud

import "time"

type LogProcessorConfig struct {
	LogTrackInterval time.Duration `yaml:"log_track_interval" default:"10s"`
	TaskWarmupDelay  time.Duration `yaml:"task_warmup_delay" default:"3m"`
}

type Config struct {
	TaskQuality          float64            `yaml:"task_quality" required:"true"`
	QualityCheckInterval time.Duration      `yaml:"quality_check_interval" default:"15s"`
	ConnectionTimeout    time.Duration      `yaml:"connection_timeout" default:"30s"`
	LogProcessorConfig   LogProcessorConfig `yaml:"log_processor"`
}
