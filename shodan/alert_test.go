package shodan

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_DeleteAlert(t *testing.T) {
	mux, tearDownTestServe, client := setUpTestServe()
	defer tearDownTestServe()

	id := "ZZ4TDUUORVE1DIIP"
	path := fmt.Sprintf(alertDeletePath, id)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		fmt.Fprint(w, `{}`) //nolint:errcheck
	})

	result, err := client.DeleteAlert(context.TODO(), id)

	assert.Nil(t, err)
	assert.True(t, result)
}

func TestClient_GetAlert(t *testing.T) {
	mux, tearDownTestServe, client := setUpTestServe()
	defer tearDownTestServe()

	id := "ZZ4TDUUORVE1DIIP"
	path := fmt.Sprintf(alertInfoPath, id)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write(getStub(t, "alert/alert")) //nolint:errcheck
	})

	alert, err := client.GetAlert(context.TODO(), id)
	alertExpected := &Alert{
		ID:         "ZZ4TDUUORVE1DIIP",
		Name:       "Test alert",
		Created:    "2017-09-24T18:30:43.592000",
		Expires:    0,
		Expired:    false,
		Expiration: "",
		Filters: &AlertFilters{
			IP: []string{"198.20.22.0/24"},
		},
		Size: 256,
	}

	assert.Nil(t, err)
	assert.Equal(t, alertExpected, alert)
}

func TestClient_GetAlerts(t *testing.T) {
	mux, tearDownTestServe, client := setUpTestServe()
	defer tearDownTestServe()

	mux.HandleFunc(alertsInfoListPath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write(getStub(t, "alert/alerts")) //nolint:errcheck
	})

	alerts, err := client.GetAlerts(context.TODO())
	alertsExpected := []*Alert{
		{
			ID:         "ZZ4TDUUORVE1DIIP",
			Expired:    true,
			Name:       "Test alert",
			Created:    "2017-09-24T18:30:43.592000",
			Expires:    0,
			Expiration: "",
			Filters: &AlertFilters{
				IP: []string{"198.20.22.0/24"},
			},
			Size: 256,
		},
		{
			ID:         "IU0CJDXNNEXBOPK3",
			Name:       "Test alert 2",
			Expired:    false,
			Created:    "2017-09-24T20:08:51.815000",
			Expires:    100,
			Expiration: "2017-09-24T20:10:31.815000",
			Filters: &AlertFilters{
				IP: []string{"198.20.88.0/24"},
			},
			Size: 256,
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, alertsExpected, alerts)
}

func TestClient_CreateAlert(t *testing.T) {
	mux, tearDownTestServe, client := setUpTestServe()
	defer tearDownTestServe()

	mux.HandleFunc(alertCreatePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.Write(getStub(t, "alert/create_alert")) //nolint:errcheck
	})

	alert, err := client.CreateAlert(context.TODO(), "Test alert API", []string{"198.20.88.0/24"}, 0)
	alertExpected := &Alert{
		ID:         "JZT8NVWEZWCY79OO",
		Name:       "Test alert API",
		Created:    "2017-09-24T23:08:43.434646",
		Expires:    0,
		Expired:    false,
		Expiration: "",
		Filters: &AlertFilters{
			IP: []string{"198.20.88.0/24"},
		},
		Size: 256,
	}

	assert.Nil(t, err)
	assert.Equal(t, alertExpected, alert)
}
