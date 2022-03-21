package opensky

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const openskyAddressAll string = "https://opensky-network.org/api/states/all"

//Reader struct defines the opensky connector parameters
//boundixBoxParam defines teh bounding box for the request
//apiKey is the api key for the request to access paid services
//requestInterval is the interval in seconds for a new request
//client httpClient used to request the resource
type Reader struct {
	boundingBoxParam string
	apiKey           string //not sure if they have that.. seems to be USERNAME:PASSWORD as part of Request URL???
	requestInterval  uint
	client           *http.Client
}

//NewReader constructs a new default reader with given request interval
func NewReader(interval uint) *Reader {
	return &Reader{requestInterval: interval}
}

//SetBoundingBox sets a bounding box for the given reader. This will be used to limit the geographical boundary
func (o *Reader) SetBoundingBox(lamin, lomin, lamax, lomax float64) {

	o.boundingBoxParam = fmt.Sprintf("lamin=%.5f&lomin=%.5f&lamax=%.5f&lomax=%.5f", lamin, lomin, lamax, lomax)
}

//Connect to the OpenSky Network and request new data at configured interval. Resonses will be made available through the provided channel
func (o *Reader) Connect(c chan *StateVectorResponse) {
	ticker := time.NewTicker(time.Duration(o.requestInterval) * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		s, err := o.SendRequest()
		if err != nil {
			fmt.Println(err)
		} else {
			c <- s
		}
	}

}

//getRequestUrl used to construct the final request Url given all possible parameters, e.g. Bounding Box
func (o *Reader) getRequestUrl() (url string) {
	if o.boundingBoxParam != "" {
		url = fmt.Sprintf("%s?%s", openskyAddressAll, o.boundingBoxParam)
	}

	return
}

//SendRequest to send a request using the given reader
//returns  the resulting StateVectorResponse and error if present
func (o *Reader) SendRequest() (*StateVectorResponse, error) {
	if o.client == nil {

		o.client = &http.Client{Timeout: time.Duration(2 * time.Second)}
	}

	url := o.getRequestUrl()

	r, err := o.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var result StateVectorResponse
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = result.UnmarshalResponse(b) //json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
