import (
        "flag"
        "log"
        "net/http"
)

var (
        dirname string
        port    string
)

func main() {
        flag.StringVar(&dirname, "d", ".", "directory to serve. default (.)")
        flag.StringVar(&port, "p", "8002", "port to bind to. default (8002)")
        log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir(dirname))))
}
