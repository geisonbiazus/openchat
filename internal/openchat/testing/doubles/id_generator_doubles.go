package doubles

type IDGeneratorStub struct {
	ID string
}

func NewIDGeneratorStub(id string) IDGeneratorStub {
	return IDGeneratorStub{
		ID: id,
	}
}

func (g IDGeneratorStub) Generate() string {
	return g.ID
}
