package useronline

type Payload struct {
	HostName string `json:"hostName"`
	SocketID string `json:"socketId"`
	UserID   string `json:"id"`
}

func (p *Payload) convertToModel() UserOnline {
	u := UserOnline{
		UserID: p.UserID,
	}
	return u
}
