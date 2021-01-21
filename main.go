package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/jasontconnell/fedproxy/conf"
)

type handler struct {
	cfg        conf.Config
	intercepts map[string]string
}

type response struct {
	body    []byte
	headers http.Header
	status  int
}

func main() {
	c := flag.String("c", "config.json", "config file")
	flag.Parse()

	if *c == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg := conf.LoadConfig(*c)
	if len(cfg.Intercepts) == 0 || len(cfg.ProxyHost) == 0 {
		log.Println("need intercepts and proxy host to start up")
		os.Exit(1)
	}

	if cfg.LocalPort == 0 {
		cfg.LocalPort = 7676

	}

	if !filepath.IsAbs(cfg.LocalStartPath) {
		wd, _ := os.Getwd()
		cfg.LocalStartPath = filepath.Join(wd, cfg.LocalStartPath)

		d, err := os.Stat(cfg.LocalStartPath)
		if err != nil || !d.IsDir() {
			log.Fatal("local start path not found or not a directory", cfg.LocalStartPath)
		}
	}

	fn := getHandler(cfg)
	hnd := fn

	log.Println("starting fedproxy")
	log.Println("listening on port", cfg.LocalPort)
	log.Println("intercepting requests for resources to serve from", cfg.LocalStartPath)
	for _, v := range cfg.Intercepts {
		log.Println("intercepting file type", v.Extension, v.MimeType)
	}
	log.Println("forwarding all other requests to", cfg.ProxyHost)

	url := ":" + strconv.Itoa(cfg.LocalPort)
	log.Fatal(http.ListenAndServe(url, hnd))
}

func getHandler(cfg conf.Config) handler {
	h := handler{cfg: cfg}
	imap := make(map[string]string)
	for _, icp := range cfg.Intercepts {
		imap["."+icp.Extension] = icp.MimeType
	}
	h.intercepts = imap
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ext := path.Ext(req.URL.Path)
	var resp response
	var err error

	if mimetype, ok := h.intercepts[ext]; ok {
		resp, err = getLocalContent(h.cfg.LocalStartPath, req.URL.Path, mimetype)
	} else {
		resp, err = getRemoteContent(h.cfg.ProxyScheme, h.cfg.ProxyHost, req, h.cfg.RequestHeaders)
	}

	if err != nil {
		http.Error(w, err.Error(), resp.status)
		return
	}

	for k, v := range resp.headers {
		for _, s := range v {
			w.Header().Add(k, s)
		}
	}
	w.WriteHeader(resp.status)
	w.Write(resp.body)
}

func getLocalContent(local, path, mime string) (response, error) {
	r := response{status: 404}
	p := filepath.Join(local, path)
	f, err := os.Open(p)
	if err != nil {
		return r, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return r, err
	}
	r.status = 200
	r.body = b
	r.headers = make(http.Header)
	r.headers["Content-Type"] = []string{mime}

	return r, nil
}

func getRemoteContent(scheme, host string, req *http.Request, reqHeaders map[string]string) (response, error) {
	c := http.Client{}
	r := response{status: 500}

	var u *url.URL = req.URL
	u.Host = host
	u.Scheme = scheme

	nreq, err := http.NewRequest(req.Method, u.String(), req.Body)
	if err != nil {
		return r, err
	}

	for k, v := range reqHeaders {
		nreq.Header.Add(k, v)
	}

	nresp, err := c.Do(nreq)
	if err != nil {
		return r, err
	}
	defer nresp.Body.Close()

	b, err := ioutil.ReadAll(nresp.Body)
	if err != nil {
		return r, err
	}

	r.body = b
	r.headers = nresp.Header
	r.status = nresp.StatusCode

	log.Println(u, "content length:", len(r.body), "status code:", r.status)

	return r, nil
}
