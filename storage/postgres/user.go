package postgres

import (
	"database/sql"

	pb "root/genprotos"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) LoginUser(user *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	query := `
			SELECT id, user_name, soldier_id from admins 
			where user_name =$1 and password=$2 and delated_at=0
		`
	row := p.db.QueryRow(query, user.UserName, user.Password)

	var users pb.LoginUserResponse

	err := row.Scan(&users.Id,&users.UserName,
		&users.SolderId)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
