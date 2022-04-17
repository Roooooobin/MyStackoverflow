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
		var par int
		curTopic, err := topichierarchydao.Find("tid = ?", rootTid)
		if err != nil {
			break
		}
		if curTopic != nil {
			par = curTopic.ParentTid
		}
		parentTids := make([]int, 0)
		parentTids = append(parentTids, rootTid)
		// when rootTid == par, it is the top topic
		for rootTid != par {
			parentTids = append(parentTids, par)
			t, err := topichierarchydao.Find("tid = ?", par)
			if err != nil {
				break
			}
			rootTid = par
			if t != nil {
				par = t.ParentTid
			}
		}
		ParentTopics[tid] = parentTids
	}
}
