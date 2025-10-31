package persistence

// ConvertSliceToDomain converts a slice of persistence models to domain entities
// This eliminates the repeated loop-and-convert pattern in every repository
//
// Usage:
//   tags, err := ConvertSliceToDomain(models, ports.TagPersistenceToDomainEntity)
func ConvertSliceToDomain[TModel any, TDomain any](
	models []TModel,
	converter func(*TModel) (TDomain, error),
) ([]TDomain, error) {
	result := make([]TDomain, len(models))
	for i, model := range models {
		entity, err := converter(&model)
		if err != nil {
			return nil, err
		}
		result[i] = entity
	}
	return result, nil
}

// ConvertSliceToDomainPtr converts a slice of persistence models to pointers of domain entities
// This is for cases where the converter returns a pointer
//
// Usage:
//   tags, err := ConvertSliceToDomainPtr(models, ports.TagPersistenceToDomainEntity)
func ConvertSliceToDomainPtr[TModel any, TDomain any](
	models []TModel,
	converter func(*TModel) (*TDomain, error),
) ([]*TDomain, error) {
	result := make([]*TDomain, len(models))
	for i, model := range models {
		entity, err := converter(&model)
		if err != nil {
			return nil, err
		}
		result[i] = entity
	}
	return result, nil
}
