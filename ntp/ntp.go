// Copyright 2017 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ntp syncs system time with NTP server.
package ntp

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	DefaultNTP    = "/usr/sbin/ntpdate"
	DefaultServer = "ntp.ntsc.ac.cn"
	// DefaultServer = "ntp.aliyun.com"
	// DefaultServer = "ntp.tencent.com"
	// DefaultServer = "time.windows.com"
	// DefaultServer = "time.google.com"
)

// DefaultSync syncs system time with NTP server.
func DefaultSync() (string, error) {
	return Sync(DefaultNTP, DefaultServer)
}

func DefaultSyncService(switch_timer bool, hours int) {
	DefaultSync()

	if switch_timer {
		if hours < 1 {
			hours = 8
		}
		ticker := time.NewTicker(time.Duration(hours) * time.Hour)
		for {
			select {
			case <-ticker.C:
				DefaultSync()
			}
		}
	}
}

var inUseErr = "NTP socket is in use"

// Sync syncs system time with NTP server.
//
//    sudo service ntp stop
//    sudo ntpdate time.google.com
//    sudo service ntp start
//
func Sync(ntpPath string, server string) (string, error) {
	buf := new(bytes.Buffer)
	err := startNTP(ntpPath, server, buf)
	o := strings.TrimSpace(buf.String())
	if err == nil && strings.Contains(o, "adjust time server") {
		return o, nil
	}

	if !strings.Contains(o, inUseErr) && (err != nil && !strings.Contains(err.Error(), inUseErr)) {
		return o, err
	}

	so, err := serviceNTP("stop")
	if err != nil {
		return so, err
	}

	buf.Reset()
	err = startNTP(ntpPath, server, buf)
	o = strings.TrimSpace(buf.String())
	if err != nil {
		return o, err
	}

	so, err = serviceNTP("start")
	if err != nil {
		return so, err
	}
	if so != "" {
		o += ";" + so
	}
	return o, nil
}

func startNTP(ntpPath string, server string, w io.Writer) error {
	if !exist(ntpPath) {
		return fmt.Errorf("%q does not exist", ntpPath)
	}
	cmd := exec.Command("sudo", ntpPath, server)
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

// when 'NTP socket is in use' error is returned
func serviceNTP(command string) (string, error) {
	buf := new(bytes.Buffer)
	cmd := exec.Command("sudo", "service", "ntp", command)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	return strings.TrimSpace(buf.String()), err
}

// exist returns true if the file or directory exists.
func exist(fpath string) bool {
	st, err := os.Stat(fpath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if st.IsDir() {
		return true
	}
	if _, err := os.Stat(fpath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
