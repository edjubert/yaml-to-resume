---
documentclass: extarticle
title: Edouard Jubert
geometry: margin=1cm
mainfont: DejaVu Sans
monofont: DejaVu Sans Mono
fontsize: 9pt
header-includes: |
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
\adjustimage{width=\imagewidth,clip,raise=20pt}{./assets/me.png}
\end{textblock*}

\vspace{\defaultvspace}

**Backend Developer Freelance | Golang, TypeScript**

\textcolor{gray}{Paris, France}

\vspace{\defaultvspace}

\small \href{mailto:edouard.jubert@proton.me}{\textcolor{black}{\includegraphics[height=\logoheight]{./assets/email.png} edouard.jubert@proton.me}} | [\includegraphics[height=\logoheight]{./assets/github.png} \textcolor{black}{Github}](https://github.com/edjubert) | [\includegraphics[height=\logoheight]{./assets/linkedin.png} \textcolor{black}{LinkedIn}](https://www.linkedin.com/in/edouard-jubert-9a348b58/) | [\includegraphics[height=\logoheight]{./assets/malt.png} \textcolor{black}{Malt}](https://www.malt.fr/profile/edouardjubert)

# SKILLS

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item \textbf{Languages:} Golang, TypeScript
\item \textbf{Databases:} PostgreSQL, Redis, ClickHouse, Elasticsearch
\item \textbf{Technologies:} Kafka, gRPC
\item \textbf{DevOps:} Docker, Kubernetes, Bazel, CI/CD, AWS, GCP
\end{itemize}

\vspace{\defaultvspace}

# EXPERIENCE

## \textcolor{darkgray}{Axens}

#### Backend Developer Golang - Freelance

\textcolor{gray}{April 2024 - February 2025}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Added a Go backend to a Grafana datasource
\item Migrated TypeScript API to Go backend
\item Created a Docker development environment
\end{itemize}

\textit{Technologies:} Go, Typescript, Grafana, Docker, MySQL, InfluxDB

\vspace{\defaultvspace}

## \textcolor{darkgray}{Polyfire}

#### Backend Developer Golang - Freelance

\textcolor{gray}{December 2024 - February 2025}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Created a scraping tool for National Assembly amendments
\item Created a facial recognition system for National Assembly deputies using event-driven and microservices architecture
\end{itemize}

\textit{Technologies:} Go, Typescript, GCP, Terraform, MySQL

\vspace{\defaultvspace}

## \textcolor{darkgray}{Leboncoin.fr}

#### Backend Developer Golang - Freelance

\textcolor{gray}{April 2023 - April 2024}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Integration with the Feature Team Care for improving user incident workflows.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Developed backend API in microservices
\item Designed and implemented a Go API in event-driven architecture
\item Created internal tools to improve team testing phases
\item Created internal tools to fix production deployment bugs
\end{itemize}

\textit{Technologies:} Go, Kafka, Datadog, PostgreSQL

\vspace{\defaultvspace}

## \textcolor{darkgray}{Salesramp}

#### Backend Developer Golang - Freelance

\textcolor{gray}{July 2022 - April 2023}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Collaborated with the CTO for product creation.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Architected and implemented ClickHouse database
\item Created a Go gRPC server for ClickHouse database communication
\item Refactored and optimized the API
\item Migrated API to TypeScript
\item Created gRPC communication functions between TypeScript and Go servers
\item Implemented best practices
\end{itemize}

\textit{Technologies:} Go, PostgreSQL, ClickHouse, NodeJS, Typescript, Parse Server

\vspace{\defaultvspace}

## \textcolor{darkgray}{Yuzu App}

#### Full Stack Developer - Freelance

\textcolor{gray}{May 2022 - July 2022}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Collaborated with the founder for application production deployment.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Cleaned and secured codebase
\item Optimized application
\item Developed features
\item Fixed bugs
\item Deployed to Play Store and AppStore
\item Set up staging environment
\end{itemize}

\textit{Technologies:} React Native, Node.JS, Typescript, Firebase

\vspace{\defaultvspace}

## \textcolor{darkgray}{Fastory.io}

#### Lead Full Stack Developer

\textcolor{gray}{June 2020 - April 2022}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Lead developer of a 5-7 developer team.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Organized rituals
\item Prioritized projects
\item Mentored junior developers
\item Wrote technical part of CIR documentation
\item Recruitment
\end{itemize}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Structural improvement project to handle increased traffic.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Migrated from Heroku to AWS with ECS and Fargate
\item Implemented autoscaling
\item Migrated from PostgreSQL to Elasticsearch
\end{itemize}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Codebase improvement project.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Migrated from Javascript to Typescript
\item Updated React router from v1 to v4
\item Adopted "Atlaskit" design system
\item Created React and NextJS components
\end{itemize}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Feature additions.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Created multiple add-ons within Fastory platform
\item Implemented SAP/CDC connection system in collaboration with Paris 2024 Olympics
\end{itemize}

\textit{Technologies:} TypeScript, React.js, Next.js, Hapi.js, AWS, Elasticsearch, PostgreSQL

\vspace{\defaultvspace}

#### Full Stack Developer

\textcolor{gray}{August 2019 - May 2020}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Integration with a 4-developer team.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Migrated part of Javascript server to Next server in Typescript
\item Implemented features
\item Fixed technical debt
\end{itemize}

\textit{Technologies:} TypeScript, JavaScript, Heroku, React.js, PostgreSQL, Hapi.js

\vspace{\defaultvspace}

## \textcolor{darkgray}{HAPPLYFACE}

#### Developer - "Shindy Me" Project

\textcolor{gray}{September 2016 - July 2018}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Collaborated with the founder to create Shindy Me application.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Developed "Shindy Me" iOS/Android app in Ionic3
\item UX Design of the application
\item Set up database (Firebase Realtime Database)
\item Server development (Firebase Functions, NodeJS)
\item Maintenance, improvement and implementation
\item Set up migration to ReactNative and Firebase Cloud Firestore
\end{itemize}

\textit{Technologies:} Ionic, Angular 2, Firebase, React Native

\vspace{\defaultvspace}

# EDUCATION

**42** - \textit{Digital Technology Architect}

\textcolor{gray}{2018 - 2024}

\vspace{\defaultvspace}
**Panthéon-Assas University (Paris II)** - \textit{Economics and Management}

\textcolor{gray}{2010 - 2014}

\vspace{\defaultvspace}

# LANGUAGES

**English** - Full professional proficiency

**French** - Native
