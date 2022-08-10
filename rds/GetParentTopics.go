package rds

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/topichierarchydao"
	"MyStackoverflow/dao/topicsdao"
	"MyStackoverflow/model"
)

func GetParentTopics() {

	sql := dao.MyDB.Table(topicsdao.TableTopics)
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
		key := FormParentsKey(tid)
		_ = DeleteKey(key)
		for _, p := range parentTids {
			_ = RPush(key, p)
		}
	}
	//var cursor uint64
	//topics := make([]string, 0)
	//keys, cursor, _ := RedisClient.Scan(cursor, "*:parents", 100).Result()
	//for _, key := range keys {
	//	parents, _ := RedisClient.LRange(key, 0, -1).Result()
	//	topics = append(topics, strings.Join(parents, ","))
	//}
	//topicNames := strings.Join(topics, ";")
	//fmt.Println(topicNames)
}
