package source

type Source interface {
	NextID() (uint16, error)
}
