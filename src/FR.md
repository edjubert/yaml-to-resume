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

**Backend Developer freelance | Golang, Typescript**

\textcolor{gray}{Paris, France}

\vspace{\defaultvspace}

\small \href{mailto:edouard.jubert@proton.me}{\textcolor{black}{\includegraphics[height=\logoheight]{./assets/email.png} edouard.jubert@proton.me}} | [\includegraphics[height=\logoheight]{./assets/github.png} \textcolor{black}{Github}](https://github.com/edjubert) | [\includegraphics[height=\logoheight]{./assets/linkedin.png} \textcolor{black}{LinkedIn}](https://www.linkedin.com/in/edouard-jubert-9a348b58/) | [\includegraphics[height=\logoheight]{./assets/malt.png} \textcolor{black}{Malt}](https://www.malt.fr/profile/edouardjubert)

# SKILLS

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item \textbf{Languages:} Golang, TypeScript
\item \textbf{Bases de Données:} PostgreSQL, Redis, ClickHouse, Elasticsearch
\item \textbf{Technologies:} Kafka, gRPC
\item \textbf{DevOps:} Docker, Kubernetes, Bazel, CI/CD, AWS, GCP
\end{itemize}

\vspace{\defaultvspace}

# EXPERIENCES

## \textcolor{darkgray}{Axens}

#### Développeur Backend Golang - Freelance

\textcolor{gray}{avril 2024 - février 2025}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Ajout d'un backend Go à une datasource Grafana.
\item Migration de l'API TypeScript vers le backend Go.
\item Création d'un environement de développement sous Docker.
\end{itemize}

\textit{Technologies:} Go, Typescript, Grafana, Docker, MySQL, InfluxDB

\vspace{\defaultvspace}

## \textcolor{darkgray}{Polyfire}

#### Développeur Backend Golang - Freelance

\textcolor{gray}{décembre 2024 - février 2025}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Création d'un outil de scrapping des amendements de l'assemblée nationale
\item Création d'un système de reconnaissance faciale des députés de l'assemblée nationale en architecture orientée événements et microservices
\end{itemize}

\textit{Technologies:} Go, Typescript, GCP, Terraform, MySQL

\vspace{\defaultvspace}

## \textcolor{darkgray}{Leboncoin.fr}

#### Développeur Backend Golang - Freelance

\textcolor{gray}{avril 2023 - avril 2024}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Intégration à la Feature Team Care pour l'amélioration des parcours d'incident utilisateur.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Développement API back-end en micro services
\item Conception et mise en place d'une API Go en event driven
\item Création d'outils internes pour améliorer les phases de test de l'équipe
\item Création d'outils internes pour corriger des bugs de mise en production
\end{itemize}

\textit{Technologies:} Go, Kafka, Datadog, PostgreSQL

\vspace{\defaultvspace}

## \textcolor{darkgray}{Salesramp}

#### Développeur Backend Golang - Freelance

\textcolor{gray}{juillet 2022 - avril 2023}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Travail en collaboration avec le CTO pour la création du produit.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Architecture et mise en place de la base de données ClickHouse.
\item Création d'un serveur gRPC en Go pour communiquer avec la base de
données Clickhouse.
\item Refonte et optimisation de l'API.
\item Migrations de l'API en Typescript.
\item Création de fonctions de communication gRPC entre le serveur Typescript et
Go
\item Mise en place de bonnes pratiques.
\end{itemize}

\textit{Technologies:} Go, PostgreSQL, ClickHouse, NodeJS, Typescript, Parse Server

## \textcolor{darkgray}{Yuzu App}

#### Développeur Full Stack - Freelance

\textcolor{gray}{mai 2022 - juillet 2022}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Travail en collaboration avec le fondateur pour la mise en production de l'application.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Nettoyage et sécurisation de la code base
\item Optimisation de l'application
\item Développement de features
\item Correction de bugs
\item Mise en production sur Play Store et AppStore
\item Mise en place d'un environnement staging
\end{itemize}

\textit{Technologies:} React Native, Node.JS, Typescript, Firebase

\vspace{\defaultvspace}

## \textcolor{darkgray}{Fastory.io}

#### Développeur Lead Full Stack

\textcolor{gray}{juin 2020 - avril 2022}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Lead développeur d'une équipe de 5 à 7 développeurs.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Organisation de rituels
\item Priorisation des projets
\item Suivi et accompagnement des juniors
\item Écriture de la partie technique du dossier pour le CIR
\item Recrutement
\end{itemize}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Chantier d'amélioration structurelle du produit pour faire face a l'augmentation du trafique.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Migration Heroku vers AWS avec ECS et Fargate
\item Mise en place de l'autoscaling
\item Migration PostgreSQL vers Elasticsearch
\end{itemize}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Chantier d'amélioration de la codebase.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Migration de Javascript vers Typescipt
\item Mise a jour du router React v1 vers v4
\item Adoption du design system "Atlaskit"
\item Création de composants React et NextJS
\end{itemize}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Ajouts de features.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Création de multiples add-ons au sein de la plateforme Fastory
\item Mise en place d'un système de connexion via SAP/CDC en collaboration avec les JO2024
\end{itemize}

\textit{Technologies:} TypeScript, React.js, Next.js, Hapi.js, AWS, Elasticsearch, PostgreSQL

\vspace{\defaultvspace}

#### Développeur Full Stack

\textcolor{gray}{août 2019 - mai 2020}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Intégration à une équipe de 4 développeurs.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Migration d'une partie du serveur Javascript vers un serveur Next en
Typescript
\item Mise en place de fonctionnalités
\item Correction de dette technique
\end{itemize}

\textit{Technologies:} TypeScript, JavaScript, Heroku, React.js, PostgreSQL, Hapi.js

\vspace{\defaultvspace}

## \textcolor{darkgray}{HAPPLYFACE}

#### Développeur - Projet "Shindy Me"

\textcolor{gray}{septembre 2016 - juillet 2018}

\vspace{\defaultvspace}
\textbf{\textcolor{darkgray}{Travail en collaboration avec le fondateur pour la création de l'application Shindy Me.}}

\begin{itemize}
\setlength{\leftskip}{\defaultleftskip}
\item Développement de l'application " Shindy Me " iOS/Android en Ionic3
\item UXDesign de l'application
\item Mise en place de la base de données (Firebase Realtime Database)
\item Serveur (Firebase Functions, NodeJS)
\item Maintenance, amélioration et implémentation
\item Mise en place d'une migration vers ReactNative et Firebase Cloud Firestore.
\end{itemize}

\textit{Technologies:} Ionic, Angular 2, Firebase, React Native

\vspace{\defaultvspace}

# EDUCATION

**42** - \textit{Architecte en technologie du numérique}

\textcolor{gray}{2018 - 2024}

\vspace{\defaultvspace}
**Université Panthéon-Assas (Paris II)** - \textit{Economie Gestion}

\textcolor{gray}{2010 - 2014}

\vspace{\defaultvspace}

# LANGUES

**Anglais** - Capacité professionnelle complète

**Français** - Langue maternelle
