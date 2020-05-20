package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Req is to store user delivery request
type Req struct {
	FromCode string `json:"fromcode"`
	ToCode   string `json:"tocode"`
	Provider string `json:"provider"`
}

// ResultType type struct
type ResultType struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Tariff      int    `json:"tariff"`
	Etd         string `json:"etd"`
}

// StatusType type struct
type StatusType struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// ProviderType type struct
type ProviderType struct {
	Results []ResultType `json:"results"`
	Status  StatusType   `json:"status"`
}

// ProviderSiCepat type struct
type ProviderSiCepat struct {
	SiCepat ProviderType `json:"sicepat"`
}

// Delivery exported function
func Delivery(w http.ResponseWriter, r *http.Request) {
	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	fCode := checkProviderCode(req.FromCode, req.Provider)
	tCode := checkProviderCode(req.ToCode, req.Provider)

	API := readAPIConfig(req.Provider)

	respond := requestAPI(fmt.Sprintf(API["url"], fCode, tCode), API["key"])

	if respond.SiCepat.Status.Code == 200 {
		for i := 0; i < len(respond.SiCepat.Results); i++ {
			provider := &req.Provider
			service := &respond.SiCepat.Results[i].Service
			cost := &respond.SiCepat.Results[i].Tariff
			serviceDescription := &respond.SiCepat.Results[i].Description
			etd := &respond.SiCepat.Results[i].Etd

			sql := fmt.Sprintf("INSERT INTO api03.cost (from_code, to_code, provider, service, cost, service_description, etd) VALUES ('%s','%s','%s','%s',%v,'%s','%s')", fCode, tCode, *provider, *service, *cost, *serviceDescription, *etd)
			_, err := db.Exec(sql)
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Fprintf(w, "Success, log to database")
	} else {
		fmt.Fprintf(w, "Failed, Input again")
	}
}

func requestAPI(url string, contentValue string) ProviderSiCepat {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("api-key", contentValue)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	var respond ProviderSiCepat
	json.NewDecoder(resp.Body).Decode(&respond)
	return respond
}

func readAPIConfig(provider string) map[string]string {
	apiConfigJSON, err := os.Open("./app/APIConfig.json")
	if err != nil {
		panic(err.Error())
	}

	defer apiConfigJSON.Close()

	var API map[string]map[string]string
	json.NewDecoder(apiConfigJSON).Decode(&API)

	var sortedAPI map[string]string
	sortedAPI = API[provider]

	return sortedAPI
}

func checkProviderCode(code, provider string) string {
	statement := fmt.Sprintf("SELECT provider_code FROM api03.location WHERE location_code = '%s' AND provider_name = '%s'", code, provider)

	query, err := db.Query(statement)
	if err != nil {
		panic(err.Error())
	}

	var result string
	for query.Next() {
		query.Scan(&result)
	}

	return result
}
