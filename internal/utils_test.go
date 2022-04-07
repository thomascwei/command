package internal

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path"
	"testing"
)

func TestInit(t *testing.T) {
	Trace.Println("start test initial")
	file, err := ioutil.ReadFile(path.Join(rootPath, "test_data/temp.json"))
	require.NoError(t, err)
	var data []WriteCommandTemplate
	_ = json.Unmarshal(file, &data)
	docs := make([]interface{}, len(data))
	for i, v := range data {
		docs[i] = v
	}
	// insert前先清空
	CommandCollection.Drop(Ctx)
	// insert
	_, err = CommandCollection.InsertMany(Ctx, docs)
	require.NoError(t, err)


}
