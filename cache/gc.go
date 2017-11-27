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
		//	cr.mu.Lock()
		// ignore duplicates that share data
		/*if cr.equalImmutable != nil && len(cr.equalImmutable.refs) > 0 || cr.equalMutable != nil && len(cr.refs) == 0 {
			cr.mu.Unlock()
			fmt.Println("Lets skip, id : ", cr.ID())
			continue
		}*/
		fmt.Println("Determining Size ", cr.ID())
		size, err := cr.Size(ctx)
		if err != nil {
			//	cr.mu.Unlock()
			fmt.Println("Cannot determine size ", cr.ID())
			return nil, err
		}
		if size >= 0 {
			fmt.Println("Removing.. : ", cr.ID())
			if cr.Parent() != nil {
				fmt.Println("Tree Below")
			}
			for temp := cr.Parent(); temp == nil; temp = temp.Parent() {
				fmt.Println("-> ", temp.ID())
				//	temp = temp.Parent()

			}
			cr.Parent().ID()
			/*err := cr.remove(ctx, true)
			if err != nil {
				//	cr.mu.Unlock()
				fmt.Println("Cannot delete the cache ", cr.ID(), err)
				return nil, err
			}*/
			fmt.Println("Success..  ")
		}
		//	cr.mu.Unlock()
		//		fmt.Println("Size determined : ", size)
	}
	fmt.Println("Returning")
	return nil, nil
}

func (cm *cacheManager) GC(ctx context.Context) error {
	return errors.New("GC not implemented")
}
