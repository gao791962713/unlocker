// SPDX-FileCopyrightText: © 2014-2022 David Parsons
// SPDX-License-Identifier: MIT

//go:build linux && !esx
// +build linux,!esx

package vmwpatch

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func IsAdmin() bool {
	if os.Geteuid() == 0 {
		return true
	}
	return false
}

//goland:noinspection GoUnusedParameter
func VMWStart(v *VMwareInfo) {
	// Dummy function on Linux
	return
}

//goland:noinspection GoUnusedParameter
func VMWStop(v *VMwareInfo) {
	// Dummy function on Linux
	return
}

func VMWInfo() *VMwareInfo {
	v := &VMwareInfo{}

	// Store known service names
	// Not used on Linux
	v.AuthD = ""
	v.HostD = ""
	v.USBD = ""

	// Access /etc/vmware/config for version, build and installation path
	file, err := os.Open("/etc/vmware/config")
	if err != nil {
		panic(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()
	config := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
					value = trimQuotes(value)
				}
				config[key] = value
			}
		}
	}

	// Basic product settings
	v.ProductVersion = config["product.version"]
	v.BuildNumber = fmt.Sprintf("%s.%s", v.ProductVersion, config["product.buildNumber"])
	v.InstallDir = config["libdir"]

	// Construct needed filenames from reg settings
	v.InstallDir64 = ""
	v.Player = "vmplayer"
	v.Workstation = "vmware"
	v.KVM = "vmware-kvm"
	v.REST = "vmrest"
	v.Tray = "vmware-tray"
	// Linux has no shell extension
	v.ShellExt = ""
	v.VMXDefault = "vmware-vmx"
	v.VMXDebug = "vmware-vmx-debug"
	v.VMXStats = "vmware-vmx-stats"
	v.VMwareBase = "libvmwarebase.so"
	v.PathVMXDefault = filepath.Join(v.InstallDir, "bin", v.VMXDefault)
	v.PathVMXDebug = filepath.Join(v.InstallDir, "bin", v.VMXDebug)
	v.PathVMXStats = filepath.Join(v.InstallDir, "bin", v.VMXStats)
	v.PathVMwareBase = filepath.Join(v.InstallDir, "lib", v.VMwareBase, v.VMwareBase)
	currentFolder, _ := os.Getwd()
	v.BackDir = filepath.Join(currentFolder, "backup", v.ProductVersion)
	v.BackVMXDefault = filepath.Join(v.BackDir, v.VMXDefault)
	v.BackVMXDebug = filepath.Join(v.BackDir, v.VMXDebug)
	v.BackVMXStats = filepath.Join(v.BackDir, v.VMXStats)
	v.BackVMwareBase = filepath.Join(v.BackDir, v.VMwareBase)
	v.PathISOMacOSX = filepath.Join(v.InstallDir, "isoimages", "darwinPre15.iso")
	v.PathISOmacOS = filepath.Join(v.InstallDir, "isoimages", "darwin.iso")
	return v
}

//goland:noinspection GoUnusedParameter
func setCTime(path string, ctime time.Time) error {
	// Dummy function on Linux
	return nil
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' {
			s = s[1:]
		}
		if i := len(s) - 1; s[i] == '"' {
			s = s[:i]
		}
	}
	return s
}
