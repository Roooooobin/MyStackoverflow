package cache

func Init() {

	ParentTopics = GetParentTopics()
	Topic2SubTopics = StoreSubTopics()
}
