package gosns

/*** Common ***/
type ResponseMetadata struct {
	RequestId string                `xml:"RequestId"`
}


/*** List Topics Response */
type TopicArnResult struct {
	TopicArn string				`xml:"TopicArn"`
}

type TopicNamestype struct {
	Member []TopicArnResult			`xml:"member"`
}

type  ListTopicsResult struct {
	Topics TopicNamestype                        `xml:"Topics"`
}

type ListTopicsResponse struct {
	Xmlns  		string  		`xml:"xmlns,attr"`
	Result		ListTopicsResult	`xml:"ListTopicsResult"`
	Metadata 	ResponseMetadata	`xml:"ResponseMetadata"`
}
