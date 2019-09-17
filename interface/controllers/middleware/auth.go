package middleware

import (
	"context"

	"CleanArchitecture_SampleApp/interface/database"
	"CleanArchitecture_SampleApp/interface/dcontext"
	"CleanArchitecture_SampleApp/interface/network"
)

type middleware struct {
	userRepository database.UserRepository
}

type MiddleWare interface {
	UserAuthorize(ar network.ApiResponser) network.ApiResponser
}

func NewMiddleWare(db database.ConnectedSql) MiddleWare {
	return &middleware{
		userRepository: database.NewUserRepository(db),
	}
}

//Headerにあるx-tokenからユーザーを特定して情報を保存する、通信の前処理
func (mw *middleware) UserAuthorize(ar network.ApiResponser) network.ApiResponser {

	ctx := ar.GetRequestContext()
	if ctx == nil {
		ctx = context.Background()
	}

	// リクエストヘッダからx-token(認証トークン)を取得
	token := ar.GetRequest().GetHeaderValue("x-token")
	if len(token) == 0 {
		ar.BadRequest("x-token is empty")
	}

	// データベースから認証トークンに紐づくユーザの情報を取得
	user, err := mw.userRepository.FindByAuthToken(token)
	if err != nil {
		ar.InternalServerError("User is not found: Not matching token found")
	}

	// userIdをContextへ保存して以降の処理に利用する
	ctx = dcontext.SetUserID(ctx, user.UserID)
	// requestにctxをセット
	ar.SetRequestContext(ctx)
	// 前処理（認証）を終えて、実際の処理＝HandleFuncを実行する
	return ar
}
