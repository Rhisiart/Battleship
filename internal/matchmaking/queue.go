package matchmaking

import (
	"fmt"
	"sync"
)

type Queue struct {
	players     []*Player
	lock        sync.Mutex
	maxPlayers  int
	matchmaking chan *Match
}

func NewQueue(maxPlayers int) *Queue {
	return &Queue{
		players:     []*Player{},
		matchmaking: make(chan *Match),
		maxPlayers:  maxPlayers,
	}
}

func (q *Queue) Enqueue(player *Player) {
	fmt.Println("Enqueue the player")
	q.lock.Lock()
	defer q.lock.Unlock()

	q.players = append(q.players, player)
	q.CheckQueue()
}

func (q *Queue) CheckQueue() {
	fmt.Println("Checking the queue")
	if len(q.players) > q.maxPlayers {
		match := NewMatch(q.players[0], q.players[1])
		q.players = q.players[q.maxPlayers:]
		q.matchmaking <- match
	}
}

func (q *Queue) StartMatch() {
	go func() {
		for match := range q.matchmaking {
			q.processMatch(match)
		}
	}()
}

func (q *Queue) processMatch(match *Match) {
	fmt.Println("Match found:", match)
	fmt.Printf("The navy player is %s\n", match.Navy.Username)
	fmt.Printf("The seals player is %s\n", match.Seals.Username)
}
