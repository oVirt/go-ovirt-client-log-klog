package kloglogger_test

import (
	"bytes"
	"strings"
	"testing"

	kloglogger "github.com/ovirt/go-ovirt-client-log-klog"
	"k8s.io/klog/v2"
)

func TestDebugf(t *testing.T) {
	output := &bytes.Buffer{}
	klog.SetOutput(output)
	logger := kloglogger.New()

	logger.Debugf("Testing...")

	if !strings.Contains("Testing...", output.String()) {
		t.Fatalf("Output did not contain written string: %s", output.String())
	}
}

func TestInfof(t *testing.T) {
	output := &bytes.Buffer{}
	klog.SetOutput(output)
	logger := kloglogger.New()

	logger.Infof("Testing...")

	if !strings.Contains("Testing...", output.String()) {
		t.Fatalf("Output did not contain written string: %s", output.String())
	}
}

func TestWarningf(t *testing.T) {
	output := &bytes.Buffer{}
	klog.SetOutput(output)
	logger := kloglogger.New()

	logger.Warningf("Testing...")

	if !strings.Contains("Testing...", output.String()) {
		t.Fatalf("Output did not contain written string: %s", output.String())
	}
}

func TestErrorf(t *testing.T) {
	output := &bytes.Buffer{}
	klog.SetOutput(output)
	logger := kloglogger.New()

	logger.Errorf("Testing...")

	if !strings.Contains("Testing...", output.String()) {
		t.Fatalf("Output did not contain written string: %s", output.String())
	}
}
