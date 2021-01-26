package resolver

// LiveStreamFeedResolver ...
type LiveStreamFeedResolver struct {
	Data         *[]*LivestreamResolver
	CursorHolder *string
}

func (l *LiveStreamFeedResolver) List() *[]*LivestreamResolver {
	return l.Data
}

func (l *LiveStreamFeedResolver) Cursor() *string {
	return l.CursorHolder
}
