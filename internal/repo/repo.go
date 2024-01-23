package repo

import (
	"database/sql"
	"errors"
	"log"

	"github.com/leoff00/ta-pago-bot/internal/models"
)

type DiscordUserRepository struct{}

func (dur *DiscordUserRepository) GetUsers() []models.DiscordReturnType {
	db, err := sql.Open("sqlite3", "ta_pago.db")
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return nil
	}

	var arr []models.DiscordReturnType

	rows, err := db.Query(`SELECT username, count FROM DISCORD_USERS ORDER BY count DESC LIMIT 10`)

	if err != nil {
		log.Default().Println("Cannot get users from DB on Repo.", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var username string
		var count int

		if err := rows.Scan(&username, &count); err != nil {
			log.Default().Println("Cannot attach db returned values into var on Repo.", err.Error())
		}

		drt := models.DiscordReturnType{
			Username: username,
			Count:    count,
		}
		arr = append(arr, drt)
	}
	return arr
}

func (dur *DiscordUserRepository) Save(du models.DiscordUser) error {

	db, err := sql.Open("sqlite3", "ta_pago.db")
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return err
	}

	usr := dur.getUserById(du.Id)

	if usr.Id != "" {
		return errors.New("você já está inscrito na lista fera")
	}

	rows, err := db.Exec(`INSERT INTO DISCORD_USERS (id, username, count) VALUES (?, ?, ?)`, du.Id, du.Username, du.Count)

	if err != nil {
		log.Default().Println("Cannot insert data into DB on Repo.", err.Error())
		return err
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}

	log.Default().Println("Rows affected on Save User ->", affected)
	return err

}

func (dur *DiscordUserRepository) getUserById(discordId string) *models.DiscordUser {
	db, err := sql.Open("sqlite3", "ta_pago.db")
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return nil
	}

	var du models.DiscordUser
	row := db.QueryRow(`SELECT id, username, count FROM DISCORD_USERS WHERE id = ?`, discordId)

	row.Scan(&du.Id, &du.Username, &du.Count)

	defer db.Close()
	return &du
}

func (dur *DiscordUserRepository) IncrementCount(discordId string) error {

	db, err := sql.Open("sqlite3", "ta_pago.db")
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return err
	}

	usr := dur.getUserById(discordId)

	if usr.Id == "" {
		return errors.New("você precisa antes se inscrever na lista fera")
	}

	rows, err := db.Exec(`UPDATE DISCORD_USERS SET count = count + 1 WHERE id = ?`, discordId)

	if err != nil {
		log.Default().Println("Cannot update the count from DB on Repo.", err.Error())
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}

	defer db.Close()

	log.Default().Println("Rows affected on Update Count ->", affected)
	return err

}

func (dur *DiscordUserRepository) RestartCount() error {
	db, err := sql.Open("sqlite3", "ta_pago.db")
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return err
	}

	rows, err := db.Exec(`UPDATE DISCORD_USERS SET count = 0`)

	if err != nil {
		log.Default().Println("Cannot restart the count from DB on Repo.", err.Error())
		return err
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}

	defer db.Close()

	log.Default().Println("Rows affected on Update Count ->", affected)
	return err
}
