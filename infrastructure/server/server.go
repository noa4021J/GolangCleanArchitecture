package server

import (
	"log"
	"net/http"
)

type server struct{}

type Server interface {
	Start(addr string)
	Get(endPoint string, apiFunc http.HandlerFunc)
	Post(endPoint string, apiFunc http.HandlerFunc)
}

func New() Server {
	return &server{}
}

func (s *server) Start(addr string) {
	// サーバの起動
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// GETリクエストを処理する
func (s *server) Get(endPoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endPoint, httpMethod(apiFunc, http.MethodGet))
}

// POSTリクエストを処理する
func (s *server) Post(endPoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endPoint, httpMethod(apiFunc, http.MethodPost))
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
