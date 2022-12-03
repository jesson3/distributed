package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":9877"
const ServivesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registratins []Registeration
	mutex        *sync.Mutex
}

func (r *registry) add(reg Registeration) error {
	r.mutex.Lock()
	r.registratins = append(r.registratins, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{
	registratins: make([]Registeration, 0),
	mutex:        new(sync.Mutex),
}

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registeration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %v", r.ServiceName,
			r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 
	}
}
