// ЗАДАЧА 7: GetFirstResult и GetResults
// Напишите функцию getFirstResult, которая принимает контекст и запускает конкурентый поиск, возвращая первый
// доступный результат из replicas. Возвращать ошибку контекста, если контекст завершился раньше, чем стал доступен
// какой-то результат из реплики.
// Напишите функцию getResults, которая запускает конкурентный поиск для каждого набора реплик из replicaKinds,
// использую getFirstResult, и возвращает результат для каждого набора реплик.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	msg string
	err error
}

type searh func() *result
type replicas []searh

func fakeSearch(kind string) searh {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	resChan := make(chan *result, len(replicas))

	for _, search := range replicas {
		go func() {
			resChan <- search()
		}()
	}

	var res *result
	select {
	case <-ctx.Done():
		res = &result{
			err: ctx.Err(),
		}
		return res
	case res = <-resChan:
		return res
	}
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	results := make([]*result, len(replicaKinds))
	var wg sync.WaitGroup

	wg.Add(len(replicaKinds))
	for i, k := range replicaKinds {
		go func() {
			defer wg.Done()
			results[i] = getFirstResult(ctx, k)
		}()
	}

	wg.Wait()
	return results
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	replicaKinds := []replicas{
		replicas{fakeSearch("web1"), fakeSearch("web2")},
		replicas{fakeSearch("image1"), fakeSearch("image2")},
		replicas{fakeSearch("video1"), fakeSearch("video2")},
	}

	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}
