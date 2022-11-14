package repo

import (
	"context"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"

	"exusiai.dev/backend-next/internal/model"
)

type RecognitionDefect struct {
	DB *bun.DB
}

func NewRecognitionDefect(db *bun.DB) *RecognitionDefect {
	return &RecognitionDefect{DB: db}
}

func (s *RecognitionDefect) CreateDefectReportDraft(ctx context.Context, defectReport *model.RecognitionDefect) error {
	if defectReport.DefectID == "" {
		defectReport.DefectID = strings.ToLower(ulid.Make().String())
	}

	_, err := s.DB.NewInsert().
		Model(defectReport).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecognitionDefect) FinalizeDefectReport(ctx context.Context, defectId, imageUri string) error {
	_, err := s.DB.NewUpdate().
		Model((*model.RecognitionDefect)(nil)).
		Set("image_uri = ?", imageUri).
		Set("updated_at = ?", time.Now()).
		Where("defect_id = ?", defectId).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecognitionDefect) GetDefectReports(ctx context.Context, limit, offset int) ([]*model.RecognitionDefect, error) {
	var defectReports []*model.RecognitionDefect

	err := s.DB.NewSelect().
		Model(&defectReports).
		Limit(limit).
		Offset(offset).
		Order("defect_id DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return defectReports, nil
}