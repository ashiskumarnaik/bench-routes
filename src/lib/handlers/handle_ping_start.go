package handlers

import (
	"github.com/zairza-cetb/bench-routes/src/lib/logger"
	"github.com/zairza-cetb/bench-routes/src/lib/parser"
)

//// HandleFloodPingStart starts the flood ping service
//func HandleFloodPingStart(config *parser.YAMLBenchRoutesType, floodPingServiceState string) {
//	floodPingConfig := config.Config.Routes
//	floodPingInterval := GetInterval(config.Config.Interval, "ping")
//	if floodPingInterval == (TestInterval{}) {
//		logger.Terminal("interval not found in configuration file for ping", "f")
//		return
//	}
//
//	urlStack := make(map[string]string)
//	for _, route := range floodPingConfig {
//		url := route.URL
//		urlHash := utils.GetHash(url)
//		// maintain urls uniquely
//		_, ok := urlStack[urlHash]
//		if !ok {
//			urlStack[urlHash] = *filters.HTTPPingFilter(&url)
//		}
//	}
//
//	doFloodPing(config, urlStack, floodPingInterval)
//}
//
//func doFloodPing(config *parser.YAMLBenchRoutesType, urlStack map[string]string, interval TestInterval) {
//	i := 0
//	for {
//		i++
//		config.Refresh()
//
//		switch config.Config.UtilsConf.ServicesSignal.FloodPing {
//		case "active":
//			var wg sync.WaitGroup
//			wg.Add(len(urlStack))
//			for _, u := range urlStack {
//				go ping.HandleFloodPing(utils.FPingc, u, 500, u, &wg, false, config.Config.Password)
//			}
//
//			wg.Wait()
//		case "passive":
//			//terminate the goroutine
//			logger.Terminal("terminating flood ping goroutine", "p")
//			return
//		default:
//			logger.Terminal("invalid service-state value for flood-ping", "f")
//			return
//		}
//
//		intrv := time.Duration(interval.duration)
//		switch interval.ofType {
//		case "hr":
//			time.Sleep(intrv * time.Hour)
//		case "min":
//			time.Sleep(intrv * time.Minute)
//		case "sec":
//			time.Sleep(intrv * time.Second)
//		default:
//			logger.Terminal("invalid interval-type for flood-ping", "f")
//			return
//		}
//	}
//}
//
//TestInterval stores the value of the duration and the type of test
type TestInterval struct {
	ofType   string
	duration int64
}

//GetInterval extracts the interval type and value from the config file for the particular test
func GetInterval(intervals []parser.Interval, testName string) TestInterval {
	for _, intrv := range intervals {
		if testName == intrv.Test {
			return TestInterval{
				ofType:   intrv.Type,
				duration: *intrv.Duration,
			}
		}
	}

	// if the execution reaches this then it is an error as the interval was not found in the configuration file
	logger.Terminal("interval not found in the configuration file", "pa")
	return TestInterval{}
}
