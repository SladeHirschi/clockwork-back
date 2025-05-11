package api

func (c *ClockworkApi) addRoutes() {
	c.Router.HandleFunc("/clock-in", c.ClockIn)
}
