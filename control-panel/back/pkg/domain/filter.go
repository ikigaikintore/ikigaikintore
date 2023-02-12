package domain

type (
	Filter struct {
		Name     string
		Birthday string
	}

	FilterOption func(*Filter)
)

func NewFilter(opts ...FilterOption) Filter {
	df := Filter{}
	for _, opt := range opts {
		opt(&df)
	}
	return df
}

func WithName(name string) FilterOption {
	return func(filter *Filter) {
		filter.Name = name
	}
}

func WithBirthday(birthday string) FilterOption {
	return func(filter *Filter) {
		filter.Birthday = birthday
	}
}
