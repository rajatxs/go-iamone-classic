package config

const (
	PUBLIC_DIR = "public"
	PAGE_404   = PUBLIC_DIR + "/404.html"
	PAGE_500   = PUBLIC_DIR + "/500.html"

	TEMPLATE_EXT             = "hbs"
	TEMPLATE_DIR             = PUBLIC_DIR + "/x/templates"
	TEMPLATE_DEFAULT_CONTENT = "default" + "." + TEMPLATE_EXT
	TEMPLATE_DEFAULT_LAYOUT  = "default.layout" + "." + TEMPLATE_EXT
	TEMPLATE_DEFAULT_STYLE   = "default.style" + "." + TEMPLATE_EXT

	THEMES_EXT = "json"
	THEMES_DIR = PUBLIC_DIR + "/x/themes"
)
