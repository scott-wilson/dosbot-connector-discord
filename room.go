package discord

type room struct {
	name string
	id   string
}

func newRoom(name string, id string) room {
	return room{name: name, id: id}
}

func (r room) Name() string {
	return r.name
}

func (r room) ID() interface{} {
	return r.id
}
