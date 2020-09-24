package fixtures

const DoctoralThesisContent = `name: Tesis de doctorado
path: "/html[1]/body[1]/table[66]/tbody[1]/tr"
parsing:
  lookahead-manual: true
fields:
  -
    name: type
    id: Tesis de doctorado
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
`
const DoctoralThesisDefinitionName = "Tesis de doctorado"
