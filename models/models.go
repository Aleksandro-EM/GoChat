package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	CreatedChats []Chat       `gorm:"foreignKey:CreatedBy"`
	ChatMembers  []ChatMember `gorm:"foreignKey:UserID"`
	Messages     []Message    `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

type Chat struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	Name      string // For group chats
	IsGroup   bool   `gorm:"default:false"`
	CreatedBy string `gorm:"type:char(36);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Creator  User         `gorm:"foreignKey:CreatedBy"`
	Members  []ChatMember `gorm:"foreignKey:ChatID"`
	Messages []Message    `gorm:"foreignKey:ChatID"`
	Webhooks []Webhook    `gorm:"foreignKey:ChatID"`
}

func (c *Chat) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

type ChatMember struct {
	ChatID   string    `gorm:"type:char(36);primaryKey"`
	UserID   string    `gorm:"type:char(36);primaryKey"`
	JoinedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// Relationships
	Chat Chat `gorm:"foreignKey:ChatID"`
	User User `gorm:"foreignKey:UserID"`
}

type Message struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	ChatID    string `gorm:"type:char(36);not null;index"`
	UserID    string `gorm:"type:char(36);not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Chat Chat `gorm:"foreignKey:ChatID"`
	User User `gorm:"foreignKey:UserID"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}

type Webhook struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	ChatID    string `gorm:"type:char(36);not null;index"`
	Name      string `gorm:"not null"`
	URL       string `gorm:"type:text;not null"`
	Secret    string // For HMAC signing
	Events    string `gorm:"type:text"` // JSON array as string: ["message.sent", "user.joined"]
	IsActive  bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Chat        Chat         `gorm:"foreignKey:ChatID"`
	WebhookLogs []WebhookLog `gorm:"foreignKey:WebhookID"`
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = uuid.New().String()
	}
	return nil
}

type WebhookLog struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	WebhookID string    `gorm:"type:char(36);not null;index"`
	EventType string    `gorm:"size:100"`
	Payload   string    `gorm:"type:json"` // Store as JSON string
	Status    int       // HTTP status code
	Response  string    `gorm:"type:text"`
	SentAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// Relationships
	Webhook Webhook `gorm:"foreignKey:WebhookID"`
}

func (wl *WebhookLog) BeforeCreate(tx *gorm.DB) error {
	if wl.ID == "" {
		wl.ID = uuid.New().String()
	}
	return nil
}
