package cache

import (
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/dao/topicsdao"
)

// ParentTopics topic and all its parent topics
var ParentTopics map[int][]int

func Init() {

	ParentTopics = make(map[int][]int)
	allTopics, _ := topicsdao.List("tid > ?", 0)
	for _, topic := range allTopics {
		rootTid := topic.Tid
		tid := rootTid
		parentTids := make([]int, 0)
		for rootTid != 0 {
			parentTids = append(parentTids, rootTid)
			t, err := topichierarchydao.Find("tid = ?", rootTid)
			if err != nil {
				break
			}
			if t != nil {
				rootTid = t.ParentTid
			}
		}
		ParentTopics[tid] = parentTids
	}
}
