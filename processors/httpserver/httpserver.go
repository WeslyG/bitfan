//go:generate bitfanDoc
// Listen and read a http request to build events with it.
//
// Processor respond with a HTTP code as :
//
// * `202` when request has been accepted, in body : the total number of event created
// * `500` when an error occurs, in body : an error description
//
// Use codecs to process body content as json / csv / lines / json lines / ....
//
// URL is available as http://webhookhost/pluginLabel/URI
//
// * webhookhost is defined by bitfan at startup
// * pluginLabel is defined in pipeline configuration, it's the named processor if you put one, or `input_httpserver` by default
// * URI is defined in plugin configuration (see below)
package httpserverprocessor

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/vjeantet/bitfan/codecs"
	"github.com/vjeantet/bitfan/processors"
)

func New() processors.Processor {
	return &processor{opt: &options{}}
}

type options struct {
	// Add a field to an event
	Add_field map[string]interface{}

	// Add any number of arbitrary tags to your event.
	// This can help with processing later.
	Tags []string

	// Add a type field to all events handled by this input
	Type string

	// The codec used for input data. Input codecs are a convenient method for decoding
	// your data before it enters the input, without needing a separate filter in your bitfan pipeline
	// @Default "plain"
	// @Type codec
	Codec codecs.Codec

	// URI path
	// @Default "events"
	Uri string
}

// Reads events from standard input
type processor struct {
	processors.Base

	opt  *options
	q    chan bool
	host string
}

func (p *processor) Configure(ctx processors.ProcessorContext, conf map[string]interface{}) error {
	defaults := options{
		Codec: codecs.New("plain", nil, ctx.Log(), ctx.ConfigWorkingLocation()),
		Uri:   "events",
	}
	p.opt = &defaults
	err := p.ConfigureAndValidate(ctx, conf, p.opt)

	if p.host, err = os.Hostname(); err != nil {
		p.Logger.Warnf("can not get hostname : %s", err.Error())
	}

	return err
}
func (p *processor) Start(e processors.IPacket) error {
	p.q = make(chan bool)
	p.WebHook.Add(p.opt.Uri, p.HttpHandler)
	return nil
}

// Handle Request received by bitfan for this agent (url hook should be registered during p.Start)
func (p *processor) HttpHandler(w http.ResponseWriter, r *http.Request) {
	p.Logger.Debug("reading request")

	// Create a reader
	var dec codecs.Decoder
	var err error

	if dec, err = p.opt.Codec.NewDecoder(r.Body); err != nil {
		p.Logger.Errorln("decoder error : ", err.Error())
		return
	}
	headersBytes, _ := httputil.DumpRequest(r, false)
	headers := string(headersBytes)

	req := map[string]interface{}{
		"remoteAddr":  r.RemoteAddr,
		"rawHeaders":  headers,
		"method":      r.Method,
		"requestURI":  r.RequestURI,
		"proto":       r.Proto,
		"host":        r.Host,
		"headers":     r.Header,
		"requestPath": r.URL.Path,
		"querystring": r.URL.Query(),
	}
	if r.Method == "POST" {
		r.ParseForm()
		req["formvalues"] = r.PostForm
	}

	var nbEvents int
	p.Logger.Debug("request = ", req)
	p.Logger.Debug("start reading body content")

	for dec.More() {
		var record interface{}
		if err = dec.Decode(&record); err != nil {
			if err == io.EOF {
				p.Logger.Warnln("error while http read docoding : ", err)
			} else {
				p.Logger.Errorln("error while http read docoding : ", err)
				break
			}
		}

		var e processors.IPacket
		switch v := record.(type) {
		case nil:
			e = p.NewPacket("", map[string]interface{}{
				"request": req,
			})
		case string:
			e = p.NewPacket(v, map[string]interface{}{
				"request": req,
			})
		case map[string]interface{}:
			e = p.NewPacket("", v)
			e.Fields().SetValueForPath(req, "request")
		}

		processors.ProcessCommonFields(e.Fields(), p.opt.Add_field, p.opt.Tags, p.opt.Type)
		p.Send(e)
		nbEvents++
		select {
		case <-p.q:
			return
		default:
		}
	}

	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("500 - " + err.Error()))
	} else {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Sprintf("%d", nbEvents)))
	}

}

func (p *processor) Stop(e processors.IPacket) error {
	close(p.q)
	return nil
}
