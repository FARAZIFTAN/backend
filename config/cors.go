package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"https://indrariksa.github.io",
	"http://localhost:5173/",
	"http://localhost:5174/",
	"http://localhost:5175/",
	"http://localhost:5176/",
	"http://localhost:5177/",
	"http://localhost:5178/",
	"http://localhost:5179/",
	"http://127.0.0.1:8088/",
	"https://fe-jifpolte.netlify.app/",
}

func GetALLowedOrigins() []string {
	return allowedOrigins
}
