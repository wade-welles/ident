package dynamodb

type Source struct {
	ClusterName, TableName string
}

func (s *Source) NextID() (uint16, error) {
	counter := Counter{
		Cluster: s.ClusterName,
	}

	table := NewTable(s.TableName)

	if err := table.Increment(&counter); err != nil {
		return 0, err
	}

	return counter.Seq, nil
}

func New(clusterName, tableName string) *Source {
	return &Source{
		ClusterName: clusterName,
		TableName:   tableName,
	}
}
