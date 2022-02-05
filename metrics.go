package ttlcache

// Metrics contains common cache metrics calculated over the course
// of the cache's lifetime.
type Metrics struct {
	// Inserts specifies how many items were inserted.
	Inserts uint64

	// Hits specifies how many items were successfully retrieved
	// from the cache.
	// Retrievals made with a loader function are not tracked.
	Hits uint64

	// Misses specifies how many items were not found in the cache.
	// Retrievals made with a loader function are tracked as well.
	Misses uint64

	// Evictions specifies how many items were removed from the
	// cache.
	Evictions uint64
}
