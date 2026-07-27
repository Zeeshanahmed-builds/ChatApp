package repo
import (
	"chat-app/db"
	"database/sql"
	"chat-app/models"
	"log"
	"golang.org/x/crypto/bcrypt"

)
type Users struct{
	db   *sql.DB
}

func NewUsers(db *sql.DB) Users {
	return Users{
		db: db,
	}
}

func (u *Users) SignUp(users *models.Users) error{
	
	log.Println("Signing up user:", users.Name, users.Email)


	hashPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}


	_, err = db.DB.Exec(
		"INSERT INTO users(name, email, password) VALUES ($1, $2, $3)",
		users.Name,
		users.Email,
		string(hashPassword),
	)

	log.Println("User signed up successfully:", users.Name)
	return err
	
}



func (u *Users)Login(login *models.Login) (*models.Users, error) {

	row := db.DB.QueryRow(`
		SELECT id, name, email, password
		FROM users
		WHERE email = $1 
	`, login.Email)

	var Resp models.Users

	err := row.Scan(
		&Resp.Users_ID,
		&Resp.Name,
		&Resp.Email,
		&Resp.Password,
	)
	if err != nil {
		return nil, err
	}

	return &Resp, err
}



