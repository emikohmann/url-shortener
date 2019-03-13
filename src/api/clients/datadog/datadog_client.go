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

func IncrementSimpleApplicationMetric(metricName string, tags ...string) {
    IncrementApplicationMetric(metricName, 1, tags...)
}

func IncrementApplicationMetric(metricName string, value float64, tags ...string) {
    go func() {
        c, err := statsd.New(config.DatadogMetricAddress)
        if err != nil {
            config.Logger.Println(errSendingMetricToDD, err)
            return
        }
        // prefix every metric with the app name
        c.Namespace = fmt.Sprintf(applicationNamePattern, config.ApplicationName)
        // send every metric with application name tag
        c.Tags = append(c.Tags, fmt.Sprintf(applicationNameTag, config.ApplicationName))
        if err = c.Gauge(metricName, value, tags, 1); err != nil {
            config.Logger.Println(errSendingMetricToDD, err)
            return
        }
    }()
}
