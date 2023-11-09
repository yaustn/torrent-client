package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	return
}

// // Open parses a torrent file
// func Open(r io.Reader) (*bencodeTorrent, error) {
// 	bto := bencodeTorrent{}
// 	err := bencode.Unmarshal(r, &bto)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &bto, nil
// }
