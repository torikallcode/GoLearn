package seed

func foundationPaths(lessons map[string]*lessonFile) []seedPathDef {
	module := func(title string, slugs ...string) seedModuleDef {
		m := seedModuleDef{Title: title}
		for _, slug := range slugs {
			if l, ok := lessons[slug]; ok {
				m.Lessons = append(m.Lessons, l)
			}
		}
		return m
	}
	return []seedPathDef{{
		Slug:        "fondasi-go",
		Title:       "Fondasi Go",
		Description: "Mulai dari nol: sintaks, tipe data, fungsi, struct, pointer, interface, sampai error handling dan struktur package.",
		Level:       "Pemula",
		Order:       1,
		Modules: []seedModuleDef{
			module("Dasar Go",
				"pengenalan-go",
				"variabel-dan-tipe-data",
				"operator-dan-control-flow",
			),
			module("Fungsi & Koleksi",
				"fungsi",
				"array-slice-map",
				"string-rune-strconv",
			),
			module("Struktur & Abstraksi",
				"struct-dan-method",
				"pointer",
				"interface-dan-polimorfisme",
				"error-handling",
				"package-dan-modul",
			),
		},
	}}
}

func skeletonPaths() []seedPathDef {
	return []seedPathDef{
		{
			Slug:        "go-menengah",
			Title:       "Go Menengah",
			Description: "Generics, goroutine, channel, sync, context, JSON, testing, dan standard library penting. Segera hadir.",
			Level:       "Menengah",
			Order:       2,
		},
		{
			Slug:        "backend-go",
			Title:       "Backend Engineering dengan Go",
			Description: "REST API dengan Gin, MySQL, auth JWT, caching, Docker, CI/CD, dan observability. Segera hadir.",
			Level:       "Lanjut",
			Order:       3,
		},
		{
			Slug:        "profesional-go",
			Title:       "Profesional & Siap Kerja",
			Description: "Clean architecture, concurrency patterns, gRPC, keamanan, profiling, dan persiapan interview. Segera hadir.",
			Level:       "Profesional",
			Order:       4,
		},
	}
}
