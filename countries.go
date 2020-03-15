package covid19

func (api options) ByAllCountriesInfo() string {
	url := api.url + "/countries"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func (api options) ByCountrySummary(country string) string {
	url := api.url + "/countries/" + country
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	// util := resty.Client{}
	// data, _ := util.JSONMarshal(resp)
	return resp.String()
}

func (api options) ByCountryConfirmed(country string) string {
	url := api.url + "/countries/" + country + "/confirmed"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func (api options) ByCountryRecovered(country string) string {
	url := api.url + "/countries/" + country + "/recovered"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func (api options) ByCountryDeaths(country string) string {
	url := api.url + "/countries/" + country + "/deaths"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}
