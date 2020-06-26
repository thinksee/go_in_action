package mapreduce

import (
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
)

func doMap(
	jobName string,
	mapTask int,
	inFile string,
	nReduce int,
	mapF func(filename string, contents string) []KV,
) {

	encoders := make([]*json.Encoder, nReduce)
	for encoder := range encoders {
		file, err := os.Create(reduceName(jobName, mapTask, encoder))
		if err != nil {
			log.Fatal(err)
		}
		encoders[encoder] = json.NewEncoder(file)
		defer file.Close()
	}

	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	kvs := mapF(inFile, string(contents))
	for _, kv := range kvs {
		r := ihash(kv.Key) % nReduce
		err := encoders[r].Encode(kv)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7ffffff)
}
