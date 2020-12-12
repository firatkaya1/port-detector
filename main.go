package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

const dot = "\xE2\x80\xA2"

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(false)

	m.Row(60, func() {
		m.Col(12, func() {
			m.FileImage("/home/kaya/go/src/port-detector/assets/port.png", props.Rect{
				Percent: 100,
				Center:  true,
			})
		})
	})
	m.Row(50, func() {
		m.Col(12, func() {
			m.Text("PORT DETECTOR", props.Text{
				Top:   30,
				Size:  25,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("\"Find Listening ports easily\"", props.Text{
				Top:   30,
				Size:  18,
				Style: consts.Italic,
				Align: consts.Center,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("Created By Firat Kaya", props.Text{
				Top:   30,
				Size:  13,
				Style: consts.Italic,
				Align: consts.Center,
			})
		})
	})

	//License Page
	m.AddPage()
	m.Row(200, func() {
		m.Col(12, func() {
			m.Text("Apache License", props.Text{
				Top:   95,
				Size:  18,
				Style: consts.Bold,
				Align: consts.Center,
			})
			m.Text("Version 2.0, January 2004", props.Text{
				Top:   115,
				Size:  18,
				Style: consts.Bold,
				Align: consts.Center,
			})
			m.Text("Source Codes", props.Text{
				Top:   140,
				Size:  10,
				Style: consts.Italic,
				Align: consts.Center,
			})
			m.Text("https://github.com/firatkaya1/port-detector ", props.Text{
				Top:   150,
				Size:  10,
				Style: consts.BoldItalic,
				Align: consts.Center,
			})
		})
	})
	//Operation System will be here
	m.AddPage()
	m.Row(50, func() {
		m.Col(12, func() {
			m.FileImage("/home/kaya/go/src/port-detector/assets/ubuntu.png", props.Rect{
				Percent: 100,
				Center:  true,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("Ubuntu", props.Text{
				Top:         10,
				Size:        30,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Center,
			})
		})

	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Version 19.10", props.Text{
				Size:        30,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Center,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("We are detect to your operation system is Ubuntu 19.10. Under the these page all suggestions re-write for you personal computer. Feel free", props.Text{
				Top:         10,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Normal,
				Align:       consts.Left,
			})
		})
	})
	m.Row(120, func() {
		m.Col(12, func() {
			m.Text("Please before the start check out dependencies for firewall settings. If you do not have these dependencies, open the original documentation(github.com/firatkaya1) and install to your operation system.", props.Text{
				Top:         20,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Normal,
				Align:       consts.Left,
			})
			m.Text(dot+" We are using iptables command on ubuntu. ", props.Text{
				Top:         40,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text(dot+" sudo apt-get update && sudo apt-get install iptables ", props.Text{
				Top:         50,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text("Settings ", props.Text{
				Top:         60,
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})

			m.Text("Extract Path ", props.Text{
				Top:         70,
				Size:        13,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text("/home/kaya/test.pdf ", props.Text{
				Top:         78,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text("Language ", props.Text{
				Top:         90,
				Size:        13,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text("English ", props.Text{
				Top:         98,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
		})

	})

	//Add the open ports
	m.AddPage()
	headers := []string{"#", "Host", "Port", "Status", "?"}
	contents := [][]string{
		{"1", "0.0.0.0", "3306", "Listening", "MYSQL"},
		{"2", "0.0.0.0", "8080", "Listening", "Apache Tomcat"},
		{"3", "0.0.0.0", "443", "Listening", "Apache Server"},
		{"4", "0.0.0.0", "69", "Listening", "Java"},
		{"5", "0.0.0.0", "56", "Listening", "Java"},
		{"6", "0.0.0.0", "4453", "Listening", "MYSQL"},
		{"7", "0.0.0.0", "22", "Listening", "Dovecot"},
	}
	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      17.0,
			GridSizes: []uint{1, 2, 3, 3, 3},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      15.0,
			GridSizes: []uint{1, 2, 3, 3, 3},
		},
		Align: consts.Center,
		AlternatedBackground: &color.Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		},
		HeaderContentSpace: 10.0,
		Line:               true,
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("We are detecting 11 ports and 3 port came from unknown, if you are considering these ports not expected, please forbid these ports.", props.Text{
				Top:         5,
				Size:        12,
				Extrapolate: false,
				Style:       consts.BoldItalic,
				Align:       consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Why should you close unused ports?", props.Text{
				Top:         20,
				Size:        15,
				Extrapolate: false,
				Style:       consts.BoldItalic,
				Align:       consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Open ports on a server are a security vulnerability that can "+
				"potentially allow a hacker to exploit services on your network. If those services "+
				"are unpatched, a hacker can easily take advantage of the system by running a simple "+
				"port scan using free software like nmap to discover the open ports. It’s important that you "+
				"understand some basics about port security and how to manage ports with the principal of least privilege.", props.Text{
				Top:         15,
				Size:        12,
				Family:      consts.Arial,
				Extrapolate: false,
				Style:       consts.Normal,
				Align:       consts.Left,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("How to block  a port?", props.Text{
				Top:         40,
				Size:        15,
				Extrapolate: false,
				Style:       consts.BoldItalic,
				Align:       consts.Left,
			})
		})
	})
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text("In linux worlds, we recommend to you use commands which is help you to understand and managed your firewall. ", props.Text{
				Top:         10,
				Size:        12,
				Family:      consts.Arial,
				Extrapolate: false,
				Style:       consts.Normal,
				Align:       consts.Left,
			})
		})
	})
	m.Row(80, func() {
		m.Col(12, func() {
			m.Text("Step 1", props.Text{
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text("Install iptables library to your machine. ", props.Text{
				Top:         10,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text(dot+" sudo apt-get update && sudo apt-get install iptables ", props.Text{
				Top:         20,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text("Step 2", props.Text{
				Top:         30,
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text("Block the incoming request which is protect to you from unknown resources.", props.Text{
				Top:         40,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text(dot+" sudo iptables -I OUTPUT -p tcp --dport {YOUR-PORT-NUMBER} -j DROP ", props.Text{
				Top:         50,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})

			m.Text("Step 3", props.Text{
				Top:         60,
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text("Check the ports, are they closed from outside request?", props.Text{
				Top:         70,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text(dot+" sudo iptables -L -nv ", props.Text{
				Top:         80,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
		})
	})

	//Frequently Asked Questions
	m.AddPage()
	m.Row(300, func() {
		m.Col(12, func() {
			m.Text("Frequently Asked Questions", props.Text{
				Top:    5,
				Size:   20,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
			m.Text("1-Why I should close open ports?", props.Text{
				Top:    20,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Open ports can accept outside request to your machine and represent vulnerability of your "+
				"information. We are recommended to close outside request for protect itself.", props.Text{
				Top:    30,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("2-Why should some ports be open?", props.Text{
				Top:    45,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Even we want to control all the ports, we have to continue our communication between machine "+
				"and world. For example, we do not want to close 80 port because this port working with HTTP and connect to"+
				"other machine.Another example is 22 port. This port helps to us connect our machine via SSH. ", props.Text{
				Top:    55,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("3-What happen if i close all ports?", props.Text{
				Top:    70,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Definitely worst idea. You can not contact your machine if machine is a server. All communication will lose.", props.Text{
				Top:    80,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("4-How can i close all ports only one command ? ", props.Text{
				Top:    95,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("iptables -A INPUT -p tcp -m tcp -m multiport ! --dports 80,443 -j DROP   ", props.Text{
				Top:    105,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.BoldItalic,
				Align:  consts.Left,
			})
			m.Text("5-How can i open a port that i closed? ", props.Text{
				Top:    120,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Let's we assume, we closed 22 port after that we wants to open same port again.", props.Text{
				Top:    130,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
			m.Text("iptables -I INPUT -p tcp --dport 22 --syn -j ACCEPT ", props.Text{
				Top:    140,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
			m.Text("6-What is iptables command ? ", props.Text{
				Top:    155,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Iptables is a Linux command line firewall that allows system administrators to manage incoming "+
				"and outgoing traffic via a set of configurable table rules. Iptables uses a set of tables which have "+
				"chains that contain set of built-in or user defined rules", props.Text{
				Top:    165,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
			m.Text("6-What is extract path in this document? ", props.Text{
				Top:    180,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Extract path is represent your document where will save. ", props.Text{
				Top:    190,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
			m.Text("7-What is the language? ", props.Text{
				Top:    200,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Language is represent main language of this document. You can change before create this document. ", props.Text{
				Top:    210,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("8-What is the language? ", props.Text{
				Top:    220,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("Language is represent main language of this document. You can change before create this document. ", props.Text{
				Top:    230,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
		})
	})

	m.AddPage()
	commandHeaders := []string{"Command", "Explain"}
	commandContents := [][]string{
		{"-I", "Append this rule to a rule chain. Valid chains for what we're doing are INPUT, FORWARD and OUTPUT, but " +
			"we mostly deal with INPUT in this tutorial, which affects only incoming traffic. "},
		{"OUTPUT", "Represent incoming traffic."},
		{"-p", "The connection protocol used."},
		{"tcp", "Represent Transmission Control Protocol"},
		{"--dport", "The destination port(s) required for this rule. A single port may be given, or a range may be given as start:end, which will match all ports from start to end, inclusive."},
		{"{YOUR-PORT-NUMBER} ", " Which port do you want to block ?"},
		{"-j", "Jump to the specified target. By default, iptables allows four targets:ACCEPT,REJECT,DROP,LOG"},
		{"ACCEPT", "Accept the packet and stop processing rules in this chain."},
		{"REJECT", "Reject the packet and notify the sender that we did so, and stop processing rules in this chain."},
		{"DROP", " Silently ignore the packet, and stop processing rules in this chain."},
		{"LOG", "Log the packet, and continue processing more rules in this chain. Allows the use of the --log-prefix and --log-level options."},
	}
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("9-What is ' sudo iptables -A OUTPUT -p tcp --dport {YOUR-PORT-NUMBER} -j DROP' ? ", props.Text{
				Size:   13,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})
	m.TableList(commandHeaders, commandContents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      16.0,
			GridSizes: []uint{3, 9},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      13.0,
			GridSizes: []uint{3, 9},
		},
		Align: consts.Left,
		AlternatedBackground: &color.Color{
			Red:   235,
			Green: 255,
			Blue:  255,
		},
		HeaderContentSpace: 5.0,
		Line:               true,
	})
	iptableHeaders := []string{"Command", "Explain"}
	iptableContent := [][]string{
		{"-L", "List of current rules in iptables."},
		{"-nv", "Filter only invalid rules."},
	}
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("9-What is ' sudo iptables -L -nv' ? ", props.Text{
				Top:    5,
				Size:   13,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("This command list of all rules in iptables.", props.Text{
				Top:    15,
				Size:   13,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
		})
	})
	m.TableList(iptableHeaders, iptableContent, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      16.0,
			GridSizes: []uint{3, 9},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      13.0,
			GridSizes: []uint{3, 9},
		},
		Align: consts.Left,
		AlternatedBackground: &color.Color{
			Red:   235,
			Green: 255,
			Blue:  255,
		},
		HeaderContentSpace: 5.0,
		Line:               true,
	})

	//Port and Explain
	m.AddPage()
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text("Most Common Ports and Services", props.Text{
				Align: consts.Center,
				Style: consts.Bold,
				Size:  20,
			})
		})
	})
	portHeaders := []string{"Port", "Service", "Explain"}
	portContents := [][]string{
		{"20,21", "FTP", "File Transfer Protocol"},
		{"22", "SSH", "Secure Shell (SSH)"},
		{"23", "TELNET", "Unencrypted text communication"},
		{"25", "SMTP", "Simple Mail Transfer Protocol"},
		{"53", "DNS", "Domain Name System"},
		{"80", "HTTP", "Hyper Text Transfer Protocol"},
		{"110", "POP3", "Post Office Protocol"},
		{"119", "NNTP", "Network News Transfer Protocol"},
		{"123", "NTP", "Network Time Protocol"},
		{"143", "IMAP", "Internet Message Access Protocol"},
		{"443", "HTTPS", "Hyper Text Transfer Protocol Security"},
		{"465", "SMTPS", "Simple Mail Transfer Protocol Over SSL"},
		{"3306", "mysql", "MySQL Database Service"},
		{"5432", "postgres", "PostgreSQL Database"},
		{"8080", "Tomcat", "Apache Tomcat Server"},
	}
	m.TableList(portHeaders, portContents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      16.0,
			GridSizes: []uint{2, 2, 8},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      12.0,
			GridSizes: []uint{2, 2, 8},
		},
		Align: consts.Left,
		AlternatedBackground: &color.Color{
			Red:   255,
			Green: 230,
			Blue:  255,
		},
		HeaderContentSpace: 5.0,
		Line:               true,
	})

	err := m.OutputFileAndClose("/home/kaya/Music/test.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
