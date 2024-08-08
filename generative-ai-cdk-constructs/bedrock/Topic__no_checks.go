//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Topic) validateCreateTopicParameters(props *TopicProps) error {
	return nil
}

func (t *jsiiProxy_Topic) validateFinancialAdviceTopicParameters(props *TopicProps) error {
	return nil
}

func (t *jsiiProxy_Topic) validateInappropriateContentParameters(props *TopicProps) error {
	return nil
}

func (t *jsiiProxy_Topic) validateLegalAdviceParameters(props *TopicProps) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMedicalAdviceParameters(props *TopicProps) error {
	return nil
}

func (t *jsiiProxy_Topic) validatePoliticalAdviceTopicParameters(props *TopicProps) error {
	return nil
}

func validateTopic_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTopicParameters(scope constructs.Construct, id *string) error {
	return nil
}

