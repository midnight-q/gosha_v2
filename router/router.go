package router

import (
	"encoding/json"
	"gosha_v2/settings"
	"gosha_v2/webapp"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Router - маршрутизатор
func Router() http.Handler {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(settings.HomePageRoute, homePage).Methods("GET")

	//[ Model ]
	router.HandleFunc(settings.ModelRoute, webapp.ModelFind).Methods("GET")
	router.HandleFunc(settings.ModelRoute, webapp.ModelCreate).Methods("POST")
	router.HandleFunc(settings.ModelRoute+"/list", webapp.ModelMultiCreate).Methods("POST")
	router.HandleFunc(settings.ModelRoute+"/list", webapp.ModelMultiUpdate).Methods("PUT")
	router.HandleFunc(settings.ModelRoute+"/{id}", webapp.ModelUpdate).Methods("PUT")

	//[ Field ]
	router.HandleFunc(settings.FieldRoute, webapp.FieldCreate).Methods("POST")
	router.HandleFunc(settings.FieldRoute+"/list", webapp.FieldMultiCreate).Methods("POST")
	router.HandleFunc(settings.FieldRoute+"/list", webapp.FieldMultiUpdate).Methods("PUT")
	router.HandleFunc(settings.FieldRoute+"/list", webapp.FieldMultiDelete).Methods("DELETE")
	router.HandleFunc(settings.FieldRoute+"/{id}", webapp.FieldUpdate).Methods("PUT")
	router.HandleFunc(settings.FieldRoute+"/{id}", webapp.FieldDelete).Methods("DELETE")

	//[ Field ]
	router.HandleFunc(settings.FieldRoute, webapp.FieldCreate).Methods("POST")
	router.HandleFunc(settings.FieldRoute+"/list", webapp.FieldMultiCreate).Methods("POST")
	router.HandleFunc(settings.FieldRoute+"/list", webapp.FieldMultiUpdate).Methods("PUT")
	router.HandleFunc(settings.FieldRoute+"/{id}", webapp.FieldUpdate).Methods("PUT")

	//[ FieldType ]
	router.HandleFunc(settings.FieldTypeRoute, webapp.FieldTypeFind).Methods("GET")

	//[ Application ]
	router.HandleFunc(settings.ApplicationRoute, webapp.ApplicationFind).Methods("GET")
	router.HandleFunc(settings.ApplicationRoute, webapp.ApplicationCreate).Methods("POST")
	router.HandleFunc(settings.ApplicationRoute+"/list", webapp.ApplicationMultiCreate).Methods("POST")
	router.HandleFunc(settings.ApplicationRoute+"/list", webapp.ApplicationMultiUpdate).Methods("PUT")
	router.HandleFunc(settings.ApplicationRoute+"/{id}", webapp.ApplicationUpdate).Methods("PUT")

	//[ Sdk ]
	router.HandleFunc(settings.SdkRoute, webapp.SdkCreate).Methods("POST")

	return applyCors(router)
}

func applyCors(router *mux.Router) http.Handler {
	return cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"token", "content-type"},
	}).Handler(router)
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	type Response struct {
		Version string
		Date    string
	}

	_ = json.NewEncoder(w).Encode(Response{
		Version: "0.0.1",
		Date:    "2021.08.03 13:41:33",
	})
}
