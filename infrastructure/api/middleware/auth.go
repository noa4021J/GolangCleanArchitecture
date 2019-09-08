package middleware

import (
	"context"
	"log"
	"net/http"

	"CleanArchitecture_SampleApp/infrastructure/api/dcontext"
	"CleanArchitecture_SampleApp/infrastructure/datastore"
	"CleanArchitecture_SampleApp/infrastructure/server/response"
	"CleanArchitecture_SampleApp/interface/database"
)

type middleware struct {
	db *datastore.ConnectedSql
}

type MiddleWare interface {
	UserAuthorize(nextFunc http.HandlerFunc) http.HandlerFunc
}

func NewMiddleWare(db *datastore.ConnectedSql) MiddleWare {
	return &middleware{db: db}
}

//Headerにあるx-tokenからユーザーを特定して情報を保存する、通信の前処理
func (mw *middleware) UserAuthorize(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if len(token) == 0 {
			log.Println("x-token is empty")
			return
		}

		// データベースから認証トークンに紐づくユーザの情報を取得
		row := mw.db.QueryRow("SELECT * FROM user WHERE auth_token=?", token)
		if row == nil {
			response.BadRequest(writer, "User is not found: Not matching token found")
		}

		user, err := database.ConvertToUser(row)
		if err != nil {
			return
		}

		// userIdをContextへ保存して以降の処理に利用する
		ctx = dcontext.SetUserID(ctx, user.UserID)

		// 前処理（認証）を終えて、実際の処理＝HandleFuncを実行する
		nextFunc(writer, request.WithContext(ctx))
	}
}
