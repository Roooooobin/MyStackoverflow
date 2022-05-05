package cache

func Init() {

	ParentTopics = GetParentTopics()
	TopicID2Name = GetTopicNameByID()
	Topic2SubTopics = StoreSubTopics()
}
