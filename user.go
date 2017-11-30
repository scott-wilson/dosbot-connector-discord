package discord

type user struct {
	name string
	id   string
}

func newUser(name string, id string) user {
	return user{name: name, id: id}
}

func (u user) Name() string {
	return u.name
}

func (u user) ID() interface{} {
	return u.id
}
