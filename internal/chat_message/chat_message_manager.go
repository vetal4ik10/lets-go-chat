package chat_message

import (
	"database/sql"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"time"
)

type ChatMessageManager interface {
	SaveMessage(cM ChatMessage) (ChatMessage, error)
	LoadAllMessages() []chatMessage
}

type chatMessageManager struct {
	db       *sql.DB
	userRepo reposetories.UserRepo
}

func NewChatMessageManager(db *sql.DB, userRepo reposetories.UserRepo) *chatMessageManager {
	return &chatMessageManager{db, userRepo}
}

func (cMM *chatMessageManager) SaveMessage(cM ChatMessage) (ChatMessage, error) {
	mTime := cM.GetTime()
	if mTime == nil {
		now := time.Now()
		mTime = &now
	}

	sqlStatement := `INSERT INTO "chat_message" ("uid", "text", "create") VALUES ($1, $2, $3)`
	_, err := cMM.db.Exec(sqlStatement, cM.GetUser().Uid, cM.GetText(), mTime)
	if err != nil {
		return nil, err
	}
	return &chatMessage{cM.GetUser(), cM.GetText(), mTime}, nil
}

func (cMM *chatMessageManager) LoadAllMessages() []chatMessage {
	messages := make([]chatMessage, 0)

	sqlStatement := `SELECT "uid", "text", "create" FROM "chat_message" ORDER BY "create" ASC`
	rows, err := cMM.db.Query(sqlStatement)
	if err != nil {
		return messages
	}

	for rows.Next() {
		var uid string
		var text []byte
		var create time.Time

		err = rows.Scan(&uid, &text, &create)
		if err != nil {
			return messages
		}

		u, _ := cMM.userRepo.GetByUid(uid)
		messages = append(messages, chatMessage{u, text, &create})
	}

	return messages
}
