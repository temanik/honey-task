// ЗАДАЧА 12: WaitGroup с maxParallel
// Есть интерфейс waiter который должен:
// 1. параллельно запускать переданные в run функции с указанным контекстом.
// 2. количество параллельных запусков определяется параметром maxParallel при создании waiter через newGroupWait
// 3. возвращать ошибку из wait, если хотя бы одна функция из run вернула ее
// 4. возвращать комбинацию ошибок от вызовов run, если несколько задач завершились с ошибками (можно использовать errors.Join)
// Реализуйте поля и методы структуры waitGroup для интерфейса waiter.

package main

import (
	"context"
	"errors"
	"sync"
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

type waitGroup struct {
	mu     sync.Mutex
	wg     sync.WaitGroup
	bucket chan struct{}
	err    error
}

func (g *waitGroup) wait() error {
	g.wg.Wait()
	return g.err
}

func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	select {
	case <-ctx.Done():
	case g.bucket <- struct{}{}:
	}

	g.wg.Add(1)
	go func() {
		defer func() {
			<-g.bucket
			g.wg.Done()
		}()

		err := fn(ctx)
		if err != nil {
			g.mu.Lock()
			defer g.mu.Unlock()

			if g.err == nil {
				g.err = err
			} else {
				g.err = errors.Join(g.err, err)
			}
		}
	}()

	return
}

func newGroupWait(maxParallel int) waiter {
	return &waitGroup{
		bucket: make(chan struct{}, maxParallel),
	}
}

func main() {
	g := newGroupWait(2)

	ctx := context.Background()

	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")

	g.run(ctx, func(ctx context.Context) error {
		return nil
	})

	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})

	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})

	err := g.wait()
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		panic("wrong code")
	}
}
