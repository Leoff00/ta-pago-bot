package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	"github.com/leoff00/ta-pago-bot/internal/domain"
	"github.com/leoff00/ta-pago-bot/internal/models"
	"github.com/leoff00/ta-pago-bot/pkg/discord"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetUsersRank gets the top 10 users from the database
func (ur *UserRepository) GetUsersRank() ([]models.DiscordRankType, error) {
	var rankList []models.DiscordRankType
	rows, err := ur.db.Query(`SELECT username, nickname, count
										  FROM DISCORD_USERS
										  ORDER BY count DESC LIMIT 10`)
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

// Insert create a new user into the database
func (ur *UserRepository) Insert(user *domain.User) error {
	rows, err := ur.db.Exec(`
								INSERT INTO DISCORD_USERS (id, username, updated_at, count, nickname)
								VALUES (?, ?, ?, ?, ?)`, user.Id, user.Username, 0, user.Count, user.Nickname)
	if err != nil {
		log.Default().Println("Cannot insert data into DB on Repo.", err.Error())
		return err
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}
	log.Default().Println("Rows affected on Insert User ->", affected)
	return err
}

// GetUserById gets specific hydrated user entity from the database
func (ur *UserRepository) GetUserById(discordId string) (*domain.User, error) {
	var du domain.User
	row := ur.db.QueryRow(`SELECT id, username, updated_at, count, nickname
											FROM DISCORD_USERS WHERE id = ?`, discordId)
	err := row.Scan(&du.Id, &du.Username, &du.Updated_at, &du.Count, &du.Nickname)
	if errors.Is(err, sql.ErrNoRows) { // expected error
		return &du, nil
	}
	if err != nil {
		log.Default().Println("Cannot get the user from DB on Repo.", err.Error())
		return nil, err
	}
	return &du, nil
}

// Save updates/persist existent user on the database
func (ur *UserRepository) Save(aggregate *models.UserAggregate) error {
	user := aggregate.User
	normalize(user, aggregate.DiscordUser)
	userToJson, _ := json.Marshal(user)
	log.Default().Println("User to save ->", string(userToJson))
	rows, err := ur.db.Exec(`UPDATE DISCORD_USERS 
									 SET updated_at = ?,nickname = ?, count = ? WHERE id = ?`,
		user.Updated_at, user.Nickname, user.Count, user.Id)
	if err != nil {
		log.Default().Println("Cannot update the count from DB on Repo.", err.Error())
		return err
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}
	log.Default().Println("Rows affected on Update Count ->", affected)
	return err
}

func (ur *UserRepository) ResetCount() error {
	rows, err := ur.db.Exec(`UPDATE DISCORD_USERS SET count = 0`)
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

func (ur *UserRepository) EditCount(userId string, countValue int) error {
	rows, err := ur.db.Exec(`UPDATE DISCORD_USERS SET count = ? WHERE id = ?`, countValue, userId)
	if err != nil {
		log.Default().Println("Cannot restart the count from DB on Repo.", err.Error())
		return err
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		log.Default().Println("Cannot get the affected row line numbers on Repo.", err.Error())
		return err
	}
	log.Default().Println("Rows affected on Update Count ->", affected)
	return nil
}

func (ur *UserRepository) ExistsById(id string) (bool, error) {
	var count int
	err := ur.db.QueryRow("SELECT COUNT(*) FROM DISCORD_USERS WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func normalize(actual *domain.User, coming *discord.UserData) {
	if coming.Nickname != actual.Nickname {
		actual.Nickname = coming.Nickname
	}
}
