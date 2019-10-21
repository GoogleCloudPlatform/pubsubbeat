/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package esx

import "github.com/vmware/govmomi/vim25/types"

var HostFirewallInfo = types.HostFirewallInfo{
	DynamicData: types.DynamicData{},
	DefaultPolicy: types.HostFirewallDefaultPolicy{
		DynamicData:     types.DynamicData{},
		IncomingBlocked: types.NewBool(true),
		OutgoingBlocked: types.NewBool(true),
	},
	Ruleset: []types.HostFirewallRuleset{
		{
			DynamicData: types.DynamicData{},
			Key:         "CIMHttpServer",
			Label:       "CIM Server",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        5988,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "sfcbd-watchdog",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "CIMHttpsServer",
			Label:       "CIM Secure Server",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        5989,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "sfcbd-watchdog",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "CIMSLP",
			Label:       "CIM SLP",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        427,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        427,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        427,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        427,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "DHCPv6",
			Label:       "DHCPv6",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        547,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        546,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        547,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        546,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "DVFilter",
			Label:       "DVFilter",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        2222,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "DVSSync",
			Label:       "DVSSync",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        8302,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8301,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8301,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8302,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "HBR",
			Label:       "HBR",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        31031,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        44046,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "NFC",
			Label:       "NFC",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        902,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        902,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "WOL",
			Label:       "WOL",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        9,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "activeDirectoryAll",
			Label:       "Active Directory All",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        88,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        88,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        123,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        137,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        139,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        389,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        389,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        445,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        464,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        464,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        3268,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        7476,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        2020,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "cmmds",
			Label:       "Virtual SAN Clustering Service",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        12345,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        23451,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        12345,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        23451,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        12321,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        12321,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "dhcp",
			Label:       "DHCP Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        68,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        68,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "src",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "dns",
			Label:       "DNS Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        53,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        53,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        53,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "esxupdate",
			Label:       "esxupdate",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        443,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "faultTolerance",
			Label:       "Fault Tolerance",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        80,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8300,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8300,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "ftpClient",
			Label:       "FTP Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        21,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        20,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "src",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "gdbserver",
			Label:       "gdbserver",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        1000,
					EndPort:     9999,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        50000,
					EndPort:     50999,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "httpClient",
			Label:       "httpClient",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        80,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        443,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "iSCSI",
			Label:       "Software iSCSI Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        3260,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "iofiltervp",
			Label:       "iofiltervp",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        9080,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "ipfam",
			Label:       "NSX Distributed Logical Router Service",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        6999,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        6999,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "nfs41Client",
			Label:       "nfs41Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        0,
					EndPort:     65535,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "nfsClient",
			Label:       "NFS Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        0,
					EndPort:     65535,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "ntpClient",
			Label:       "NTP Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        123,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "ntpd",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "pvrdma",
			Label:       "pvrdma",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        28250,
					EndPort:     28761,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        28250,
					EndPort:     28761,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "rabbitmqproxy",
			Label:       "rabbitmqproxy",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        5671,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "rdt",
			Label:       "Virtual SAN Transport",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        2233,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        2233,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "remoteSerialPort",
			Label:       "VM serial port connected over network",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        0,
					EndPort:     65535,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        23,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        1024,
					EndPort:     65535,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "snmp",
			Label:       "SNMP Server",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        161,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "snmpd",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "sshClient",
			Label:       "SSH Client",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        22,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "sshServer",
			Label:       "SSH Server",
			Required:    true,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        22,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "syslog",
			Label:       "syslog",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        514,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        514,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        1514,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "updateManager",
			Label:       "vCenter Update Manager",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        80,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        9000,
					EndPort:     9100,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vMotion",
			Label:       "vMotion",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        8000,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8000,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vSPC",
			Label:       "VM serial port connected to vSPC",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        0,
					EndPort:     65535,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vSphereClient",
			Label:       "vSphere Web Client",
			Required:    true,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        902,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        443,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vpxHeartbeats",
			Label:       "VMware vCenter Agent",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        902,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "vpxa",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vsanEncryption",
			Label:       "vsanEncryption",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        0,
					EndPort:     65535,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vsanhealth-multicasttest",
			Label:       "vsanhealth-multicasttest",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        5001,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        5001,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "udp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vsanvp",
			Label:       "vsanvp",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        8080,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
				{
					DynamicData: types.DynamicData{},
					Port:        8080,
					EndPort:     0,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "vvold",
			Label:       "vvold",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        0,
					EndPort:     65535,
					Direction:   "outbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: false,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
		{
			DynamicData: types.DynamicData{},
			Key:         "webAccess",
			Label:       "vSphere Web Access",
			Required:    false,
			Rule: []types.HostFirewallRule{
				{
					DynamicData: types.DynamicData{},
					Port:        80,
					EndPort:     0,
					Direction:   "inbound",
					PortType:    "dst",
					Protocol:    "tcp",
				},
			},
			Service: "",
			Enabled: true,
			AllowedHosts: &types.HostFirewallRulesetIpList{
				DynamicData: types.DynamicData{},
				IpAddress:   nil,
				IpNetwork:   nil,
				AllIp:       true,
			},
		},
	},
}
