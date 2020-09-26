package lgrs

import "github.com/hobord/logger"

func stringInSlice(sl []string, val string) bool {
	for _, item := range sl {
		if item == val {
			return true
		}
	}
	return false
}

func (l *lgrs) GetTags() []string {
	var tags []string
	field := l.base.Data["tags"]
	switch v := field.(type) {
	case []string:
		tags = v
	default:
		tags = make([]string, 0)
	}

	return tags
}

func (l *lgrs) addNewTags(tags []string) []string {
	currentTags := l.GetTags()

	for _, tag := range tags {
		if stringInSlice(currentTags, tag) == false {
			currentTags = append(currentTags, tag)
		}
	}
	return currentTags
}

func (l *lgrs) WithTags(tags []string) logger.Logger {
	currentTags := l.addNewTags(tags)
	entry := l.base.WithField("tags", currentTags)

	// l.base = entry
	// return l

	lg := l.cloneLgrs()
	lg.base = entry
	return lg
}

func (l *lgrs) SetDefaultTags(tags []string) {
	currentTags := l.addNewTags(tags)
	l.base = l.base.WithField("tags", currentTags)
}
