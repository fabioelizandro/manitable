package manitable

type Transform func(Table) Table

type Pipeline struct {
	listOfTransformations []Transform
}

func NewPipeline(listOfTransformations []Transform) Pipeline {
	return Pipeline{listOfTransformations: listOfTransformations}
}

func (p Pipeline) Run(table Table) Table {
	for _, transformation := range p.listOfTransformations {
		table = transformation(table)
	}

	return table
}
