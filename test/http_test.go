package test

import (
	"testing"

	"github.com/Courtcircuits/optique-modules/http"
	httpclient "net/http"
)

func InitHTTP() http.Http {
	health := http.NewHealthController()
	app, err := http.NewHttp(http.Config{
		ListenAddr: ":8080",
	})
	if err != nil {
		app.Stop()
		panic(err)
	}
	app.WithHandler(health)
	return app
}

func IgniteHTTP(t *testing.T, app http.Http) {
	go func() {
		if err := app.Ignite(); err != nil {
			t.Fatal(err)
		}
	}()
}

func TestHTTPIsAccessible(t *testing.T) {
	app := InitHTTP()
	defer app.Stop()
	IgniteHTTP(t, app)
	resp, err := httpclient.Get("http://localhost:8080/health")
	if err != nil {
		app.Stop()
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		app.Stop()
		t.Fatal("Health check failed")
	}
}
