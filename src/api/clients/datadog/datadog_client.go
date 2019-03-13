package datadog

import (
    "fmt"
    "github.com/DataDog/datadog-go/statsd"
    "github.com/emikohmann/url-shortener/src/api/config"
)

const (
    errSendingMetricToDD   = "error sending metric to datadog"
    applicationNamePattern = "%s."
    applicationNameTag     = "application_name:%s"
)

var (
    client *statsd.Client
)

func init() {
    client, err := statsd.New(config.DatadogMetricAddress)
    if err != nil {
        config.Logger.Println(errSendingMetricToDD, err)
        panic(err)
    }
    // prefix every metric with the app name
    client.Namespace = fmt.Sprintf(applicationNamePattern, config.ApplicationName)
    // send every metric with application name tag
    client.Tags = append(client.Tags, fmt.Sprintf(applicationNameTag, config.ApplicationName))
}

func IncrementSimpleApplicationMetric(metricName string, tags ...string) {
    IncrementApplicationMetric(metricName, 1, tags...)
}

func IncrementApplicationMetric(metricName string, value float64, tags ...string) {
    go func() {
        if err := client.Gauge(metricName, value, tags, 1); err != nil {
            config.Logger.Println(errSendingMetricToDD, err)
            return
        }
    }()
}
