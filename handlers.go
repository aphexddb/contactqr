package contactqr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aphexddb/contactqr/vcard"
)

// VCardRequest represents a request for a new vCard
type VCardRequest struct {
	First         string `json:"first"`
	Last          string `json:"last"`
	CompanyName   string `json:"company_name"`
	Title         string `json:"title"`
	Email         string `json:"email"`
	CellPhone     string `json:"cell_phone"`
	Street        string `json:"street"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
	FacebookURL   string `json:"facebook_url"`
	TwitterHandle string `json:"twitter_handle"`
	URL           string `json:"url"`
	Note          string `json:"note"`
}

// VCardResponse represents the response for creating a new vCard
type VCardResponse struct {
	Success   bool   `json:"success"`
	Errors    string `json:"errors"`
	VCardText string `json:"vcard_text"`
}

// writeVCardResponse writes a response to vcard actions
func writeVCardResponse(w http.ResponseWriter, text, errors string) {
	resp := VCardResponse{
		Success:   true,
		Errors:    errors,
		VCardText: text,
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.VCardText = ""
	}

	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

// NewVCardHandler creates a new VCard
func NewVCardHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New vCard request")

	w.Header().Set("Content-Type", "application/json")

	var in VCardRequest
	body, _ := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(body, &in)

	// validate request
	if jsonErr != nil {
		writeVCardResponse(w, "", "Invalid request")
		return
	}

	vc, vcErr := vcard.New(
		vcard.First(in.First),
		vcard.Last(in.Last),
		vcard.Org(in.CompanyName),
		vcard.Title(in.Title),
		vcard.Email(in.Email),
		vcard.CellPhone(in.CellPhone),
		vcard.HomeAddress(vcard.Address{
			Street:     in.Street,
			City:       in.City,
			State:      in.State,
			PostalCode: in.PostalCode,
		}),
		vcard.FacebookProfileURL(in.FacebookURL),
		vcard.TwitterHandle(in.TwitterHandle),
		vcard.URL(in.URL),
		vcard.Note(in.Note),
	)
	if vcErr != nil {
		writeVCardResponse(w, "", vcErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeVCardResponse(w, vc.String(), "")
}

// IndexHandler handles index file requests
func IndexHandler(indexPath string) func(w http.ResponseWriter, r *http.Request) {
	log.Println("HTML index file is being served from:", indexPath)
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, indexPath)
	}

	return http.HandlerFunc(fn)
}

// HealthCheckHandler provides a simple health check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"healthy": true}`))
}

// NotFoundHandler handles 404 requests
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "404 Not Found"}`))
}
