package covid19

func (api options) ByGlobalConfirmed() string {
	url := api.url + "/confirmed"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func (api options) ByGlobalDeaths() string {
	url := api.url + "/deaths"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func (api options) ByGlobalRecovered() string {
	url := api.url + "/recovered"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

// GlobalSummary is defining all infos about covid19
func (api options) ByGlobalSummary() string {
	resp, err := api.req.Get(api.url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}
