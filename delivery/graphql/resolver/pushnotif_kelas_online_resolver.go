package resolver

import "github.com/nugrohoac/livestream/entity"

type PushNotifResolver struct {
	Data *entity.PushNotif
}

func (p *PushNotifResolver) Title() *string {
	return &p.Data.Title
}

func (p *PushNotifResolver) Description() *string {
	return &p.Data.Description
}
