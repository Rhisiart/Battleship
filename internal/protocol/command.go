package protocol

/*
Message Types
	LOGIN: A client joins the game.
	LOGOUT: A client left the game.
	JOIN: A client join a queue.
	START: Start the game when enough players have joined.
	PLACE: Place a ship on the board.
	ATTACK: Attack a position on the opponent’s board.
	RESULT: The result of an attack.
	LEAVE: A client leaves the game.

Message Format
	LOGIN: LOGIN <username>
	LOGOUT: LOGOUT
	JOIN: JOIN <username>
	START: START
	PLACE: PLACE <x> <y> <direction> <length> (e.g., PLACE 1 1 H 5 for a horizontal ship of length 5 starting at (1,1))
	ATTACK: ATTACK <x> <y>
	RESULT: RESULT <hit/miss> <sunk> (e.g., RESULT hit true if the attack hit and sunk a ship)
	LEAVE: LEAVE <username>
*/

type ID int

const (
	LOGIN ID = iota
	LOGOUT
	CLIENTS
	JOIN
	START
	PLACE
	ATTACK
	RESULT
	LEAVE
)

type Command struct {
	id     ID
	sender string
	body   []byte
}
