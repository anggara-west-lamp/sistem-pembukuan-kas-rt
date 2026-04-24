package models

import "time"

type Role struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `gorm:"size:50;uniqueIndex" json:"name"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type User struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `gorm:"size:120" json:"name"`
    Email     string `gorm:"size:180;uniqueIndex" json:"email"`
    Password  string `json:"-"`
    RoleID    uint   `json:"role_id"`
    Role      Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Kas mewakili entitas kas (tipe & kategori) dari diagram.
type Kas struct {
    ID          uint    `gorm:"primaryKey" json:"id"`
    Type        string  `gorm:"size:20" json:"type"`      // masuk|keluar
    Category    string  `gorm:"size:100" json:"category"`
    Amount      float64 `json:"amount"`                    // opsional, master nilai default
    Description string  `gorm:"size:255" json:"description"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Transaction struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    UserID    uint    `json:"user_id"`
    User      User    `json:"user"`
    KasID     uint    `json:"kas_id"`
    Kas       Kas     `json:"kas"`
    Amount    float64 `json:"amount"`
    Status    string  `gorm:"size:30" json:"status"` // pending|approved|rejected|settled
    ProofURL  string  `gorm:"size:255" json:"proof_url"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Report dapat berupa snapshot bulanan; dapat juga dihitung on-the-fly.
type Report struct {
    ID           uint    `gorm:"primaryKey" json:"id"`
    Month        string  `gorm:"size:7;index" json:"month"` // YYYY-MM
    TotalIncome  float64 `json:"total_income"`
    TotalExpense float64 `json:"total_expense"`
    Balance      float64 `json:"balance"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

