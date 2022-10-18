package v2

import (
	"github.com/goccy/go-json"

	"github.com/uptrace/bun"

	"github.com/penguin-statistics/backend-next/internal/model"
)

type DropInfo struct {
	bun.BaseModel `bun:"drop_infos"`

	DropID     int             `bun:",pk,autoincrement" json:"-"`
	Server     string          `json:"-"`
	StageID    int             `json:"-"`
	ItemID     int             `json:"-"`
	ArkStageID string          `bun:"-" json:"-"`
	ArkItemID  string          `bun:"-" json:"itemId,omitempty"`
	DropType   string          `json:"dropType"`
	RangeID    int             `json:"-"`
	Bounds     json.RawMessage `json:"bounds" swaggertype:"object"`
	Extras     json.RawMessage `json:"-" swaggertype:"object"`

	Item      *Item            `bun:"rel:belongs-to,join:item_id=item_id" json:"-"`
	Stage     *Stage           `bun:"rel:belongs-to,join:stage_id=stage_id" json:"-"`
	TimeRange *model.TimeRange `bun:"rel:belongs-to,join:range_id=range_id" json:"-"`
}
