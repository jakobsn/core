// +build linux

package optimus

import (
	"fmt"
	"os"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	defaultCPUPeriod = uint64(100000)
)

func RestrictUsage(cfg *RestrictionsConfig) (Deleter, error) {
	path := cgroups.StaticPath(cfg.Name)
	control, err := cgroups.Load(cgroups.V1, path)
	switch err {
	case nil:
	case cgroups.ErrCgroupDeleted:
		// Create an empty cgroup, we update it later.
		control, err = cgroups.New(cgroups.V1, path, &specs.LinuxResources{})
		if err != nil {
			return nil, fmt.Errorf("failed to create cgroup %s: %v", cfg.Name, err)
		}
	default:
		return nil, fmt.Errorf("failed to load cgroup %s: %v", cfg.Name, err)
	}

	quota := int64(float64(defaultCPUPeriod) * cfg.CPUCount)
	period := defaultCPUPeriod
	resources := &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota:  &quota,
			Period: &period,
		},
	}

	if err = control.Update(resources); err != nil {
		return nil, fmt.Errorf("failed to set resource limit on cgroup %s: %v", cfg.Name, err)
	}

	if err := control.Add(cgroups.Process{Pid: os.Getpid()}); err != nil {
		return nil, fmt.Errorf("failed to add Optimus into cgroup: %v", err)
	}

	return control, nil
}
