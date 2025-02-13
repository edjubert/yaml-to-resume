---
documentclass: extarticle
title: {{ .Info.Name }}
geometry: margin=1cm
fontsize: 9pt
header-includes: |
  \usepackage{fontspec}
  \setmainfont{DejaVu Sans}
  \setmonofont{DejaVu Sans Mono}
  \usepackage{graphicx}
  \usepackage{wrapfig}
  \usepackage[export]{adjustbox}
  \usepackage{xcolor}
  \usepackage{hyperref}
  \usepackage[absolute]{textpos}

  \newlength{\logoheight}
  \setlength{\logoheight}{0.8em}

  \newlength{\defaultleftskip}
  \setlength{\defaultleftskip}{0.1cm}

  \newlength{\defaultvspace}
  \setlength{\defaultvspace}{0.2cm}

  \newlength{\imagewidth}
  \setlength{\imagewidth}{2.5cm}

  \setlength{\parindent}{0pt}
  \setlength{\topskip}{0pt}
  \setlength{\parskip}{1pt}

  \pagenumbering{gobble}
  \makeatletter
  \renewcommand{\maketitle}{\begin{flushleft}
    \huge\@title
  \end{flushleft}}
  \makeatother
---

\begin{textblock*}{3.5cm}(18.5cm,1cm)
\adjustimage{width=\imagewidth,clip,raise=20pt}{ {{- .Info.ProfilePicture -}} }
\end{textblock*}

\vspace{\defaultvspace}

\textbf{ {{- .Info.Title -}} }

\vspace{0cm}
{{- if .Info.Location -}}
\textcolor{gray}{ {{- .Info.Location -}} }
{{- end }}

\vspace{\defaultvspace}

{{- if or .Info.Email .Info.GitHub .Info.Linkedin .Info.Malt -}}

{{- if .Info.Email -}}
\small \href{mailto:{{ .Info.Email }}}{\textcolor{black}{\includegraphics[height=\logoheight]{./assets/logos/email.png} {{ .Info.Email }}}}{{- if or .Info.GitHub .Info.Linkedin .Info.Malt -}} \space|\space {{- end -}}
{{- end -}}

{{- if .Info.GitHub -}}
\small [\includegraphics[height=\logoheight]{./assets/logos/github.png} \textcolor{black}{Github}](https://github.com/{{ .Info.GitHub }})
{{- if or .Info.Linkedin .Info.Malt -}}\space|\space{{- end -}}
{{- end -}}

{{- if .Info.Linkedin -}}
\small [\includegraphics[height=\logoheight]{./assets/logos/linkedin.png} \textcolor{black}{LinkedIn}](https://www.linkedin.com/in/edouard-jubert-9a348b58/)
{{- if or .Info.Malt -}} \space|\space {{- end -}}
{{- end -}}

{{- if .Info.Malt -}}
\small [\includegraphics[height=\logoheight]{./assets/logos/malt.png} \textcolor{black}{Malt}](https://www.malt.fr/profile/edouardjubert)
{{- end -}}

{{- end }}


# SKILLS

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
{{- range $key, $value := .Skills }}
\item \textbf{ {{ $value.Label }}: }
{{- range $index, $value := $value.Items }}{{ if $index }}, {{ end }}{{ $value }}{{- end }}
{{- end }}
\end{itemize}

\vspace{\defaultvspace}

# EXPERIENCES

{{- range .Experiences }}

## \textcolor{darkgray}{ {{- .Company -}} }

{{- range .Roles }}
#### {{ .Title }}

\textcolor{gray}{ {{ .Date }} }

{{ range .Items }}
\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{ {{ .Description }} }}
\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
{{- range .Tasks }}
\item {{ . }}
{{- end }}
\end{itemize}
{{- end }}

\textit{Technologies: } {{- range $index, $value := .Technologies }}{{ if $index }}, {{ end }}{{ $value }}{{ end }}

{{ end }}
{{- end }}


\vspace{\defaultvspace}

# EDUCATION

{{- range .Education }}
**{{ .School }}** - \textit{ {{ .Degree }} }

\textcolor{gray}{ {{ .Date }} }

{{ end }}


\vspace{\defaultvspace}

# LANGUES

{{- range .Languages }}
**{{ .Language }}** - {{ .Level }}
{{ end }}
