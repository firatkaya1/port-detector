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

	m.AddPage()

	m.Row(200, func() {
		m.Col(12, func() {
			m.Text("Frequently Asked Questions", props.Text{
				Top:    5,
				Size:   20,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
			m.Text("1-Why open ports has vulnerability?", props.Text{
				Top:    20,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("We believe that open ports can be under attack from outside attack to hack your server or personal computer.", props.Text{
				Top:    30,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("2-Why open ports has vulnerability?", props.Text{
				Top:    45,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("We believe that open ports can be under attack from outside attack to hack your server or personal computer.", props.Text{
				Top:    55,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("2-Why open ports has vulnerability?", props.Text{
				Top:    70,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("We believe that open ports can be under attack from outside attack to hack your server or personal computer.", props.Text{
				Top:    80,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text("2-Why open ports has vulnerability?", props.Text{
				Top:    95,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text("We believe that open ports can be under attack from outside attack to hack your server or personal computer.", props.Text{
				Top:    105,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
		})
	})

	err := m.OutputFileAndClose("/home/kaya/Music/test.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
