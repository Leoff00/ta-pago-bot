package repo

import (
	"database/sql"
	"errors"
	"log"

	"github.com/leoff00/ta-pago-bot/internal/models"
)

// func GetUsers() []models.DiscordUser {
// 	db, err := sql.Open("sqlite3", "ta_pago.db")

// 	if err != nil {
// 		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
// 	}

// 	var arr []models.DiscordUser

// 	rows, err := Db.Query(`SELECT id, username, count FROM DISCORD_USERS`)

// 	if err != nil {
// 		log.Default().Println("Cannot get users from DB on Repo.", err.Error())
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var id string
// 		var username string
// 		var count int

// 		if err := rows.Scan(&id, &username, &count); err != nil {
// 			log.Default().Println("Cannot attach db returned values into var on Repo.", err.Error())
// 		}

// 		du := models.DiscordUser{
// 			Id:       id,
// 			Username: username,
// 			Count:    count,
// 		}
// 		arr = append(arr, du)
// 	}
// 	defer db.Close()

// 	return arr
// }

func Save(du models.DiscordUser) error {
	db, err := sql.Open("sqlite3", "ta_pago.db")

	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
	}

	usr := getUserById(du.Id)

	if usr != nil {
		return errors.New("você já está inscrito na lista fera")
	}

	rows, err := db.Exec(`INSERT INTO DISCORD_USERS (id, username, count) VALUES (?, ?, ?)`, du.Id, du.Username, du.Count)

	if err != nil {
		log.Default().Println("Cannot insert data into DB on Repo.", err.Error())
		return nil
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return nil
	}

	log.Default().Println("Rows affected on Save User ->", affected)

	return nil

}

func getUserById(discordId string) *models.DiscordUser {
	db, err := sql.Open("sqlite3", "ta_pago.db")
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return nil
	}

	var du models.DiscordUser
	row := db.QueryRow(`SELECT id, username, count FROM DISCORD_USERS WHERE id = ?`, discordId)

	if err = row.Scan(&du.Id, &du.Username, &du.Count); err != nil {
		log.Default().Println("Error fetching user from DB on Repo.", err.Error())
		return nil
	}

	return &du
}

// * Need to understand why updatecount is with
// * call side effect in /ta-pago cmd handler.
func UpdateCount(discordId string) error {
	db, err := sql.Open("sqlite3", "ta_pago.db")

	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
	}

	usr := getUserById(discordId)

	if err != nil || usr == nil {
		return errors.New("você já está inscrito na lista fera")
	}

	rows, err := db.Exec(`UPDATE DISCORD_USERS SET count = count + 1 WHERE id = ?`, discordId)

	if err != nil {
		log.Default().Println("Cannot update the count from DB on Repo.", err.Error())
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return nil
	}

	log.Default().Println("Rows affected on Update Count ->", affected)

	defer db.Close()
	return nil

}
