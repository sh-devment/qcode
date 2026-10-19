package main

var regionCOM = regionDef{
	Domain: "sh-development.com",
	Lang:   "en",

	Text: pageText{
		Login: "log in",
		About: "The QR code builder of the sh-development platform.",

		TabCodes:  "qcode",
		TabAI:     "ai",
		TabManual: "manual",
		Apps:      "apps",
		Readme:    "readme",
		Logout:    "log out",

		Create:     "Create",
		EmptyTitle: "Nothing saved yet.",
		EmptyHint:  "Press <b>Create</b> to build your first code.",

		Name:         "Name",
		Unsaved:      "unsaved changes",
		Save:         "Save",
		DownloadPNG:  "Download PNG",
		DownloadWebP: "Download WebP",
		DownloadSVG:  "Download SVG",
		CopySVG:      "Copy SVG",
		Randomize:    "Randomize",
		Reset:        "Reset",
		RasterScale:  "Raster scale",

		AIStep1:          "Copy the instructions — every parameter, plus this code as it stands.",
		AIStep2:          "Paste them into your own AI and describe the code you want.",
		AIStep3:          "Paste the JSON it returns below and apply it.",
		CopyInstructions: "Copy instructions",
		Apply:            "Apply",
		Images:           "Images",
		Logo:             "Logo",
		Background:       "Background",
		Choose:           "Choose",
		Clear:            "Clear",

		JS: jsText{
			Copied:           "Copied",
			ClipboardBlocked: "Clipboard blocked",
			CopySVG:          "Copy SVG",
			Applied:          "Applied",
			Uploaded:         "uploaded",
			None:             "none",
			Edit:             "edit",
			Delete:           "delete",
			UnsavedMsg:       "Unsaved changes in “%s”. Save them?",
			Save:             "Save",
			DontSave:         "Don't save",
			DeleteMsg:        "Delete “%s”? This cannot be undone.",
			DeleteConfirm:    "Delete",
			Cancel:           "Cancel",
			ExportNoSupport:  "%s export is not supported by this browser. Download the SVG or PNG instead.",
			ExportFailed:     "%s export failed. Download the SVG instead.",
		},
	},
}
