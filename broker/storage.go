package broker

type TopicStorage struct {
	Messages map[string][][]byte
}

func NewTopicStorage() *TopicStorage {
	messages := make(map[string][][]byte)
	ts := TopicStorage{
		Messages: messages,
	}
	return &ts
}

func (ts *TopicStorage) AddMessage(topic string, message []byte) {
	ts.Messages[topic] = append(ts.Messages[topic], message)
}

func (ts *TopicStorage) GetMessages(topic string) [][]byte {
	return ts.Messages[topic]
}

func (ts *TopicStorage) ListTopics() []string {
	var topics []string

	// Iterate over the map to get the topics
	for topic := range ts.Messages {
		topics = append(topics, topic)
	}

	return topics
}
