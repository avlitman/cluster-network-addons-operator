package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring"
)

const (
	opening = `# Cluster Network Addons Operator metrics
> This file is auto generated by metricsdocs, run 'make generate-doc' in order to update it.

This document aims to help users that are not familiar with metrics exposed by the Cluster Network Addons Operator.
All metrics documented here are auto-generated by the utility tool 'tools/metricsdocs' and reflects exactly what is being exposed.
## Cluster Network Addons Operator Metrics List
`

	footer = `## Developing new metrics
After developing new metrics or changing old ones, please run 'make generate-doc' to regenerate this document.`
)

func main() {
	metricsList := metricsOptsToMetricList(monitoring.MetricsOptsList)
	sort.Sort(metricsList)
	writeToFile(metricsList)
}

func writeToFile(metricsList metricList) {
	fmt.Print(opening)
	metricsList.writeOut()
	fmt.Print(footer)
}

type metric struct {
	name        string
	description string
}

func metricsOptsToMetricList(mdl map[monitoring.MetricsKey]monitoring.MetricsOpts) metricList {
	res := make([]metric, 0)
	for _, element := range mdl {
		res = append(res, metricDescriptionToMetric(element))
	}

	return res
}

func metricDescriptionToMetric(rrd monitoring.MetricsOpts) metric {
	return metric{
		name:        rrd.Name,
		description: rrd.Help,
	}
}

func (m metric) writeOut() {
	fmt.Println("###", m.name)
	fmt.Println(m.description)
}

type metricList []metric

// Len implements sort.Interface.Len
func (m metricList) Len() int {
	return len(m)
}

// Less implements sort.Interface.Less
func (m metricList) Less(i, j int) bool {
	return m[i].name < m[j].name
}

// Swap implements sort.Interface.Swap
func (m metricList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *metricList) add(line string) {
	split := strings.Split(line, " ")
	name := split[2]
	split[3] = strings.Title(split[3])
	description := strings.Join(split[3:], " ")
	*m = append(*m, metric{name: name, description: description})
}

func (m metricList) writeOut() {
	for _, met := range m {
		met.writeOut()
	}
}
