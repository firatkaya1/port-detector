package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/matishsiao/goInfo"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Best port detector CLI",
	Long: `port-detector an command based interface which is help you to find opened ports and 
report them easily. Start to using, just enter the detect command and watch the magic. Enjoy...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start to detecting...")
		scanPorts()
	},
}
var language string
var path string
var fileName string
var unknownPorts int

func init() {
	rootCmd.AddCommand(detectCmd)
	detectCmd.PersistentFlags().StringVarP(&language, "language", "l", "en", "required")
	detectCmd.PersistentFlags().StringVarP(&path, "path", "p", "/", "required")
	detectCmd.PersistentFlags().StringVarP(&fileName, "name", "n", "port-detector", "required")

	detectCmd.MarkFlagRequired("lang")
	detectCmd.MarkFlagRequired("path")
	detectCmd.MarkFlagRequired("name")

}

func scanPorts() {
	var ports []string
	for i := 0; i < 65532; i++ {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("0.0.0.0", strconv.Itoa(i)), timeout)
		if err != nil {
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Port Detect :", net.JoinHostPort("0.0.0.0", strconv.Itoa(i)))

			ports = append(ports, net.JoinHostPort("0.0.0.0", strconv.Itoa(i))[8:])
		}
	}
	createPDF(ports)
}

const dot = "\xE2\x80\xA2"

var langMap = make(map[string]string)

func createPDF(ports []string) {
	langMap = jsonConsume(langMap)

	osInfos := getOperationSystem()

	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(false)
	m.Row(60, func() {
		m.Col(12, func() {
			m.FileImage("assets/port.png", props.Rect{
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
			m.Text(langMap["p1_motto"], props.Text{
				Top:   30,
				Size:  18,
				Style: consts.Italic,
				Align: consts.Center,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text(langMap["p1_creator"], props.Text{
				Top:   30,
				Size:  13,
				Style: consts.Italic,
				Align: consts.Center,
			})
		})
	})
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
			m.Text(langMap["p2_source_codes"], props.Text{
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
	m.AddPage()
	m.Row(50, func() {
		m.Col(12, func() {
			m.FileImage(getOsLogo(osInfos.GoOS), props.Rect{
				Percent: 100,
				Center:  true,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text(osInfos.OS, props.Text{
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
			m.Text(fmt.Sprintf(langMap["p3_os_info"], osInfos.OS, string(osInfos.CPUs)), props.Text{
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
			m.Text(langMap["p3_os_sub_info"], props.Text{
				Top:         20,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Normal,
				Align:       consts.Left,
			})
			m.Text(fmt.Sprintf(dot+langMap["p3_deb_info"], osInfos.GoOS), props.Text{
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
			m.Text(langMap["p3_settings"], props.Text{
				Top:         60,
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})

			m.Text(langMap["p3_path"], props.Text{
				Top:         70,
				Size:        13,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text(getPath(), props.Text{
				Top:         78,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text(langMap["p3_lang"], props.Text{
				Top:         90,
				Size:        13,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text(getLanguageFormat(), props.Text{
				Top:         98,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
			m.Text(langMap["p3_created_date"], props.Text{
				Top:         108,
				Size:        13,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text(time.Now().Format("2006/01/02 15:04:05"), props.Text{
				Top:         116,
				Size:        12,
				Extrapolate: false,
				Style:       consts.Italic,
				Align:       consts.Left,
			})
		})

	})
	//Add the open ports
	m.AddPage()
	headers := []string{"#", "Host", "Port", langMap["p4_status"], "?"}
	var contents = make([][]string, 1, len(ports))
	for i := range ports {
		contents = append(contents, []string{strconv.Itoa(i + 1), "0.0.0.0", ports[i], langMap["listen"], getPortService(ports[i])})
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
			m.Text(fmt.Sprintf(langMap["p4_port_detect"], len(ports), unknownPorts), props.Text{
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
			m.Text(langMap["p4_question_1"], props.Text{
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
			m.Text(langMap["p4_answer_1"], props.Text{
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
			m.Text(langMap["p4_question_2"], props.Text{
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
			m.Text(langMap["p4_answer_2"], props.Text{
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
			m.Text(langMap["p4_step_1"], props.Text{
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text(langMap["p4_step_1_info"], props.Text{
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
			m.Text(langMap["p4_step_2"], props.Text{
				Top:         30,
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text(langMap["p4_step_2_info"], props.Text{
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

			m.Text(langMap["p4_step_3"], props.Text{
				Top:         60,
				Size:        18,
				Extrapolate: false,
				Style:       consts.Bold,
				Align:       consts.Left,
			})
			m.Text(langMap["p4_step_3_info"], props.Text{
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
			m.Text(langMap["p5_title"], props.Text{
				Top:    5,
				Size:   20,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
			m.Text(langMap["p5_question_1"], props.Text{
				Top:    20,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_1"], props.Text{
				Top:    30,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text(langMap["p5_question_2"], props.Text{
				Top:    45,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_2"], props.Text{
				Top:    55,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text(langMap["p5_question_3"], props.Text{
				Top:    70,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_3"], props.Text{
				Top:    80,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text(langMap["p5_question_4"], props.Text{
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
			m.Text(langMap["p5_question_5"], props.Text{
				Top:    120,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_5"], props.Text{
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
			m.Text(langMap["p5_question_6"], props.Text{
				Top:    155,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_6"], props.Text{
				Top:    165,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_question_7"], props.Text{
				Top:    180,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_7"], props.Text{
				Top:    190,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_question_8"], props.Text{
				Top:    200,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_8"], props.Text{
				Top:    210,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})

			m.Text(langMap["p5_question_9"], props.Text{
				Top:    220,
				Size:   15,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_9"], props.Text{
				Top:    230,
				Size:   12,
				Family: consts.Arial,
				Style:  consts.Normal,
				Align:  consts.Left,
			})
		})
	})
	m.AddPage()
	commandHeaders := []string{langMap["p5_answer_10_command"], langMap["p5_answer_10_explain"]}
	commandContents := [][]string{
		{"-I", langMap["p5_answer_10_1"]},
		{"OUTPUT", langMap["p5_answer_10_2"]},
		{"-p", langMap["p5_answer_10_3"]},
		{"tcp", langMap["p5_answer_10_4"]},
		{"--dport", langMap["p5_answer_10_5"]},
		{"{YOUR-PORT-NUMBER} ", langMap["p5_answer_10_6"]},
		{"-j", langMap["p5_answer_10_7"]},
		{"ACCEPT", langMap["p5_answer_10_8"]},
		{"REJECT", langMap["p5_answer_10_9"]},
		{"DROP", langMap["p5_answer_10_10"]},
		{"LOG", langMap["p5_answer_10_11"]},
	}
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(langMap["p5_question_11"], props.Text{
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
	iptableHeaders := []string{langMap["p5_answer_10_command"], langMap["p5_answer_10_explain"]}
	iptableContent := [][]string{
		{"-L", langMap["p5_answer_11_1"]},
		{"-nv", langMap["p5_answer_11_2"]},
	}
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text(langMap["p5_question_11"], props.Text{
				Top:    5,
				Size:   13,
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
			m.Text(langMap["p5_answer_11"], props.Text{
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
			m.Text(langMap["p6_title"], props.Text{
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

	err := m.OutputFileAndClose(getPath())
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	printTerminal(end, begin)
}

func jsonConsume(mm map[string]string) map[string]string {
	var result map[string]interface{}
	file, _ := ioutil.ReadFile("language/" + getLanguage() + ".json")
	json.Unmarshal([]byte(file), &result)
	for k, v := range result {
		mm[k] = fmt.Sprint(v)
	}
	return mm
}

func printTerminal(end, begin time.Time) {
	if language == "en" {
		fmt.Println("Language :", getLanguageFormat())
		fmt.Println("Extract Path : " + path)
		fmt.Println("File Name : " + fileName)
		fmt.Println("Exact path : " + getPath())
		fmt.Println("Created report time : ", end.Sub(begin))

	} else {
		fmt.Println("Dil :", getLanguageFormat())
		fmt.Println("Dosya'nın çıkarıldığı yer : " + path)
		fmt.Println("Dosya Adı : " + fileName)
		fmt.Println("Tam adresi : " + getPath())
		fmt.Println("Rapor oluşturma süresi : ", end.Sub(begin))
	}

}

func getLanguage() string {
	if strings.Contains(language, "en") {
		return "en"
	} else {
		return "tr"
	}
}

func getPath() string {
	if path[len(path)-1:] == "/" {
		return path + fileName
	} else {
		return path + "/" + fileName + ".pdf"
	}
}

func getOperationSystem() *goInfo.GoInfoObject {
	return goInfo.GetInfo().VarDump()
}
func getOsLogo(osName string) string {
	switch osName {
	case "windows":
		return "assets/microsoft.png"
	case "darwin":
		return "assets/macos.png"
	default:
		return "assets/ubuntu.png"
	}
}

var portServiceMap = map[string]string{
	"20":   "FTP",
	"21":   "FTP",
	"22":   "SSH",
	"23":   "TELNET",
	"25":   "SMTP",
	"53":   "DNS",
	"80":   "HTTP",
	"110":  "POP3",
	"119":  "NNTP",
	"123":  "NTP",
	"143":  "IMAP",
	"443":  "HTTPS",
	"465":  "SMTPS",
	"631":  "IPP",
	"3306": "mySQL",
	"5432": "postgres",
	"5939": "TCP",
	"8080": "Apache Tomcat",
}

func getPortService(portService string) string {
	if _, ok := portServiceMap[portService]; ok {
		return portServiceMap[portService]
	}
	unknownPorts = unknownPorts + 1
	return langMap["unknown"]
}
func getLanguageFormat() string {
	if strings.Contains(language, "en") {
		return "English"
	} else {
		return "Turkish"
	}
}
