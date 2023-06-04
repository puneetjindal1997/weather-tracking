package router

import (
	"net/http"
	"weather/controller"
)

var WeatherRoutes = Routes{
	Route{"Get temp of perticular city", http.MethodGet, "/getWeather", controller.GetCityWeather},
}
