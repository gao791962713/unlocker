// SPDX-FileCopyrightText: © 2014-2022 David Parsons
// SPDX-License-Identifier: MIT

//go:build linux && esx
// +build linux,esx

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
	// Dummy function on ESXi
	return true
}

//goland:noinspection GoUnusedParameter
func VMWStart(v *VMwareInfo) {
	// Dummy function on ESXi
	return
}

//goland:noinspection GoUnusedParameter
func VMWStop(v *VMwareInfo) {
	// Dummy function on ESXi
	return
}

func VMWInfo() *VMwareInfo {
	v := &VMwareInfo{}

	// Store known service names
	// Not used on ESXi
	v.AuthD = ""
	v.HostD = ""
	v.USBD = ""

	// ESXi command for version esxcli --formatter=keyvalue system version get
	// File /etc/vmware/.buildInfo
	// GITHASH:0
	// CHANGE:9235482
	// BRANCH:esx-7.0.3
	// UID:201
	// VMTREE:/build/mts/release/bora-18644231/bora
	// VMBLD:release
	// BUILDTAG:gobuild
	// BUILDNUMBER:18644231
	// STAGEPATH:build/esx/release

	file, err := os.Open("/etc/vmware/.buildInfo")
	if err != nil {
		panic(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()
	config := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, ":"); equal >= 0 {
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
	v.ProductVersion = config["BRANCH"]
	v.BuildNumber = fmt.Sprintf("%s.%s", v.ProductVersion, config["BUILDNUMBER"])
	v.InstallDir = ""

	// ESXi has only VMX executables to patch
	v.InstallDir64 = ""
	v.Player = ""
	v.Workstation = ""
	v.KVM = ""
	v.REST = ""
	v.Tray = ""
	v.ShellExt = ""
	v.VMXDefault = "vmx"
	v.VMXDebug = "vmx-debug"
	v.VMXStats = "vmx-stats"
	v.VMwareBase = ""
	v.PathVMXDefault = filepath.Join("/bin", v.VMXDefault)
	v.PathVMXDebug = filepath.Join("/bin", v.VMXDebug)
	v.PathVMXStats = filepath.Join("/bin", v.VMXStats)
	v.PathVMwareBase = ""
	currentFolder, _ := os.Getwd()
	v.BackDir = filepath.Join(currentFolder, "backup", v.ProductVersion)
	v.BackVMXDefault = ""
	v.BackVMXDebug = ""
	v.BackVMXStats = ""
	v.BackVMwareBase = ""
	v.PathISOMacOSX = ""
	v.PathISOmacOS = ""
	return v
}

//goland:noinspection GoUnusedParameter
func setCTime(path string, ctime time.Time) error {
	// Dummy function on ESXi
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
