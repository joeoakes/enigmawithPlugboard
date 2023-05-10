package main

import (
	"fmt"
	"strings"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Rotor struct {
	wiring         string
	position       int
	notchPositions []int
}

func (r *Rotor) shift() {
	r.position = (r.position + 1) % len(alphabet)
}

func (r *Rotor) notch() bool {
	for _, notch := range r.notchPositions {
		if r.position == notch {
			return true
		}
	}
	return false
}

func (r *Rotor) encrypt(input int) int {
	offset := (input + r.position) % len(alphabet)
	output := (offset + int(r.wiring[offset]-'A')) % len(alphabet)
	return (output - r.position + len(alphabet)) % len(alphabet)
}

type Plugboard struct {
	pairs map[int]int
}

func (p *Plugboard) encrypt(input int) int {
	output, exists := p.pairs[input]
	if exists {
		return output
	}
	return input
}

type Enigma struct {
	rotors    []*Rotor
	plugboard *Plugboard
}

func (e *Enigma) encrypt(input string) string {
	input = strings.ToUpper(input)
	output := ""
	for _, char := range input {
		inputIndex := int(char - 'A')
		inputIndex = e.plugboard.encrypt(inputIndex)
		for _, rotor := range e.rotors {
			inputIndex = rotor.encrypt(inputIndex)
		}
		inputIndex = e.plugboard.encrypt(inputIndex)
		output += string(alphabet[inputIndex])
		for _, rotor := range e.rotors {
			if rotor.notch() {
				rotor.shift()
			}
		}
	}
	return output
}

func main() {
	// Define the rotors
	rotor1 := &Rotor{wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", position: 0, notchPositions: []int{16}}
	rotor2 := &Rotor{wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE", position: 0, notchPositions: []int{4}}
	rotor3 := &Rotor{wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO", position: 0, notchPositions: []int{21}}
	rotors := []*Rotor{rotor1, rotor2, rotor3}

	// Define the plugboard settings
	//0:24 a:y, 2:8 c:i, 4:12 e:m, 6:18 g:s, 10:17 k:r, 11:22 l:w,13:16 n:q,
	//14:19 o:t, 20:23 u:x, 21:25 v:z
	plugboardPairs := map[int]int{
		0:  24,
		2:  8,
		4:  12,
		6:  18,
		10: 17,
		11: 22,
		13: 16,
		14: 19,
		20: 23,
		21: 25,
	}
	plugboard := &Plugboard{pairs: plugboardPairs}

	// Create the Enigma machine
	enigma := &Enigma{rotors, plugboard}

	// Encrypt a message
	message := "HELLO"
	encrypted := enigma.encrypt(message)
	fmt.Println("Encrypted message:", encrypted)
	// decrypt a message
	decrypted := enigma.encrypt("QYYBY")
	fmt.Println("decrypted message:", decrypted)
}
