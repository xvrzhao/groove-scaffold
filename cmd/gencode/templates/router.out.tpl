// {{.upperCamelModel}}s
{{.lowerCamelModel}} := <parent_router>.Group("/{{.lowerCamelModel}}s")
{
  {{.lowerCamelModel}}Ctrl := handler.New{{.upperCamelModel}}Ctrl()
  {{.lowerCamelModel}}.GET("", {{.lowerCamelModel}}Ctrl.Page)
  {{.lowerCamelModel}}.POST("", {{.lowerCamelModel}}Ctrl.Create)
  {{.lowerCamelModel}}.PUT("/:id", {{.lowerCamelModel}}Ctrl.Update)
  {{.lowerCamelModel}}.DELETE("/:id", {{.lowerCamelModel}}Ctrl.Delete)
}