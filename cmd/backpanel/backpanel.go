package main

import (
	"expvar"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/thecodearchive/gitarchive/index"
	"github.com/thecodearchive/go-trello"
)

type Backpanel struct {
	c     *trello.Client
	i     *index.Index
	exp   *expvar.Map
	pause time.Duration

	blacklistID string

	closing uint32
}

type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

var (
	// Hardcoded because I'm awful.
	// curl "https://trello.com/1/boards/04pbw4Gv/labels?value=&key=$key&token=$token"
	labelGH        = "5778e0ee84e677fd365e136e"
	labelFork      = "5778e8c51040134fc3126996"
	labelTooBig    = "5778e0ee84e677fd365e136d"
	labelCrash     = "5778e0ee84e677fd365e136c"
	label10Stars   = "5778e0ee84e677fd365e1369"
	label100Stars  = "5778e0ee84e677fd365e136a"
	label1000Stars = "5778e0ee84e677fd365e136b"
)

func (b *Backpanel) Run() error {
	for atomic.LoadUint32(&b.closing) == 0 {
		board, err := b.c.Board(b.blacklistID)
		if err != nil {
			return errors.Wrapf(err, "getting board %s", b.blacklistID)
		}

		var whitelist, blacklist *trello.List
		lists, err := board.Lists()
		if err != nil {
			return errors.Wrapf(err, "getting board %s lists", b.blacklistID)
		}
		for _, list := range lists {
			if list.Name == "Whitelist" {
				l := list
				whitelist = &l
			}
			if list.Name == "Blacklist" {
				l := list
				blacklist = &l
			}
		}
		if whitelist == nil || blacklist == nil {
			return errors.New("Lists not found.")
		}

		tr := make(map[string]trello.Card)
		cards, err := board.Cards()
		if err != nil {
			// This is quite large and sometimes the fetch fails with io.ErrUnexpectedEOF.
			// However, the trello package opaquely wraps that so we can't special case it.
			// Just sleep for a bit longer instead of crashing ¯\_(ツ)_/¯.
			// TODO Backoff
			log.Printf("[-] Error getting board %s cards: %v", b.blacklistID, err)
			if !interruptableSleep(2 * b.pause) {
				return nil
			}
			continue
		}
		for _, card := range cards {
			tr[card.Name] = card
		}

		ble, err := b.i.ListBlacklist()
		if err != nil {
			return err
		}

		for _, e := range ble {
			n := strings.TrimPrefix(e.Name, "github.com/")
			if card, ok := tr[n]; !ok {
				// New line in the database, add card.
				card := trello.Card{Name: n}
				card.Desc = "https://github.com/" + n
				card.IdLabels = append(card.IdLabels, labelGH)
				if e.Reason == "Too big." {
					card.IdLabels = append(card.IdLabels, labelTooBig)
				} else {
					card.IdLabels = append(card.IdLabels, labelCrash)
					card.Desc += "\n\n" + e.Reason
				}
				list := blacklist
				if e.State == index.Whitelisted {
					list = whitelist
				}
				_, err := list.AddCard(card)
				if err != nil {
					return errors.Wrapf(err, "Adding card %#v", card)
				}
				b.exp.Add("newcard", 1)

			} else {
				// Card present, check if it matches.
				// So far the only supported action is changing lists.
				if card.IdList == whitelist.Id && e.State != index.Whitelisted {
					log.Println("Whitelisting", e.Name)
					err := b.i.SetBlacklistState(e.Name, index.Whitelisted)
					if err != nil {
						return err
					}
					_, err = card.AddComment("Applied whitelist!")
					if err != nil {
						return errors.Wrapf(err, "Adding comment to card %s", card.Name)
					}
					b.exp.Add("moved", 1)
				}
				if card.IdList == blacklist.Id && e.State != index.Blacklisted {
					log.Println("Blacklisting", e.Name)
					err := b.i.SetBlacklistState(e.Name, index.Blacklisted)
					if err != nil {
						return err
					}
					_, err = card.AddComment("Applied blacklist!")
					if err != nil {
						return errors.Wrapf(err, "Adding comment to card %s", card.Name)
					}
					b.exp.Add("moved", 1)
				}
				delete(tr, n)
			}
		}

		// Only new entries are left now in tr.
		for _, card := range tr {
			if card.IdList == blacklist.Id {
				log.Println("Blacklisting (new)", card.Name)
			} else if card.IdList == whitelist.Id {
				log.Println("Whitelisting (new)", card.Name)
			} else {
				continue
			}
			err := b.i.AddBlacklist("github.com/"+card.Name, card.Desc)
			if err != nil {
				return err
			}
			msg := "blacklist"
			if card.IdList == whitelist.Id {
				err := b.i.SetBlacklistState("github.com/"+card.Name, index.Whitelisted)
				if err != nil {
					return err
				}
				msg = "whitelist"
			}
			_, err = card.AddComment("Added to " + msg + "!")
			if err != nil {
				return errors.Wrapf(err, "Adding comment to card %s", card.Name)
			}
			b.exp.Add("newline", 1)
		}

		// TODO: also keep the labels and descriptions up to date.

		if !interruptableSleep(b.pause) {
			return nil
		}
	}
	return nil
}

func (b *Backpanel) Stop() {
	atomic.StoreUint32(&b.closing, 1)
}

func interruptableSleep(d time.Duration) bool {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(c)
	t := time.NewTimer(d)
	select {
	case <-c:
		return false
	case <-t.C:
		return true
	}
}
