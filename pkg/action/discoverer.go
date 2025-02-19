package action

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
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
		"privateIPv4":       providerPrefix + "ipv4_",
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
	logger    *slog.Logger
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

		networks, err := client.Network.All(ctx)

		if err != nil {
			d.logger.Warn("Failed to fetch networks",
				"project", project,
				"err", err,
			)

			requestFailures.WithLabelValues(project).Inc()
			continue
		}

		servers, err := client.Server.All(ctx)

		if err != nil {
			d.logger.Warn("Failed to fetch servers",
				"project", project,
				"err", err,
			)

			requestFailures.WithLabelValues(project).Inc()
			continue
		}

		requestDuration.WithLabelValues(project).Observe(time.Since(now).Seconds())

		d.logger.Debug("Requested servers",
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

			for _, priv := range server.PrivateNet {
				for _, network := range networks {
					if network.ID == priv.Network.ID {
						target.Labels[model.LabelName(normalizeNetwork(Labels["privateIPv4"]+network.Name))] = model.LabelValue(priv.IP.String())
						break
					}
				}
			}

			d.logger.Debug("Server added",
				"project", project,
				"source", target.Source,
			)

			current[target.Source] = struct{}{}
			targets = append(targets, target)
		}

	}

	for k := range d.lasts {
		if _, ok := current[k]; !ok {
			d.logger.Debug("Server deleted",
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

func normalizeNetwork(val string) string {
	return replacer.Replace(val)
}
