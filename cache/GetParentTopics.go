package cache

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
)

// ParentTopics topic and all its parent topics
var ParentTopics map[int][]int

func GetParentTopics() map[int][]int {

	sql := dao.MyDB.Table(topicsdao.TableTopics)
	parentTopics := make(map[int][]int)
	allTopics := make([]*model.Topic, 0)
	sql.Find(&allTopics)
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
		parentTopics[tid] = parentTids
	}
	return parentTopics
}
