// https://adventofcode.com/2020/day/10
package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const tilesize = 10

type tile struct {
	content      map[image.Point]byte
	orientations [8]orientation
}

type orientation struct {
	top, bottom, left, right     [tilesize]byte
	tops, bottoms, lefts, rights []key
}

type key struct {
	id, orientationID int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	tileStrings := strings.Split(string(input), "\n\n")
	tiles := make(map[int]*tile)

	for i := range tileStrings {
		lines := strings.Split(tileStrings[i], "\n")
		var id int
		fmt.Sscanf(lines[0], "Tile %d:", &id)
		t := &tile{content: make(map[image.Point]byte)}
		for y := 0; y < tilesize; y++ {
			line := lines[y+1]
			for x := 0; x < tilesize; x++ {
				if y == 0 {
					t.orientations[0].top[x] = line[x]
				} else if y == tilesize-1 {
					t.orientations[0].bottom[x] = line[x]
				} else if x > 0 && x < tilesize-1 {
					t.content[image.Pt(x-1, y-1)] = line[x]
				}
			}
			t.orientations[0].left[y] = line[0]
			t.orientations[0].right[y] = line[tilesize-1]
		}
		tiles[id] = t
	}

	for id := range tiles {
		for o := 1; o < 4; o++ {
			tiles[id].orientations[o] = rotateOrientation(tiles[id].orientations[o-1])
		}
		tiles[id].orientations[4] = flipOrientation(tiles[id].orientations[0])
		for o := 5; o < 8; o++ {
			tiles[id].orientations[o] = rotateOrientation(tiles[id].orientations[o-1])
		}
	}

	for id1 := range tiles {
		for id2 := range tiles {
			if id1 == id2 {
				continue
			}
			for i1 := range tiles[id1].orientations {
				for i2 := range tiles[id2].orientations {
					if tiles[id1].orientations[i1].left == tiles[id2].orientations[i2].right {
						tiles[id1].orientations[i1].lefts = append(tiles[id1].orientations[i1].lefts, key{id2, i2})
					}
					if tiles[id1].orientations[i1].right == tiles[id2].orientations[i2].left {
						tiles[id1].orientations[i1].rights = append(tiles[id1].orientations[i1].rights, key{id2, i2})
					}
					if tiles[id1].orientations[i1].top == tiles[id2].orientations[i2].bottom {
						tiles[id1].orientations[i1].tops = append(tiles[id1].orientations[i1].tops, key{id2, i2})
					}
					if tiles[id1].orientations[i1].bottom == tiles[id2].orientations[i2].top {
						tiles[id1].orientations[i1].bottoms = append(tiles[id1].orientations[i1].bottoms, key{id2, i2})
					}

				}
			}
		}
	}

	// Part 1
	corners := make(map[int]struct{})
	for id := range tiles {
		for _, o := range tiles[id].orientations {
			if len(o.tops) == 0 && len(o.lefts) == 0 {
				corners[id] = struct{}{}
			}
		}
	}
	prod := 1
	for c := range corners {
		prod *= c
	}
	fmt.Println(prod)

	// Part 2
	// Assuming each tile from the input has exactly zero or one match on any side.
	id, o := anyTopLeft(tiles)
	img := make(map[image.Point]byte)
	var origin image.Point
	for {
		leftID, leftO := id, o
		for {
			c := getContent(tiles, id, o)
			for pt := range c {
				img[pt.Add(origin)] = c[pt]
			}
			origin = origin.Add(image.Pt(tilesize-2, 0))
			if len(tiles[id].orientations[o].rights) == 0 {
				break
			}
			k := tiles[id].orientations[o].rights[0]
			id, o = k.id, k.orientationID
		}
		if len(tiles[leftID].orientations[leftO].bottoms) == 0 {
			break
		}
		k := tiles[leftID].orientations[leftO].bottoms[0]
		id, o = k.id, k.orientationID
		origin = image.Pt(0, origin.Y+tilesize-2)
	}
	imgSize := origin.Y + tilesize - 2

	monster := getMonster()
	for o := 0; o < 4; o++ {
		if roughness, found := findMonster(img, monster, imgSize); found {
			fmt.Println(roughness)
			return
		}
		img = rotate(img, imgSize)
	}
	img = flip(img, imgSize)
	for o := 4; o < 8; o++ {
		if roughness, found := findMonster(img, monster, imgSize); found {
			fmt.Println(roughness)
			return
		}
		img = rotate(img, imgSize)
	}
}

func rotateOrientation(in orientation) orientation {
	var out orientation
	out.top = reverse(in.left)
	out.bottom = reverse(in.right)
	out.left = in.bottom
	out.right = in.top
	return out
}

func flipOrientation(in orientation) orientation {
	var out orientation
	out.top = reverse(in.top)
	out.bottom = reverse(in.bottom)
	out.left = in.right
	out.right = in.left
	return out
}

func reverse(in [tilesize]byte) [tilesize]byte {
	var out [tilesize]byte
	for i := 0; i < tilesize; i++ {
		out[tilesize-i-1] = in[i]
	}
	return out
}

func anyTopLeft(tiles map[int]*tile) (int, int) {
	for id := range tiles {
		for i, o := range tiles[id].orientations {
			if len(o.tops) == 0 && len(o.lefts) == 0 {
				return id, i
			}
		}
	}
	return -1, -1
}

func getContent(tiles map[int]*tile, id, orientationID int) map[image.Point]byte {
	content := tiles[id].content
	transformed := make(map[image.Point]byte)
	for pt := range content {
		transformed[pt] = content[pt]
	}
	switch orientationID {
	case 0:
	case 1, 2, 3:
		for i := 0; i < orientationID; i++ {
			transformed = rotate(transformed, tilesize-2)
		}
	case 4:
		transformed = flip(transformed, tilesize-2)
	case 5, 6, 7:
		transformed = flip(transformed, tilesize-2)
		for i := 0; i < orientationID-4; i++ {
			transformed = rotate(transformed, tilesize-2)
		}
	}
	return transformed
}

func rotate(tile map[image.Point]byte, size int) map[image.Point]byte {
	rotated := make(map[image.Point]byte)
	for pt := range tile {
		rotated[image.Pt(size-1-pt.Y, pt.X)] = tile[pt]
	}
	return rotated
}
func flip(tile map[image.Point]byte, size int) map[image.Point]byte {
	flipped := make(map[image.Point]byte)
	for pt := range tile {
		flipped[image.Pt(size-1-pt.X, pt.Y)] = tile[pt]
	}
	return flipped
}

func getMonster() map[image.Point]byte {
	monsterStr := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `
	monster := make(map[image.Point]byte)
	for y, l := range strings.Split(monsterStr, "\n") {
		for x := range l {
			if l[x] == '#' {
				monster[image.Pt(x, y)] = '#'
			}
		}
	}
	return monster
}

func findMonster(img, monster map[image.Point]byte, imgSize int) (int, bool) {
	var monsterLocations []image.Point
	for x := 0; x < imgSize; x++ {
		for y := 0; y < imgSize; y++ {
			possible := true
			for pt := range monster {
				if img[pt.Add(image.Pt(x, y))] != monster[pt] {
					possible = false
					break
				}
			}
			if possible {
				monsterLocations = append(monsterLocations, image.Pt(x, y))
			}
		}
	}

	if len(monsterLocations) == 0 {
		return 0, false
	}
	numHash := 0
	for pt := range img {
		if img[pt] == '#' {
			numHash++
		}
	}
	return numHash - len(monsterLocations)*len(monster), true
}
