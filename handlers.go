package contactqr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

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
	PngBase64 string `json:"png_base64"`
}

// writeVCardResponse writes a response to vcard actions
func writeVCardResponse(w http.ResponseWriter, text, errors, pngBase64 string) {
	resp := VCardResponse{
		Success:   true,
		Errors:    errors,
		VCardText: text,
		PngBase64: pngBase64,
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.VCardText = ""
	}

	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

// CreateVCardHandler creates a new VCard
func CreateVCardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in VCardRequest
	body, _ := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(body, &in)

	// validate request
	if jsonErr != nil {
		log.Println("Error creating new vCard:", jsonErr.Error())
		writeVCardResponse(w, "", "Invalid request", "")
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
		writeVCardResponse(w, "", vcErr.Error(), "")
		return
	}

	log.Printf("Creating vCard for %s %s\n", vc.First(), vc.Last())
	w.WriteHeader(http.StatusOK)
	writeVCardResponse(w, vc.String(), "", vc.QRCode(200, 200))
}

// StaticHTMLHandler handles static html file requests
func StaticHTMLHandler(filePath, indexFile string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// service index.html
		if r.RequestURI == "/" || r.URL.Path[1:] == "/" {
			http.ServeFile(w, r, filepath.Join(filePath, indexFile))
			return
		}

		http.ServeFile(w, r, filepath.Join(filePath, r.URL.Path[1:]))
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
