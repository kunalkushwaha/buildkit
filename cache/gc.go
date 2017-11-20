package cache

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// GCPolicy defines policy for garbage collection
type GCPolicy struct {
	MaxSize         uint64
	MaxKeepDuration time.Duration
}

// // CachePolicy defines policy for keeping a resource in cache
// type CachePolicy struct {
// 	Priority int
// 	LastUsed time.Time
// }
//
// func defaultCachePolicy() CachePolicy {
// 	return CachePolicy{Priority: 10, LastUsed: time.Now()}
// }

func (cm *cacheManager) Prune(ctx context.Context) (map[string]int64, error) {
	//return nil, errors.New("Prune not implemented")
	for _, cr := range cm.records {
		cr.mu.Lock()
		// ignore duplicates that share data
		if cr.equalImmutable != nil && len(cr.equalImmutable.refs) > 0 || cr.equalMutable != nil && len(cr.refs) == 0 {
			cr.mu.Unlock()
			continue
		}
		size, err := cr.Size(ctx)
		if err != nil {
			fmt.Println("Cannot determine size ", cr.ID())
			cr.mu.Unlock()
			return nil, err
		}
		if size >= 0 {
			err := cr.remove(ctx, true)
			if err != nil {
				fmt.Println("Cannot delete the cache ", cr.ID())
				cr.mu.Unlock()
				return nil, err
			}
		}
		cr.mu.Unlock()
	}
	return nil, nil
}

func (cm *cacheManager) GC(ctx context.Context) error {
	return errors.New("GC not implemented")
}
