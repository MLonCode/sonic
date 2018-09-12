package sonic

import (
	"hash/fnv"
	"math"

	"github.com/MLonCode/sonic/src/sound"
)

func Convert(scale sound.Scale, nodes []sonicNode) []sound.Note {
	var max uint32
	for _, n := range nodes {
		if n.Lenght > max {
			max = n.Lenght
		}
	}

	notes := make([]sound.Note, len(nodes))
	hash := fnv.New32a()

	for i, n := range nodes {
		duration := float64(n.Lenght)
		duration = toLog(float64(max), duration) * 0.25

		hash.Reset()
		hash.Write([]byte(n.Token))
		note := scale.Get(hash.Sum32())

		notes[i] = sound.Note{Note: note, Duration: duration}
	}

	return notes
}

func toLog(max, value float64) float64 {
	return (math.Log(value) - math.Log(0.1)) / (math.Log(max) - math.Log(0.1))
}