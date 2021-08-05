package kloglogger_test

import (
	"bytes"
	"strings"
	"testing"

	kloglogger "github.com/ovirt/go-ovirt-client-log-klog"
	"k8s.io/klog/v2"
)

func TestVerbose(t *testing.T) {
	output := &bytes.Buffer{}
	klog.SetOutput(output)
	logger := kloglogger.NewVerbose(
		klog.V(4),
		klog.V(3),
		klog.V(2),
		klog.V(1),
	)
	klog.V(1).Enabled()
	logger.Debugf("Testing...")
	if lines := len(strings.Split(output.String(), "\n")); lines != 1 {
		t.Fatalf("Incorrect number of log messages in log output (%d lines) during debug test", lines)
	}
	logger.Infof("Testing...")
	if lines := len(strings.Split(output.String(), "\n")); lines != 1 {
		t.Fatalf("Incorrect number of log messages in log output (%d lines) during info test", lines)
	}
	logger.Warningf("Testing...")
	if lines := len(strings.Split(output.String(), "\n")); lines != 1 {
		t.Fatalf("Incorrect number of log messages in log output (%d lines) during warning test", lines)
	}
	logger.Errorf("Testing...")
	if lines := len(strings.Split(output.String(), "\n")); lines != 1 {
		t.Fatalf("Incorrect number of log messages in log output (%d lines) during error test", lines)
	}
}
