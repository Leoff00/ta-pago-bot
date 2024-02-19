package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/leoff00/ta-pago-bot/internal/domain"
	"github.com/leoff00/ta-pago-bot/internal/models"
	"github.com/leoff00/ta-pago-bot/pkg/discord"
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

var (
	DB_PATH     = env.Getenv("DB_PATH")
	DB_NAME     = env.Getenv("DB_NAME")
	DB_FULLPATH = DB_PATH + "/" + DB_NAME
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (dur *UserRepository) GetUsers() ([]models.DiscordRankType, error) {
	db, err := sql.Open("sqlite3", DB_FULLPATH)
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return nil, err
	}
	defer db.Close()

	var rankList []models.DiscordRankType
	rows, err := db.Query(`SELECT username, nickname, count FROM DISCORD_USERS ORDER BY count DESC LIMIT 10`)
	if err != nil {
		log.Default().Println("Cannot get users from DB on Repo.", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var nickname string
		var count int
		var username string

		if err := rows.Scan(&username, &nickname, &count); err != nil {
			log.Default().Println("Cannot attach db returned values into var on Repo.", err.Error())
		}
		if nickname == "" {
			nickname = username
		}
		rank := models.DiscordRankType{
			Nickname: nickname,
			Count:    count,
		}
		rankList = append(rankList, rank)
	}
	return rankList, nil
}

func (dur *UserRepository) Create(user *domain.User) error {
	db, err := sql.Open("sqlite3", DB_FULLPATH)
	defer db.Close()
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return err
	}

	usr := dur.GetUserById(user.Id)

	if usr.Id != "" {
		return errors.New("você já está inscrito na lista fera")
	}

	rows, err := db.Exec(`INSERT INTO DISCORD_USERS (id, username, updated_at, count, nickname) VALUES (?, ?, ?, ?, ?)`, user.Id, user.Username, 0, user.Count, user.Nickname)

	if err != nil {
		log.Default().Println("Cannot insert data into DB on Repo.", err.Error())
		return err
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}

	log.Default().Println("Rows affected on Create User ->", affected)
	return err
}

func (dur *UserRepository) GetUserById(discordId string) *domain.User {
	db, err := sql.Open("sqlite3", DB_FULLPATH)
	defer db.Close()
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return nil
	}

	var du domain.User
	row := db.QueryRow(`SELECT id, username, updated_at, count, nickname FROM DISCORD_USERS WHERE id = ?`, discordId)

	err = row.Scan(&du.Id, &du.Username, &du.Updated_at, &du.Count, &du.Nickname)
	if err != nil {
		log.Default().Println("Cannot get the user from DB on Repo.", err.Error())
		return &du
	}

	return &du
}

func (dur *UserRepository) Save(aggregate *models.UserAggregate) error {
	user := aggregate.User
	normalize(user, aggregate.DiscordUser)
	db, err := sql.Open("sqlite3", DB_FULLPATH)
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return err
	}
	defer db.Close()
	userToJson, _ := json.Marshal(user)
	log.Default().Println("User to save ->", string(userToJson))
	rows, err := db.Exec(`UPDATE DISCORD_USERS SET updated_at = ?,nickname = ?, count = ? WHERE id = ?`, user.Updated_at, user.Nickname, user.Count, user.Id)
	if err != nil {
		log.Default().Println("Cannot update the count from DB on Repo.", err.Error())
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
	}
	log.Default().Println("Rows affected on Update Count ->", affected)
	return err
}

func (dur *UserRepository) ResetCount() error {
	db, err := sql.Open("sqlite3", DB_FULLPATH)
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
		return err
	}
	defer db.Close()

	rows, err := db.Exec(`UPDATE DISCORD_USERS SET count = 0`)

	if err != nil {
		log.Default().Println("Cannot restart the count from DB on Repo.", err.Error())
		return err
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
	}

	log.Default().Println("Rows affected on Update Count ->", affected)
	return err
}

func (dur *UserRepository) ExistsById(id string) (bool, error) {
	db, err := sql.Open("sqlite3", DB_FULLPATH)
	if err != nil {
		log.Default().Println("Cannot open the DB on Repo ->", err.Error())
	}
	defer db.Close()
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM DISCORD_USERS WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func normalize(actual *domain.User, coming *discord.UserData) {
	if coming.Nickname != actual.Nickname {
		actual.Nickname = coming.Nickname
	}
	//if coming.Nickname != actual.Nickname {
	//	rows, err := db.Exec(`UPDATE DISCORD_USERS SET nickname = ? WHERE id = ?`, actual.Nickname, actual.Id)
	//	if err != nil {
	//		log.Default().Println("Cannot normalize the nickname in DB.", err.Error())
	//	}
	//	affected, _ := rows.RowsAffected()
	//	log.Default().Println("Rows affected on Update Nickname Normalization ->", affected)
	//}
	return
}
