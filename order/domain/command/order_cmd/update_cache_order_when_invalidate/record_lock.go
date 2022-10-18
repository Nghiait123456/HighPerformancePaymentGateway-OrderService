package update_cache_order_when_invalidate

type (
	RecordLock = DataSavedCache
)

func CreateRecordLock() RecordLock {
	return RecordLock{
		Data:       DataQuery{},
		IsHaveLock: true,
	}
}
