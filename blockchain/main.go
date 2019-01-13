/*
 *		Blockchain
 *			GO
 *		Marcos Huck
 *	 marcos@huck.com.ar
 *
 */

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	id        int
	timeStamp string
	bpm       int
	hash      string
	hashPtr   string
}

var Blockchain []Block

func calcularHash(block Block) string {
	registro := string(block.id) + block.timeStamp + string(block.bpm) + block.hashPtr
	hash := sha256.New()
	hash.Write([]byte(registro))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generarBloque(anteriorBloque Block, bpm int) (Block, error) {
	var nuevoBloque Block
	nuevoBloque.id = anteriorBloque.id + 1
	nuevoBloque.timeStamp = time.Now().String()
	nuevoBloque.bpm = bpm
	nuevoBloque.hashPtr = anteriorBloque.hash
	nuevoBloque.hash = calcularHash(nuevoBloque)
	return nuevoBloque, nil
}

func validarBloque(nuevoBloque, anteriorBloque Block) bool {
	if anteriorBloque.id != nuevoBloque.id+1 {
		return false
	}

	if anteriorBloque.hash != nuevoBloque.hashPtr {
		return false
	}

	if calcularHash(nuevoBloque) != nuevoBloque.hash {
		return false
	}
	return true
}
