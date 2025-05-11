package api

func (c *ClockworkApi) addRoutes() {
	c.Router.HandleFunc("/login", c.Login).Methods("POST")
	c.Router.HandleFunc("/signup", c.Signup).Methods("POST")

	c.Router.HandleFunc("/clock-in", c.ClockIn).Methods("POST")
}
