package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
)

//var Conn *pgx.Conn = nil

type DB struct {
	Conn *pgx.Conn
	DatabaseURL string
	Cache map[string]string
}

var Database = DB{
	Conn: nil,
	Cache: make(map[string]string),
}


//var Cache = map[string]string{}

//var DatabaseURL string

func (db *DB) Connect(databaseUrl string) error {
	db.DatabaseURL = databaseUrl
	if db.Conn != nil && !db.Conn.IsClosed() {
		return fmt.Errorf("database already connected")
	}
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return err
	}
	//defer conn.Close(context.Background())
	db.Conn = conn

	fmt.Println("Connected to database.")
	return nil
}

func (db *DB) CheckConn() {
	if db.Conn == nil || db.Conn.IsClosed() {
		err := db.Connect(db.DatabaseURL)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (db *DB) GetChannelID(guildID string) (string, bool) {
	var channelID string

	if channelID, ok := db.Cache[guildID]; ok {
		return channelID, true
	}

	err := db.Conn.QueryRow(context.Background(),
		"SELECT channel_id FROM guild_config WHERE guild_id = $1", guildID).Scan(&channelID)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	if channelID == "" {
		return "", false
	} else {
		return channelID, true
	}
}

func (db *DB) AddGuildConfig(guildID, channelID string) error {
	db.Cache[guildID] = channelID
	db.CheckConn()
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO guild_config (guild_id, channel_id) VALUES ($1, $2)", guildID, channelID)
	return err
}

func (db *DB) EditChannelID(guildID, channelID string) error {
	db.Cache[guildID] = channelID
	db.CheckConn()
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE guild_config SET channel_id = $2 WHERE guild_id = $1", guildID, channelID)
	return err
}

func (db *DB) SetChannelID(guildID, channelID string) error {
	var exists bool
	db.CheckConn()
	err := db.Conn.QueryRow(context.Background(),
		"SELECT exists(SELECT 1 FROM guild_config WHERE guild_id = $1)", guildID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return db.EditChannelID(guildID, channelID)
	} else {
		return db.AddGuildConfig(guildID, channelID)
	}
}

func (db *DB) RemoveChannelID(guildID string) error {
	return db.SetChannelID(guildID, "")
}

func (db *DB) DelChannelID(guildID string) error {
	delete(db.Cache, guildID)
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM guild_config WHERE guild_id = $1", guildID)
	return err
}