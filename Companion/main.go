package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"text/template"

	"github.com/ZengJinCheng/FicsitRemoteMonitoringCompanion/Companion/exporter"
	"github.com/ZengJinCheng/FicsitRemoteMonitoringCompanion/Companion/prometheus"
	"github.com/ZengJinCheng/FicsitRemoteMonitoringCompanion/Companion/realtime_map"
)

var Version = "0.0.0-dev"

func lookupEnvWithDefault(variable string, defaultVal string) string {
	val, exist := os.LookupEnv(variable)
	if exist {
		return val
	}
	return defaultVal
}

func main() {

	var frmHostname string
	flag.StringVar(&frmHostname, "hostname", "localhost", "hostname of Ficsit Remote Monitoring webserver")
	var frmPort string
	flag.StringVar(&frmPort, "port", "8080", "port of Ficsit Remote Monitoring webserver")

	var frmHostnames string
	flag.StringVar(&frmHostnames, "hostnames", "", "comma separated values of multiple Ficsit Remote Monitoring webservers, of the form http://myserver1:8080,http://myserver2:8080. If defined, this will be used instead of hostname+port")

	var genReadme bool
	flag.BoolVar(&genReadme, "GenerateReadme", false, "Generate readme and exit")
	var noProm bool
	flag.BoolVar(&noProm, "noprom", false, "Do not run prometheus with the app.")
	flag.Parse()

	frmHostname = lookupEnvWithDefault("FRM_HOST", frmHostname)
	frmPort = lookupEnvWithDefault("FRM_PORT", frmPort)
	frmHostnames = lookupEnvWithDefault("FRM_HOSTS", frmHostnames)
	logStdout, _ := strconv.ParseBool(lookupEnvWithDefault("FRM_LOG_STDOUT", "0"))

	if genReadme {
		generateReadme()
		os.Exit(0)
	}

	if !logStdout {
		logFile, err := createLogFile()
		if err != nil {
			fmt.Printf("创建日志文件时出错: %s", err)
			os.Exit(1)
		}
		log.Default().SetOutput(logFile)
	}

	// Create exporter
	frmUrls := []string{}
	if frmHostnames == "" {
		frmUrls = append(frmUrls, "http://"+frmHostname+":"+frmPort)
	} else {
		for _, frmServer := range strings.Split(frmHostnames, ",") {
			if !strings.HasPrefix(frmServer, "http://") && !strings.HasPrefix(frmServer, "https://") {
				frmServer = "http://" + frmServer
			}
			frmUrls = append(frmUrls, frmServer)
		}
	}
	var promExporter *exporter.PrometheusExporter
	promExporter = exporter.NewPrometheusExporter(frmUrls)

	var prom *prometheus.PrometheusWrapper
	var err error
	if !noProm {
		// Create prometheus
		prom, err = prometheus.NewPrometheusWrapper()
		if err != nil {
			fmt.Printf("准备Prometheus时出错: %s", err)
			os.Exit(1)
		}
	}

	// Create map server
	mapServ, err := realtime_map.NewMapServer()
	if err != nil {
		fmt.Printf("动态地图准备错误: %s", err)
		os.Exit(1)
	}

	// Start prometheus
	if !noProm {
		err = prom.Start()
		if err != nil {
			fmt.Printf("启动Prometheus时出错: %s", err)
			os.Exit(1)
		}
	}

	// Start exporter
	promExporter.Start()

	// Start map
	mapServ.Start()

	fmt.Printf(`
Ficsit Remote Monitoring Companion (v%s)

访问实时地图访问:
http://localhost:8000/?frmport=8080

    如果您已经配置了Ficsit远程监控
	要为其web服务器使用8080以外的端口，
	将“frmport”查询字符串参数修改为
	匹配您选择的端口并刷新页面。

要访问Prometheus中的指标，请访问:
http://localhost:9090/

按 Ctrl + C 退出。
`, Version)

	// Wait for an interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Stop the exporter
	err = promExporter.Stop()
	if err != nil {
		fmt.Printf("停止prometheus导出程序出错: %s", err)
	}

	// Stop prometheus
	if !noProm {
		err = prom.Stop()
		if err != nil {
			fmt.Printf("停止prometheus程序出错: %s", err)
		}
	}

	// Stop map
	mapServ.Stop()

	fmt.Println("退出。")
	os.Exit(0)
}

func createLogFile() (*os.File, error) {
	curExePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	curExeDir := filepath.Dir(curExePath)

	if err != nil {
		return nil, err
	}

	return os.Create(path.Join(curExeDir, "frmc.log"))
}

func generateReadme() {

	tpl := template.New("readme.tpl.md")
	tpl.Funcs(template.FuncMap{
		"List": strings.Join,
	})

	tpl = template.Must(tpl.ParseFiles("../readme/readme.tpl.md"))

	err := tpl.Execute(
		os.Stdout,
		struct {
			Metrics []exporter.MetricVectorDetails
		}{
			Metrics: exporter.RegisteredMetricVectors,
		},
	)
	if err != nil {
		fmt.Printf("写自述文件错误: %s", err)
	}
}
