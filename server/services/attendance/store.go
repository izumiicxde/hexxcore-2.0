package attendance

import (
	"fmt"
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
func (s *Store) GetAttendanceSummary(userId uint) (*types.AttendanceSummary, error) {
	var totalClasses int64
	var attended int64
	var skipped int64
	var subjectStats []types.SubjectStats

	// Get total classes attended by the user
	err := s.db.Model(&types.Attendance{}).
		Where("user_id = ?", userId).
		Count(&totalClasses).Error
	if err != nil {
		return nil, fmt.Errorf("error getting total classes: %w", err)
	}

	// Count attended classes
	err = s.db.Model(&types.Attendance{}).
		Where("user_id = ? AND status = ?", userId, "Present").
		Count(&attended).Error
	if err != nil {
		return nil, fmt.Errorf("error getting attended classes: %w", err)
	}

	// Count skipped classes
	err = s.db.Model(&types.Attendance{}).
		Where("user_id = ? AND status = ?", userId, "Absent").
		Count(&skipped).Error
	if err != nil {
		return nil, fmt.Errorf("error getting skipped classes: %w", err)
	}

	// Calculate how many more can be skipped to maintain 75% attendance
	neededFor75 := int(float64(totalClasses) * 0.75)
	canSkipMore := int(attended) - neededFor75
	if canSkipMore < 0 {
		canSkipMore = 0
	}

	// Fetch per-subject stats
	rows, err := s.db.Raw(`
		SELECT s.name, 
		       COUNT(a.id) AS total,
		       SUM(CASE WHEN a.status = 'Present' THEN 1 ELSE 0 END) AS attended,
		       SUM(CASE WHEN a.status = 'Absent' THEN 1 ELSE 0 END) AS skipped
		FROM attendances a
		JOIN subjects s ON a.subject_id = s.id
		WHERE a.user_id = ?
		GROUP BY s.name
	`, userId).Rows()
	if err != nil {
		return nil, fmt.Errorf("error fetching per-subject stats: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var subject types.SubjectStats
		if err := rows.Scan(&subject.SubjectName, &subject.Total, &subject.Attended, &subject.Skipped); err != nil {
			return nil, err
		}
		subjectStats = append(subjectStats, subject)
	}

	return &types.AttendanceSummary{
		TotalClasses: int(totalClasses),
		Attended:     int(attended),
		Skipped:      int(skipped),
		AllowedSkips: canSkipMore,
		Subjects:     subjectStats,
	}, nil
}

func (s *Store) MarkAttendance(req *types.AttendanceRequest, userId uint) error {
	parsedDate, err := time.Parse("2006-01-02", req.Date) // Expecting format "YYYY-MM-DD"
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	tx := s.db.Begin()

	for _, subject := range req.Subjects {
		var existingSubject types.Subject

		// Check if the subject exists for the user
		err := tx.Where("name = ? AND user_id = ?", subject.Name, userId).First(&existingSubject).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("subject '%s' not found for user", subject.Name)
		}

		// Convert boolean status to string
		status := "Not Marked"
		if subject.Status {
			status = "Present"
		} else {
			status = "Absent"
		}

		// Create attendance record
		attendance := types.Attendance{
			UserID:    userId,
			SubjectID: existingSubject.ID,
			Date:      parsedDate,
			Status:    status,
		}

		// Insert attendance record
		if err := tx.Create(&attendance).Error; err != nil {
			tx.Rollback()

			if strings.Contains(err.Error(), "duplicate key value") || strings.Contains(err.Error(), "unique constraint") {
				return fmt.Errorf("attendance already marked for this subject on this date")
			}

			return err
		}

		// Update subject stats
		updateFields := map[string]interface{}{
			"total_taken": gorm.Expr("total_taken + 1"),
		}

		if status == "Present" {
			updateFields["attended"] = gorm.Expr("attended + 1")
		} else if status == "Absent" {
			updateFields["missed"] = gorm.Expr("missed + 1")
		}

		// Apply updates to the subject record
		if err := tx.Model(&existingSubject).Updates(updateFields).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update subject stats for '%s'", subject.Name)
		}
	}

	// Commit transaction if everything is successful
	return tx.Commit().Error
}

func (s *Store) GetAllSubjects() ([]string, error) {
	var subjects []string
	err := s.db.Model(&types.Schedule{}).Distinct().Pluck("subject_name", &subjects).Error
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *Store) GetTodaysClasses() ([]string, error) {
	today := time.Now().Weekday().String()

	var classes []types.Schedule
	err := s.db.Select("subject_name").Where("day_of_week = ?", today).Find(&classes).Error
	if err != nil {
		return nil, err
	}

	// Extract subject names from the query result
	subjectNames := make([]string, len(classes))
	for i, class := range classes {
		subjectNames[i] = class.SubjectName
	}

	return subjectNames, nil
}
