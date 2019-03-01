package engine

import(
	"hash/fnv"
)

func GetHash(s string) uint64{
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}