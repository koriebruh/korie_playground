package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Identity struct {
	Name    string
	Address string
	Region  string
}

// CONDITION BASIC NO POINTER
func TestNoPointer(t *testing.T) {
	user1 := Identity{
		Name:    "Ipur",
		Address: "Semarang",
		Region:  "Indonesia",
	}
	user2 := user1

	assert.Equal(t, user1, user2)
}

// CONDITION Ref POINTER
func TestPointer(t *testing.T) {
	user1 := Identity{
		Name:    "Ipur",
		Address: "Semarang",
		Region:  "Indonesia",
	}
	//JADI PERUBAHAN AKAN MEMPENGARUHI SATU SAMA LAIN
	var user2 *Identity = &user1
	//user2 := &user1
	user2.Name = "NinaAsu"

	fmt.Println(user1)
	fmt.Println(user2)

	assert.Equal(t, user1, user2)
}

// NAMUN JIKA KITA MENDEKLARASIKAN ULANG DENGAN MENGUBAH SELURUH FIELD MAKA PERUBAHAN  TERSEBUT TIDAK AKAN SALING MEMPENGARUHI
func TestPointerNotEffect(t *testing.T) {
	user1 := Identity{
		Name:    "RebertusEntod",
		Address: "ANJ",
		Region:  "Indonesia",
	}

	user3 := &user1
	var user2 *Identity = &user1
	user2.Name = "NinaAsuMemek"

	// KARENA DI DEKLARASIKAN ULANG 1 FILED PENUH JADI YG REFENRECE TIDAK TERPENGARUH
	user2 = &Identity{
		Name:    "nina nngendot ama robertus",
		Address: "dosen anjg",
		Region:  "Forest",
	}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

	assert.NotEqual(t, user1, &user3)
	assert.NotEqual(t, user1, user2)
	assert.NotEqual(t, user3, user2)

}

// LALU BAGIMANA BG BG BIAR TETEP TERPENGARUH
func TestPaksa(t *testing.T) {
	user1 := Identity{
		Name:    "RebertusEntod",
		Address: "ANJ",
		Region:  "Indonesia",
	}
	var user2 *Identity = &user1
	//DENGAM MENAMBHAKAN * MEMAKSA AGAR YG TERKAIT IKUT BRUBAH
	*user2 = Identity{
		Name:    "nina nngendot ama robertus",
		Address: "dosen anjg",
		Region:  "Forest",
	}

	fmt.Println(user1)
	fmt.Println(user2)
	assert.Equal(t, &user1, user2)

}
