package auth

import (
	"errors"
	"fmt"
	"hexxcore/config"
	"hexxcore/types"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateSubjectsForUser(tx *gorm.DB, userID uint) error {
	subjectNames := []string{"ADA", "IT", "SE", "IC", "LANG", "ENG", "OE", "ADA Lab", "IT Lab"}

	for _, name := range subjectNames {
		var count int64
		tx.Model(&types.Subject{}).Where("user_id = ? AND name = ?", userID, name).Count(&count)
		if count == 0 {
			if err := tx.Create(&types.Subject{Name: name, UserID: userID, StartDate: time.Now()}).Error; err != nil {
				return err // Rollback will happen in the caller function
			}
		}
	}

	return nil
}

func (s *Store) CreateUser(user *types.User) error {
	if err := config.Validator.Struct(user); err != nil {
		return err
	}

	tx := s.db.Begin() // Start transaction

	// Create user & ensure user.ID is set
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value") {
			return fmt.Errorf("user with this email or register number already exists")
		}
		return err
	}

	// Now user.ID is set, pass the transaction to CreateSubjectsForUser
	if err := s.CreateSubjectsForUser(tx, user.ID); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (s *Store) GetUserByIdentifier(identifier string) (*types.User, error) {
	user := new(types.User)
	if err := s.db.Where("email = ? OR register = ?", identifier, identifier).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id uint) (*types.User, error) {
	user := new(types.User)

	if err := s.db.Where("id = ?", id).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (s *Store) UpdateUser(user *types.User) error {
	tx := s.db.Begin()

	currentUser := new(types.User)
	if err := tx.Where("id = ?", user.ID).First(currentUser).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("user not found %v", err)
	}
	if err := tx.Model(currentUser).Updates(user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating user %v", err)
	}
	return tx.Commit().Error
}

func (s *Store) DeleteUser(id uint) error {
	user := new(types.User)
	if err := s.db.Unscoped().Where("id = ?", id).First(user).Error; err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	tx := s.db.Begin()

	if user.DeletedAt.Valid {
		// User already soft deleted — delete permanently
		if err := tx.Unscoped().Where("user_id = ?", user.ID).Delete(&types.Subject{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Unscoped().Delete(user).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// Soft delete the user
		if err := tx.Delete(user).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
