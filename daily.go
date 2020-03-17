package covid19

func (api options) ByDailyGlobal() string {
	url := api.url + "/daily"
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func (api options) ByDailyDate(dt string) string {
	url := api.url + "/daily/" + dt
	resp, err := api.req.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.String()
}
