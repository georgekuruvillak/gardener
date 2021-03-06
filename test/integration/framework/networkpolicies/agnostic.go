// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkpolicies

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
)

//Info about pods in Shoot-namespace
var (
	// KubeControllerManagerInfoSecured points to cloud-agnostic kube-apiserver.
	KubeAPIServerInfo = &SourcePod{
		Ports: NewSinglePort(443),
		Pod: NewPod("kube-apiserver", labels.Set{
			"app":  "kubernetes",
			"role": "apiserver",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-kube-apiserver",
			"allow-to-dns",
			"allow-to-public-networks",
			"allow-to-private-networks",
			"allow-to-shoot-networks",
			"deny-all",
		),
	}

	// KubeControllerManagerInfoSecured points to cloud-agnostic kube-controller-manager running on HTTPS port.
	KubeControllerManagerInfoSecured = &SourcePod{
		Ports: NewSinglePort(10257),
		Pod: NewPod("kube-controller-manager-https", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "controller-manager",
		}, ">= 1.13"),
		ExpectedPolicies: sets.NewString(
			"allow-to-public-networks",
			"allow-to-private-networks",
			"allow-from-prometheus",
			"allow-to-dns",
			"allow-to-blocked-cidrs",
			"allow-to-shoot-apiserver",
			"deny-all",
		),
	}

	// KubeControllerManagerInfoSecured points to cloud-agnostic kube-controller-manager running on HTTP port.
	KubeControllerManagerInfoNotSecured = &SourcePod{
		Ports: NewSinglePort(10252),
		Pod: NewPod("kube-controller-manager-http", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "controller-manager",
		}, "< 1.13"),
		ExpectedPolicies: sets.NewString(
			"allow-to-public-networks",
			"allow-to-private-networks",
			"allow-from-prometheus",
			"allow-to-dns",
			"allow-to-blocked-cidrs",
			"allow-to-shoot-apiserver",
			"deny-all",
		),
	}

	// KubeSchedulerInfoSecured points to cloud-agnostic kube-scheduler running on HTTPS port.
	KubeSchedulerInfoSecured = &SourcePod{
		Ports: NewSinglePort(10259),
		Pod: NewPod("kube-scheduler-https", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "scheduler",
		}, ">= 1.13"),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-shoot-apiserver",
			"allow-to-dns",
			"deny-all",
		),
	}

	// KubeSchedulerInfoNotSecured points to cloud-agnostic kube-scheduler running on HTTP port.
	KubeSchedulerInfoNotSecured = &SourcePod{
		Ports: NewSinglePort(10251),
		Pod: NewPod("kube-scheduler-http", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "scheduler",
		}, "< 1.13"),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-shoot-apiserver",
			"allow-to-dns",
			"deny-all",
		),
	}

	// EtcdMainInfo points to cloud-agnostic etcd-main instance.
	EtcdMainInfo = &SourcePod{
		Ports: NewSinglePort(2379),
		Pod: NewPod("etcd-main", labels.Set{
			"app":                     "etcd-statefulset",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "main",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-etcd",
			"allow-to-dns",
			"allow-to-public-networks",
			"allow-to-private-networks",
			"deny-all",
		),
	}

	// EtcdMainInfo points to cloud-agnostic etcd-main instance.
	EtcdEventsInfo = &SourcePod{
		Ports: NewSinglePort(2379),
		Pod: NewPod("etcd-events", labels.Set{
			"app":                     "etcd-statefulset",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "events",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-etcd",
			"allow-to-dns",
			"allow-to-public-networks",
			"allow-to-private-networks",
			"deny-all",
		),
	}

	// CloudControllerManagerInfoNotSecured points to cloud-agnostic cloud-controller-manager running on HTTP port.
	CloudControllerManagerInfoNotSecured = &SourcePod{
		Ports: NewSinglePort(10253),
		Pod: NewPod("cloud-controller-manager-http", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "cloud-controller-manager",
		}, "< 1.13"),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-shoot-apiserver",
			"allow-to-dns",
			"allow-to-public-networks",
			"deny-all",
		),
	}

	// CloudControllerManagerInfoSecured points to cloud-agnostic cloud-controller-manager running on HTTPS port.
	CloudControllerManagerInfoSecured = &SourcePod{
		Ports: NewSinglePort(10258),
		Pod: NewPod("cloud-controller-manager-https", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "cloud-controller-manager",
		}, ">= 1.13"),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-shoot-apiserver",
			"allow-to-dns",
			"allow-to-public-networks",
			"deny-all",
		),
	}

	// ElasticSearchInfo points to cloud-agnostic elasticsearch instance.
	ElasticSearchInfo = &SourcePod{
		Ports: []Port{
			{Name: "http", Port: 9200},
			{Name: "metrics", Port: 9114},
		},
		Pod: NewPod("elasticsearch-logging", labels.Set{
			"app":                     "elasticsearch-logging",
			"garden.sapcloud.io/role": "logging",
			"role":                    "logging",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-elasticsearch",
			"deny-all",
		),
	}

	// GrafanaInfo points to cloud-agnostic grafana instance.
	GrafanaInfo = &SourcePod{
		Ports: NewSinglePort(3000),
		Pod: NewPod("grafana", labels.Set{
			"component":               "grafana",
			"garden.sapcloud.io/role": "monitoring",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-grafana",
			"allow-to-dns",
			"deny-all",
		),
	}

	// KibanaInfo points to cloud-agnostic kibana instance.
	KibanaInfo = &SourcePod{
		Ports: NewSinglePort(5601),
		Pod: NewPod("kibana-logging", labels.Set{
			"app":                     "kibana-logging",
			"garden.sapcloud.io/role": "logging",
			"role":                    "logging",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-kibana",
			"allow-to-dns",
			"allow-to-elasticsearch",
			"deny-all",
		),
	}

	// KubeStateMetricsSeedInfo points to cloud-agnostic kube-state-metrics-seed instance.
	KubeStateMetricsSeedInfo = &SourcePod{
		Ports: NewSinglePort(8080),
		Pod: NewPod("kube-state-metrics-seed", labels.Set{
			"component":               "kube-state-metrics",
			"garden.sapcloud.io/role": "monitoring",
			"type":                    "seed",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-dns",
			"allow-to-seed-apiserver",
			"deny-all",
		),
	}

	// KubeStateMetricsShootInfo points to cloud-agnostic kube-state-metrics-shoot instance.
	KubeStateMetricsShootInfo = &SourcePod{
		Ports: NewSinglePort(8080),
		Pod: NewPod("kube-state-metrics-shoot", labels.Set{
			"component":               "kube-state-metrics",
			"garden.sapcloud.io/role": "monitoring",
			"type":                    "shoot",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-dns",
			"allow-to-shoot-apiserver",
			"deny-all",
		),
	}

	// MachineControllerManagerInfo points to cloud-agnostic machine-controller-manager instance.
	MachineControllerManagerInfo = &SourcePod{
		Ports: NewSinglePort(10258),
		Pod: NewPod("machine-controller-manager", labels.Set{
			"app":                     "kubernetes",
			"garden.sapcloud.io/role": "controlplane",
			"role":                    "machine-controller-manager",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-from-prometheus",
			"allow-to-dns",
			"allow-to-public-networks",
			"allow-to-private-networks",
			"allow-to-seed-apiserver",
			"allow-to-shoot-apiserver",
			"deny-all",
		),
	}

	// PrometheusInfo points to cloud-agnostic prometheus instance.
	PrometheusInfo = &SourcePod{
		Ports: NewSinglePort(9090),
		Pod: NewPod("prometheus", labels.Set{
			"app":                     "prometheus",
			"garden.sapcloud.io/role": "monitoring",
			"role":                    "monitoring",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-prometheus",
			"allow-to-dns",
			"allow-to-public-networks",
			"allow-to-seed-apiserver",
			"allow-to-shoot-apiserver",
			"allow-to-shoot-networks",
			"deny-all",
		),
	}

	// AddonManagerInfo points to gardener-resource-manager instance.
	AddonManagerInfo = &SourcePod{
		Pod: NewPod("gardener-resource-manager", labels.Set{
			"app":                     "gardener-resource-manager",
			"garden.sapcloud.io/role": "controlplane",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-to-dns",
			"allow-to-shoot-apiserver",
			"deny-all",
		),
	}

	// DependencyWatchdog points to dependency-watchdog instance.
	DependencyWatchdog = &SourcePod{
		Pod: NewPod("dependency-watchdog", labels.Set{
			"role": "dependency-watchdog",
		}),
		ExpectedPolicies: sets.NewString(
			"allow-to-dns",
			"allow-to-seed-apiserver",
			"deny-all",
		),
	}

	// AddonManagerInfo points to busybox instance.
	BusyboxInfo = &SourcePod{
		Pod: NewPod("busybox", labels.Set{
			"app":  "busybox",
			"role": "testing",
		}),
	}

	// ExternalHost points external host.
	ExternalHost = &Host{
		Description: "External host",
		HostName:    "8.8.8.8",
		Port:        53,
	}

	// SeedKubeAPIServer points the Seed Kube APIServer.
	SeedKubeAPIServer = &Host{
		Description: "Seed Kube APIServer",
		HostName:    "kubernetes.default",
		Port:        443,
	}

	// GardenPrometheus points the Gardener Prometheus running in the seed cluster.
	GardenPrometheus = &Host{
		Description: "Garden Prometheus",
		HostName:    "prometheus-web.garden",
		Port:        80,
	}
)
