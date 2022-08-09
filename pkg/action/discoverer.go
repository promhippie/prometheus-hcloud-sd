package action

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
)

var (
	// providerPrefix defines the general prefix for all labels.
	providerPrefix = model.MetaLabelPrefix + "hcloud_"

	// Labels defines all available labels for this provider.
	Labels = map[string]string{
		"datacenterName":    providerPrefix + "datacenter",
		"imageName":         providerPrefix + "image_name",
		"imageType":         providerPrefix + "image_type",
		"labelPrefix":       providerPrefix + "label_",
		"locationCity":      providerPrefix + "city",
		"locationCountry":   providerPrefix + "country",
		"locationName":      providerPrefix + "location",
		"name":              providerPrefix + "name",
		"osFlavor":          providerPrefix + "os_flavor",
		"osVersion":         providerPrefix + "os_version",
		"project":           providerPrefix + "project",
		"publicIPv4":        providerPrefix + "public_ipv4",
		"publicIPv6":        providerPrefix + "public_ipv6",
		"serverTypeCores":   providerPrefix + "cores",
		"serverTypeCPU":     providerPrefix + "cpu",
		"serverTypeDisk":    providerPrefix + "disk",
		"serverTypeMemory":  providerPrefix + "memory",
		"serverTypeName":    providerPrefix + "type",
		"serverTypeStorage": providerPrefix + "storage",
		"status":            providerPrefix + "status",
	}

	// replacer defines a list of characters that gets replaced.
	replacer = strings.NewReplacer(
		".", "_",
		"-", "_",
	)
)

// Discoverer implements the Prometheus discoverer interface.
type Discoverer struct {
	clients   map[string]*hcloud.Client
	logger    log.Logger
	refresh   int
	separator string
	lasts     map[string]struct{}
}

// Run initializes fetching the targets for service discovery.
func (d Discoverer) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
	ticker := time.NewTicker(time.Duration(d.refresh) * time.Second)

	for {
		targets, err := d.getTargets(ctx)

		if err == nil {
			ch <- targets
		}

		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (d *Discoverer) getTargets(ctx context.Context) ([]*targetgroup.Group, error) {
	current := make(map[string]struct{})
	targets := make([]*targetgroup.Group, 0)

	for project, client := range d.clients {

		now := time.Now()
		servers, err := client.Server.All(ctx)
		requestDuration.WithLabelValues(project).Observe(time.Since(now).Seconds())

		if err != nil {
			level.Warn(d.logger).Log(
				"msg", "Failed to fetch servers",
				"project", project,
				"err", err,
			)

			requestFailures.WithLabelValues(project).Inc()
			continue
		}

		level.Debug(d.logger).Log(
			"msg", "Requested servers",
			"project", project,
			"count", len(servers),
		)

		for _, server := range servers {
			var (
				imageType string
				imageName string
				osFlavor  string
				osVersion string
			)

			if server.Image != nil {
				imageType = string(server.Image.Type)
				imageName = server.Image.Name
				osFlavor = server.Image.OSFlavor
				osVersion = server.Image.OSVersion
			}

			target := &targetgroup.Group{
				Source: fmt.Sprintf("hcloud/%d", server.ID),
				Targets: []model.LabelSet{
					{
						model.AddressLabel: model.LabelValue(server.PublicNet.IPv4.IP.String()),
					},
				},
				Labels: model.LabelSet{
					model.AddressLabel:                           model.LabelValue(server.PublicNet.IPv4.IP.String()),
					model.LabelName(Labels["project"]):           model.LabelValue(project),
					model.LabelName(Labels["name"]):              model.LabelValue(server.Name),
					model.LabelName(Labels["status"]):            model.LabelValue(server.Status),
					model.LabelName(Labels["publicIPv4"]):        model.LabelValue(server.PublicNet.IPv4.IP.String()),
					model.LabelName(Labels["publicIPv6"]):        model.LabelValue(server.PublicNet.IPv6.IP.String()),
					model.LabelName(Labels["serverTypeName"]):    model.LabelValue(server.ServerType.Name),
					model.LabelName(Labels["serverTypeCores"]):   model.LabelValue(strconv.Itoa(int(server.ServerType.Cores))),
					model.LabelName(Labels["serverTypeMemory"]):  model.LabelValue(strconv.Itoa(int(server.ServerType.Memory))),
					model.LabelName(Labels["serverTypeDisk"]):    model.LabelValue(strconv.Itoa(int(server.ServerType.Disk))),
					model.LabelName(Labels["serverTypeStorage"]): model.LabelValue(server.ServerType.StorageType),
					model.LabelName(Labels["serverTypeCPU"]):     model.LabelValue(server.ServerType.CPUType),
					model.LabelName(Labels["datacenterName"]):    model.LabelValue(server.Datacenter.Name),
					model.LabelName(Labels["locationName"]):      model.LabelValue(server.Datacenter.Location.Name),
					model.LabelName(Labels["locationCity"]):      model.LabelValue(server.Datacenter.Location.City),
					model.LabelName(Labels["locationCountry"]):   model.LabelValue(server.Datacenter.Location.Country),
					model.LabelName(Labels["imageType"]):         model.LabelValue(imageType),
					model.LabelName(Labels["imageName"]):         model.LabelValue(imageName),
					model.LabelName(Labels["osFlavor"]):          model.LabelValue(osFlavor),
					model.LabelName(Labels["osVersion"]):         model.LabelValue(osVersion),
				},
			}

			for key, value := range server.Labels {
				target.Labels[model.LabelName(normalizeLabel(Labels["labelPrefix"]+key))] = model.LabelValue(value)
			}

			level.Debug(d.logger).Log(
				"msg", "Server added",
				"project", project,
				"source", target.Source,
			)

			current[target.Source] = struct{}{}
			targets = append(targets, target)
		}

	}

	for k := range d.lasts {
		if _, ok := current[k]; !ok {
			level.Debug(d.logger).Log(
				"msg", "Server deleted",
				"source", k,
			)

			targets = append(
				targets,
				&targetgroup.Group{
					Source: k,
				},
			)
		}
	}

	d.lasts = current
	return targets, nil
}

func normalizeLabel(val string) string {
	return replacer.Replace(val)
}
