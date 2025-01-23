package user

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func readBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.WithFields(log.Fields{
			"error":   err,
			"request": r.RequestURI,
		}).Error("Read body error")

		return nil, err
	}
	return body, nil
}

func serializeJSON(w http.ResponseWriter, target any) ([]byte, error) {
	marshal, err := json.Marshal(target)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"error": err,
			"data":  target,
		}).Error("Serialize")
		return nil, err
	}
	return marshal, nil
}

func deserializeJSON(w http.ResponseWriter, data []byte, target any) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.WithFields(log.Fields{
			"error": err,
			"data":  target,
		}).Error("Deserialize")
		return err
	}
	return nil
}
