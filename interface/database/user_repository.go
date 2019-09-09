package database

import (
	"CleanArchitecture_SampleApp/domain"
	"database/sql"
	"log"
)

type userRepository struct {
	db ConnectedSql
}

type UserRepository interface {
	Insert(user domain.User) error
	SelectByAuthToken(authToken string) (*domain.User, error)
	SelectByPrimaryKey(userID string) (*domain.User, error)
	UpdateByPrimaryKey(userID string, name string) error
}

func NewUserRepository(db ConnectedSql) UserRepository {
	return &userRepository{db}
}

// Insert データベースをレコードを登録する
func (userRepository *userRepository) Insert(user domain.User) error {
	_, err := userRepository.db.Exec("INSERT INTO user(user_id, auth_token, name) VALUES (?, ? ,?)", user.UserID, user.AuthToken, user.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SelectByAuthToken auth_tokenを条件にレコードを取得する
func (userRepository *userRepository) SelectByAuthToken(authToken string) (*domain.User, error) {
	row := userRepository.db.QueryRow("SELECT * FROM user WHERE auth_token=?", authToken)
	return ConvertToUser(row)
}

// SelectByPrimaryKey 主キーを条件にレコードを取得する
func (userRepository *userRepository) SelectByPrimaryKey(userID string) (*domain.User, error) {
	row := userRepository.db.QueryRow("SELECT * FROM user WHERE user_id=?", userID)
	return ConvertToUser(row)
}

// UpdateByPrimaryKey 主キーを条件にレコードを更新する
func (userRepository *userRepository) UpdateByPrimaryKey(userID string, name string) error {
	_, err := userRepository.db.Exec("UPDATE user SET name=? WHERE user_id=?", name, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// convertToUser rowデータをUserデータへ変換する
func ConvertToUser(row Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.UserID, &user.AuthToken, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
