package cache

import (
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/penguin-statistics/backend-next/internal/utils/cache"
)

var (
	ItemFromId    cache.Cache
	ItemFromArkId cache.Cache

	StageFromId    cache.Cache
	StageFromArkId cache.Cache

	ZoneFromId    cache.Cache
	ZoneFromArkId cache.Cache

	DropPatternElementsFromId cache.Cache

	TimeRangeFromId cache.Cache

	once sync.Once
)

func Populate(client *redis.Client) {
	once.Do(func() {
		ItemFromId = cache.New(client, "item#id")
		ItemFromArkId = cache.New(client, "item#arkId")

		StageFromId = cache.New(client, "stage#id")
		StageFromArkId = cache.New(client, "stage#arkId")

		ZoneFromId = cache.New(client, "zone#id")
		ZoneFromArkId = cache.New(client, "zone#arkId")

		DropPatternElementsFromId = cache.New(client, "dropPatternElement#id")

		TimeRangeFromId = cache.New(client, "timeRange#id")
	})
}
