package main

var regionRU = regionDef{
	Domain: "sh-development.ru",
	Lang:   "ru",

	Text: pageText{
		Login: "войти",
		About: "Конструктор QR-кодов платформы sh-development",

		TabCodes:  "qcode",
		TabAI:     "ии",
		TabManual: "вручную",
		Apps:      "apps",
		Readme:    "readme",
		Logout:    "выйти",

		Create:     "Создать",
		EmptyTitle: "Пока ничего не сохранено.",
		EmptyHint:  "Нажмите <b>Создать</b>, чтобы сделать первый код.",

		Name:         "Название",
		Unsaved:      "несохранённые изменения",
		Save:         "Сохранить",
		DownloadPNG:  "Скачать PNG",
		DownloadWebP: "Скачать WebP",
		DownloadSVG:  "Скачать SVG",
		CopySVG:      "Копировать SVG",
		Randomize:    "Случайно",
		Reset:        "Сбросить",
		RasterScale:  "Масштаб растра",

		AIStep1:          "Скопируйте инструкцию — в ней все параметры и текущий код.",
		AIStep2:          "Вставьте её в свой ИИ и опишите, какой код нужен.",
		AIStep3:          "Вставьте JSON из ответа ниже и примените.",
		CopyInstructions: "Копировать инструкцию",
		Apply:            "Применить",
		Images:           "Изображения",
		Logo:             "Логотип",
		Background:       "Фон",
		Choose:           "Выбрать",
		Clear:            "Убрать",

		JS: jsText{
			Copied:           "Скопировано",
			ClipboardBlocked: "Буфер обмена недоступен",
			CopySVG:          "Копировать SVG",
			Applied:          "Применено",
			Uploaded:         "загружено",
			None:             "нет",
			Edit:             "изменить",
			Delete:           "удалить",
			UnsavedMsg:       "Несохранённые изменения в «%s». Сохранить?",
			Save:             "Сохранить",
			DontSave:         "Не сохранять",
			DeleteMsg:        "Удалить «%s»? Это нельзя отменить.",
			DeleteConfirm:    "Удалить",
			Cancel:           "Отмена",
			ExportNoSupport:  "Этот браузер не умеет экспортировать в %s. Скачайте SVG или PNG.",
			ExportFailed:     "Не удалось экспортировать в %s. Скачайте SVG.",
		},
	},
}
