package gorelicwrap

import (
	"log"
	"net/http"
)

type AppMetaData struct {
	Applications []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"applications"`
}

type AppInstMetaData struct {
	ApplicationsInstances []struct {
		ID                 int    `json:"id"`
		ApplicationName    string `json:"application_name"`
		Host               string `json:"host"`
		Language           string `json:"language"`
		ApplicationSummary struct {
			InstanceCount int     `json:"instance_count"`
			Throughput    float64 `json:"throughput"`
			ResponseTime  float64 `json:"response_time"`
			ApdexScore    float64 `json:"apdex_score"`
		} `json:"application_summary"`
		Links struct {
			Applicaiton     int `json:"applicaiton"`
			ApplicationHost int `json:"application_host"`
		} `json:"links"`
	} `json:"application_instances"`
}

type AppMetaDataParams struct {
	FilterName string `url:"filter[name],omitempty"`
}

type AppInstMetaDataParams struct {
	FilterHostName string `url:"filter[hostname],omitempty"`
	FilterIds      string `url:"filter[ids],omitempty"`
}

func (client *NRClient) GetAppID(appName string) int {
	app := AppMetaData{}
	params := &AppMetaDataParams{FilterName: appName}
	resp, err := client.sling.New().Get("applications.json").QueryStruct(params).ReceiveSuccess(&app)
	checkResponse(resp, err)
	return app.Applications[0].ID
}

func (client *NRClient) GetAppInstances(appID string) *AppInstMetaData {
	appInst := AppInstMetaData{}
	resp, err := client.sling.New().Get("applications/" + appID + "/instances.json").ReceiveSuccess(&appInst)
	checkResponse(resp, err)
	return &appInst
}

func (client *NRClient) GetAppMetricData(appID string, params *MetricParms) *MetricData {
	metricData := MetricData{}
	pathURL := "applications/" + appID + "/metrics/data.json"
	resp, err := client.sling.New().Get(pathURL).QueryStruct(params).ReceiveSuccess(&metricData)
	checkResponse(resp, err)
	return &metricData
}

func (client *NRClient) GetAppInstancesMetricData(appID, appInstID string, params *MetricParms) *MetricData {
	metricData := MetricData{}
	pathURL := "applications/" + appID + "/instances/" + appInstID + "/metrics/data.json"
	resp, err := client.sling.New().Get(pathURL).QueryStruct(params).ReceiveSuccess(&metricData)
	checkResponse(resp, err)
	return &metricData
}

func checkResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatalln("Query of NewRelic endpoint resulted in an error:", err)
	}

	if resp.StatusCode != 200 {
		log.Println(resp.Request)
		log.Fatalln("Query of NewRelic endpoint was not successful", resp.StatusCode)
	}
}
