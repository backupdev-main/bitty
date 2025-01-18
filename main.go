package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peterkwesiansah/bitty/bencodeTorrent"
	"github.com/peterkwesiansah/bitty/peers"
	"github.com/peterkwesiansah/bitty/torrentfile"
)

func main() {

	srcPath := flag.String("src", "", "Path to the source torrent file")
	dstPath := flag.String("dst", "", "Destination path for the downloaded torrent content")

	flag.Parse()
	if *srcPath == "" || *dstPath == "" {
		fmt.Println("run instead: go run main.go -src /path/to/source -dst /path/to/destination")
		flag.Usage()
		os.Exit(1)
	}
	bct, err := bencodeTorrent.Decode(*srcPath)
	if err != nil {
		log.Fatal(err)
	}
	p2p, err := peers.Peers(bct)

	if err != nil {
		log.Fatal(err)
	}

	tf := torrentfile.Torrent{
		Peers:       p2p.Peers,
		PeerID:      [20]byte(p2p.PeerId),
		InfoHash:    bct.InfoHash,
		Name:        bct.Info.Name,
		PieceLength: bct.Info.PieceLength,
		Length:      bct.Info.Length,
		PieceHashes: bct.PieceHashes,
	}

	file, err := tf.Download(*dstPath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
