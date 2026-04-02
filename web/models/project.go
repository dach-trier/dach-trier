package models

import "html/template"

type Project struct {
	Preview     *Image
	Name        template.HTML
	Description template.HTML
}
