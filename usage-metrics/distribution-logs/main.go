package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type format struct {
	DateTime      time.Time `json:"date_time"`
	Method        string    `json:"method"`
	Status        string    `json:"status"`
	UserAgent     string    `json:"user_agent"`
	URI           string    `json:"uri"`
	DriverVersion string    `json:"driver_version"`
	Target        string    `json:"target"`
	KernelRelease string    `json:"kernel_release"`
	KernelVersion string    `json:"kernel_version"`
	DriverType    string    `json:"driver_type"`
	SourceIP      net.IP    `json:"source_ip"`
}

func driverTypeFromExt(ext string) string {
	switch ext {
	case ".ko":
		return "module"
	case ".o":
		return "bpf"
	}
	return "unknown"
}
func extractDriverInfo(line string) *format {

	dets := strings.Split(line, "\t")

	if len(dets) < 10 {
		logrus.Debug("skipping because details are not available")
		return nil
	}
	dateStr := dets[0]
	timeStr := dets[1]
	methodStr := dets[5]
	urlStr := dets[7]
	statusStr := dets[8]
	userAgentStr := dets[10]
	sourceIP := dets[4]

	finalTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", dateStr, timeStr))
	if err != nil {
		logrus.Debug("skipping because error parsing date: %v", err)
		return nil
	}

	res := strings.TrimPrefix(urlStr, "/driver/")

	res = strings.TrimLeft(res, "/")

	driverSplit := strings.Split(res, "/")
	if len(driverSplit) != 2 {
		logrus.Debug("driver split expects two slices")
		return nil
	}

	driverVersion := driverSplit[0]

	driverFileName := driverSplit[1]

	extension := filepath.Ext(driverFileName)

	basename := strings.TrimSuffix(driverFileName, extension)

	distroSplit := strings.Split(basename, "_")

	if len(distroSplit) < 4 {
		logrus.Debug("skipping because distribution details must be four")
		return nil
	}

	target := distroSplit[1]
	kernelRelease := distroSplit[2]
	kernelVersion := distroSplit[3]
	return &format{
		DateTime:      finalTime,
		Method:        methodStr,
		Status:        statusStr,
		URI:           urlStr,
		UserAgent:     userAgentStr,
		DriverVersion: driverVersion,
		Target:        target,
		KernelRelease: kernelRelease,
		KernelVersion: kernelVersion,
		DriverType:    driverTypeFromExt(extension),
		SourceIP:      net.ParseIP(sourceIP),
	}
}

func main() {
	file, err := os.Open("_output/all.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dinfos := []*format{}
	for scanner.Scan() {
		line := scanner.Text()

		curf := extractDriverInfo(line)

		dinfos = append(dinfos, curf)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ts := targetStats(dinfos)
	printTargetStatsCSV(ts)
}

// must contain comparable types only
type aggregateTarget struct {
	DriverVersion string `json:"driver_version"`
	Target        string `json:"target"`
	KernelRelease string `json:"kernel_release"`
	KernelVersion string `json:"kernel_version"`
	DriverType    string `json:"driver_type"`
	SourceIP      string `json:"source_ip"`
	Stats         *targetCounter
}

func makeAggregateTarget(f *format) *aggregateTarget {
	return &aggregateTarget{
		DriverVersion: f.DriverVersion,
		Target:        f.Target,
		KernelVersion: f.KernelVersion,
		KernelRelease: f.KernelRelease,
		DriverType:    f.DriverType,
		SourceIP:      f.SourceIP.String(),
		Stats:         nil,
	}
}

type targetCounter struct {
	Statuses map[string]uint
}

func targetStats(dinfos []*format) []*aggregateTarget {

	index := map[aggregateTarget]*targetCounter{}

	for _, dinfo := range dinfos {

		// Aggregated item init and lookup
		if dinfo == nil {
			continue
		}

		aggregate := makeAggregateTarget(dinfo)
		counter, ok := index[*aggregate]

		if !ok {
			counter = &targetCounter{
				Statuses: make(map[string]uint),
			}
			index[*aggregate] = counter
		}

		// Upsert counters
		counter.Statuses[dinfo.Status]++
	}

	// Join and convert to array
	res := make([]*aggregateTarget, 0, len(index))
	for a, c := range index {
		agg := a
		agg.Stats = c
		res = append(res, &agg)
	}

	return res
}

func printTargetStatsCSV(ts []*aggregateTarget) {
	w := csv.NewWriter(os.Stdout)
	header := []string{
		"driverversion",
		"target",
		"kernelrelease",
		"kernelversion",
		"drivertype",
		"sourceip",
		"200",
		"2xx",
		"3xx",
		"404",
		"4xx",
		"5xx",
		"unknown",
	}

	records := [][]string{}
	records = append(records, header)

	for _, a := range ts {
		elem := []string{
			a.DriverVersion,
			a.Target,
			a.KernelRelease,
			a.KernelVersion,
			a.DriverType,
			a.SourceIP,
		}

		statuses := map[string]uint{
			"200":     0,
			"2xx":     0,
			"3xx":     0,
			"404":     0,
			"4xx":     0,
			"5xx":     0,
			"unknown": 0,
		}

		for s, c := range a.Stats.Statuses {
			statusAggr := s

			switch s[0] {
			case '2':
				if s == "200" {
					statusAggr = s
					break
				}
				statusAggr = "2xx"
			case '3':
				statusAggr = "3xx"
			case '4':
				if s == "404" {
					statusAggr = s
					break
				}
				statusAggr = "4xx"
			case '5':
				statusAggr = "5xx"
			default:
				statusAggr = "unknown"
			}

			statuses[statusAggr] += c
		}

		elem = append(elem,
			fmt.Sprintf("%d", statuses["200"]),
			fmt.Sprintf("%d", statuses["2xx"]),
			fmt.Sprintf("%d", statuses["3xx"]),
			fmt.Sprintf("%d", statuses["404"]),
			fmt.Sprintf("%d", statuses["4xx"]),
			fmt.Sprintf("%d", statuses["5xx"]),
			fmt.Sprintf("%d", statuses["unknown"]),
		)

		records = append(records, elem)
	}

	for _, record := range records {
		if err := w.Write(record); err != nil {
			logrus.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		logrus.Fatal(err)
	}
}
