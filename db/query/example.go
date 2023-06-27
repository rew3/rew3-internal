package query

func test() {
	Query(
		Field("first_name").Eq("raka"),
		Field("last_name").NOT().Eq("raka"),
		AND(
			Field("first_name").Eq("raka"),
			OR(
				Field("last_name").Eq("raka"),
				Field("gender").Eq("raka"),
			),
		),
	)
}