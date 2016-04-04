package main

import (
	"expvar"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"go4.org/types"

	"github.com/thecodearchive/gitarchive/camli"
	"github.com/thecodearchive/gitarchive/git"
	"github.com/thecodearchive/gitarchive/queue"
)

type Fetcher struct {
	q *queue.Queue
	u *camli.Uploader

	exp *expvar.Map

	closing uint32
}

func (f *Fetcher) Run() error {
	for atomic.LoadUint32(&f.closing) == 0 {
		name, parent, err := f.q.Pop()
		if err != nil {
			return err
		}

		if name == "" {
			interruptableSleep(30 * time.Second)
			continue
		}

		if err := f.Fetch(name, parent); err != nil {
			return err
		}
	}
	return nil
}

func (f *Fetcher) Fetch(name, parent string) error {
	url := "https://github.com/" + name + ".git"

	repo, err := f.u.GetRepo(url)
	if err != nil {
		return err
	}

	haves := make(map[string]struct{})
	if repo != nil {
		for _, have := range repo.Refs {
			haves[have] = struct{}{}
		}
	}

	// TODO: benchmark not adding the parent to the haves when we have cloned
	// the repo before, as it might trigger more optimizations on the server
	// side, even if it causes more bandwidth in fork syncs
	if parent != "" {
		mainURL := "https://github.com/" + parent + ".git"
		mainRepo, err := f.u.GetRepo(mainURL)
		if err != nil {
			return err
		}
		if mainRepo != nil {
			for _, have := range mainRepo.Refs {
				haves[have] = struct{}{}
			}
		}
	}

	logVerb, logFork := "Cloning", ""
	if repo != nil {
		logVerb = "Fetching"
	}
	if parent != "" {
		logFork = fmt.Sprintf(" (fork of %s)", parent)
	}
	log.Printf("[+] %s %s%s...", logVerb, name, logFork)

	res, err := git.Fetch(url, haves, f.u, os.Stderr)
	if err != nil {
		return err
	}

	return f.u.PutRepo(&camli.Repo{
		Name:      url,
		Retrieved: types.Time3339(time.Now()),
		Refs:      res.Refs,
	})
}

func (f *Fetcher) Stop() {
	atomic.StoreUint32(&f.closing, 1)
}

func interruptableSleep(d time.Duration) bool {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)
	t := time.NewTimer(d)
	select {
	case <-c:
		return false
	case <-t.C:
		return true
	}
}
