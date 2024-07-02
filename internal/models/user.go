package models

// User représente la structure d'un utilisateur en base de données et en JSON
type User struct {
	ID       int    `json:"id" db:"id"`             // ID de l'utilisateur
	Name     string `json:"name" db:"name"`         // Nom de l'utilisateur
	Email    string `json:"email" db:"email"`       // Email de l'utilisateur
	Password string `json:"password" db:"password"` // Mot de passe de l'utilisateur
}

// CreateUser crée un nouvel utilisateur dans la base de données
func CreateUser(user User) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetUserByEmail récupère un utilisateur à partir de son email depuis la base de données
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID récupère un utilisateur à partir de son ID depuis la base de données
func GetUserByID(id int) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
