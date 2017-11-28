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
	//return nil, errors.New("Prune not implemented")\
	fmt.Println(">>Prune: Total records : ", len(cm.records))
	//	return nil, nil
	for _, cr := range cm.records {
		fmt.Println("analyising container id : ", cr.ID())

		fmt.Println("Determining Size ", cr.ID())
		/*	size, err := cr.Size(ctx)
			if err != nil {
				//	cr.mu.Unlock()
				fmt.Println("Cannot determine size ", cr.ID())
				continue
				//	return nil, err
			}*/
		err := cr.mref().release(ctx)
		if err != nil {
			//	cr.mu.Unlock()
			fmt.Println("Cannot delete the cache ", cr.ID(), err)
			//return nil, err
			continue
		}
		fmt.Println("Success..  ")

	}
	fmt.Println("Returning")
	removeMap := make(map[string]int64)
	return removeMap, nil
}

func (cm *cacheManager) GC(ctx context.Context) error {
	return errors.New("GC not implemented")
}
