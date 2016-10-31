package source

type Source interface {
	NextID() (uint16, error)
}

func MustID(source Source) uint16 {
	id, err := source.NextID()
	if err != nil {
		panic(err)
	}

	return id
}
