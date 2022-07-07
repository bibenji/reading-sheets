# Les design patterns en Java
## Les 23 modèles de conception fondamentaux

Préface

Patterns d’interface

ADAPTER (17) fournit l’interface qu’un client attend en utilisant les services d’une
classe dont l’interface est différente.

FACADE (33) fournit une interface simplifiant l’emploi d’un sous-système.

COMPOSITE (47) permet aux clients de traiter de façon uniforme des objets indivi-
duels et des compositions d’objets.

BRIDGE (63) découple une classe qui s’appuie sur des opérations abstraites de
l’implémentation de ces opérations, permettant ainsi à la classe et à son implémen-
tation de varier indépendamment.

Patterns de responsabilité

SINGLETON (81) garantit qu’une classe ne possède qu’une seule instance, et fournit
un point d’accès global à celle-ci.

OBSERVER (87) définit une dépendance du type un-à-plusieurs (1,n) entre des objets
de manière à ce que lorsqu’un objet change d’état, tous les objets dépendants en
soient notifiés et soient actualisés afin de pouvoir réagir conformément.

MEDIATOR (103) définit un objet qui encapsule la façon dont un ensemble d’objets
interagissent. Cela promeut un couplage lâche, évitant aux objets d’avoir à se
référer explicitement les uns aux autres, et permet de varier leur interaction indé-
pendamment.

PROXY (117) contrôle l’accès à un objet en fournissant un intermédiaire pour cet
objet.

CHAIN OF RESPONSABILITY (137) évite de coupler l’émetteur d’une requête à son
récepteur en permettant à plus d’un objet d’y répondre.

FLYWEIGHT (145) utilise le partage pour supporter efficacement un grand nombre
d’objets à forte granularité.

Patterns de construction

BUILDER (159) déplace la logique de construction d’un objet en-dehors de la classe
à instancier, typiquement pour permettre une construction partielle ou pour simplifier
l’objet.

FACTORY METHOD (167) laisse un autre développeur définir l’interface permettant de
créer un objet, tout en gardant un contrôle sur le choix de la classe à instancier.

ABSTRACT FACTORY (175) permet la création de familles d’objets ayant un lien ou
interdépendants.

PROTOTYPE (187) fournit de nouveaux objets par la copie d’un exemple.

MEMENTO (193) permet le stockage et la restauration de l’état d’un objet.

Patterns d’opération

TEMPLATE METHOD (217) implémente un algorithme dans une méthode, laissant à
d’autres classes le soin de définir certaines étapes de l’algorithme.

STATE (229) distribue la logique dépendant de l’état d’un objet à travers plusieurs
classes qui représentent chacune un état différent.

STRATEGY (241) encapsule des approches, ou stratégies, alternatives dans des classes
distinctes qui implémentent chacune une opération commune.

COMMAND (251) encapsule une requête en tant qu’objet, de manière à pouvoir para-
métrer des clients au moyen de divers types de requêtes (de file d’attente, de temps
ou de journalisation) et de permettre à un client de préparer un contexte spécial dans
lequel émettre la requête.

INTERPRETER (261) permet de composer des objets exécutables d’après un ensemble
de règles de composition que vous définissez.

Patterns d’extension

DECORATOR (287) permet de composer dynamiquement le comportement d’un
objet.

ITERATOR (305) fournit un moyen d’accéder de façon séquentielle aux éléments
d’une collection.

VISITOR (325) permet de définir une nouvelle opération pour une hiérarchie sans
changer ses classes.

### I Patterns d'interface

P. 28

#### ADAPTER

L’objectif du pattern ADAPTER est de fournir l’interface qu’un client attend en
utilisant les services d’une classe dont l’interface est différente.

un adaptateur d’objet, c’est-à-dire un adaptateur qui utilise
la délégation plutôt que la dérivation de sous-classes

un adaptateur de classe étend une classe existante et implé-
mente une interface cible tandis qu’un adaptateur d’objet étend une classe cible et
délègue à une classe existante

#### FACADE

P. 50

#### COMPOSITE

#### BRIDGE

### Patterns de responsabilité

#### SINGLETON

#### OBSERVER

#### MEDIATOR

#### PROXY

#### CHAIN OF RESPONSABILITY

#### FLYWEIGHT

### Patterns de construction

#### BUILDER

#### FACTORY METHOD

#### ABSTRACT FACTORY

#### PROTOTYPE

#### MEMENTO

### Patterns d'opérations

#### TEMPLATE METHOD

#### STATE

#### STRATEGY

#### COMMAND

#### INTERPRETER

### Patterns d'extension

#### DECORATOR

#### ITERATOR

#### VISITOR