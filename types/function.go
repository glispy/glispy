package types

// Function represents a function type
type Function func(sc Scope, args List) (Expression, error)
