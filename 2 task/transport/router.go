package transport

func (s *Server) InitRoutes() {
	v1 := s.router.Group("api/v1")
	v1.GET("/currencies", s.handler.GetCurrencies)
	v1.GET("/currencies/:currency_id", s.handler.GetCurrencyById)

}
