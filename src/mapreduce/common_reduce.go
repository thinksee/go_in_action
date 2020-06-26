package mapreduce

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTask int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	kvs := make(map[string]([]string))

	for m := 0; m < nMap; m++ {
		file, err := os.Open(reduceName(jobName, m, reduceTask))
		if err != nil {
			// printf + exit
			log.Fatal(err)
		}

		// json.Decoder，若数据是从一个即将io.Reader流，或者需要多个值，从数据流进行解码
		// json.Unmarshal，若JSON数据已经在内存中存在了
		decoder := json.NewDecoder(file)
		for decoder.More() {
			var kv KV
			if err := decoder.Decode(&kv); err != nil {
				log.Fatal(err)
			}
			kvs[kv.Key] = append(kvs[kv.Key], kv.Value)
		}
		defer file.Close()
	}

	keys := make([]string, 0, len(kvs))
	for key, _ := range kvs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	file, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, key := range keys {
		encoder.Encode(KV{key, reduceF(key, kvs[key])})
	}
}
