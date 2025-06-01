func main() {
	ctx := context.Background()
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	svc := new(frontendServer)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, propagation.Baggage{}))

	baseUrl = os.Getenv("BASE_URL")
	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	addr := os.Getenv("LISTEN_ADDR")

	// -----------------------------
	// DUMMY MODE STARTS HERE 🧪
	// -----------------------------

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🎉 Frontend is up — dummy mode!")
	})

	// optional: add /_healthz so health checks pass
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	var handler http.Handler = r
	handler = &logHandler{log: log, next: handler}
	handler = ensureSessionID(handler)
	handler = otelhttp.NewHandler(handler, "frontend")

	log.Infof("🌍 Starting dummy frontend server at http://%s:%s", addr, srvPort)
	log.Fatal(http.ListenAndServe(addr+":"+srvPort, handler))
}
