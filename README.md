# collector

The collector library tries to scrape the research products of the GrupLAC pages.
# Features
- [x] File definitions using yaml
- [x] regular expressions to define each field of the research product
- [x] Filter `eq`, and `contains` to filter by and specific field. 
## Examples

The following is a simple example to scrape the content of doctoral thesis using a file definition.

```go
const (
	gruplacURL                 = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	managedWorksDefinitionPath = "examples/doctoral_thesis/doctoral_thesis_definition.yml"
)

func main() {

	api := collector.NewAPI()
	results, err := api.GetContentFromFileDefinition(gruplacURL, managedWorksDefinitionPath)

	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
``` 

The file definition looks like:
```yaml
name: Tesis de doctorado
path: "/html[1]/body[1]/table[66]/tbody[1]/tr"
parser:
  lookahead-manual: true
fields:
  -
    name: type
    ex: (<strong>).*(</strong>)
  -
    name: title
    ex: (</strong>).*(<br/>)

  - name: fecha_inicio
    ex: (Desde)(.*)(hasta)
  -
    name: tipo_orientacion
    ex: (orientación:)(.*)(<br>)
  -
    name: nombre_estudiante
    ex: (estudiante:)(.*)(Programa académico:)
  -
    name: programa_academico
    ex: (Programa académico:)(.*)(<br>)
  -
    name: numero_paginas
    ex: (Número de páginas:)(.*)(, Valoración:)
  -
    name: institucion
    ex: (Institución:)(.*)
  -
    name: autores
    ex: (Autores:)(.*)(,)

filters:
  -
    fn: eq(Tesis de doctorado)
    field: type

```

The result of the last pieces of code is map tha contains a key(the field defined in the yaml), and a value as result of the regular expression parsing : 
```text
fecha_inicio:  1 2019 
tipo_orientacion: 
nombre_estudiante:  Jeniffer Paola Gracia Rojas, 
programa_academico: 
type: Tesis de doctorado
title:  : Desarrollo de una alternativa sostenible para la producción de PHA¿s a partir de cultivos mixtos y fuentes de carbono renovables.Jeniffer Paola Gracia Rojas
numero_paginas:  0
institucion:  UNIVERSIDAD DISTRITAL FRANCISCO JOSÉ DE CALDAS
autores:  PAULO ALONSO GAONA GARCIA CARLOS ENRIQUE MONTENEGRO MARIN
```