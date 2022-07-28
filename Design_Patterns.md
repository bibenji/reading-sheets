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

L’objectif du pattern FACADE est de fournir une interface simplifiant l’emploi
d’un sous-système.

Les équations paramétriques
simplifient la modélisation
de courbes lorsque y n’est pas
une fonction monovaluée
de x. [...]

Une façade est une classe configurable et réutilisable,
avec une interface de plus haut niveau qui simplifie l’emploi du sous-système.

#### COMPOSITE

P. 62

Un COMPOSITE est un groupe d’objets contenant aussi bien des éléments individuels
que des éléments contenant d’autres objets.

groupes et feuilles

L’objectif du pattern COMPOSITE est de permettre aux clients de traiter de
façon uniforme des objets individuels et des compositions d’objets.

[...]

Vous pouvez
implémenter une méthode isTree() sur la classe abstraite MachineComponent afin
de déléguer l’appel à une méthode isTree() conservant un ensemble des nœuds
parcourus.
(pour éviter de compter plusieurs fois le même noeud)

[...]

En procédant avec soin, vous pouvez garantir qu’un modèle objet reste un arbre en
refusant tout changement qui ferait retourner false par isTree(). D’un autre côté,
vous pouvez décider d’autoriser l’existence de composites qui ne sont pas des arbres,
surtout lorsque le domaine de problèmes que vous modélisez contient des cycles.

en gros : on passe un Set (en java) qui contient les références aux composants déjà visités

[...] PARTIE A REVOIR

#### BRIDGE

P. 76

L’objectif du pattern BRIDGE est de découpler une abstraction de l’implémenta-
tion de ses opérations abstraites, permettant ainsi à l’abstraction et à son
implémentation de varier indépendamment.

Les drivers sont des abstractions. Le résultat de l’exécution de l’application dépend
du driver en place. Chaque driver est une instance du pattern ADAPTER, fournissant
l’interface qu’un client attend en utilisant les services d’une classe comportant une
interface différente.

Une conception globale qui utilise des drivers est une instance de BRIDGE.

Vous pouvez
inclure des méthodes que certains drivers ne supporteront pas, ou exclure des
méthodes pour limiter ce que les abstractions pourront faire avec un driver ou bien
les forcer à inclure du code pour un cas particulier.

(pour garder une fonctionnalité de certaines drivers, les autres l'ignoreront)

Un exemple banal d’application utilisant des drivers est l’accès à une base de
données.

### Patterns de responsabilité

Le code compile sans problème. L’accès est défini au niveau classe et non au niveau
objet. Aussi un objet Firework peut-il accéder aux variables et méthodes privées
d’un autre objet Firework, par exemple.

le développement OO promeut l’encapsulation, l’idée
qu’un objet travaille sur ses propres données.

patterns qui sont exceptions à la règle de responsabilité répartie

#### SINGLETON

P. 94

L’objectif du pattern SINGLETON est de garantir qu’une classe ne possède
qu’une seule instance et de fournir un point d’accès global à celle-ci.

private static Factory factory = new Factory();

Cette classe pourrait rendre son unique instance disponible par l’intermédiaire
d’une méthode getFactory() publique et statique.

initialisation tardive, dite "paresseuse", ou lazy-initialization

public static Factory getFactory() {
    if (factory == null)
        factory = new Factory();
        // ...

    return factory;
}

Voir Concurrent Programming in Java™

#### OBSERVER

P. 100

L’objectif du pattern OBSERVER est de définir une dépendance du type un-à-
plusieurs (1,n) entre des objets de manière que, lorsqu’un objet change d’état,
tous les objets dépendants en soient notifiés afin de pouvoir réagir conformément.

un seul objet,
l’application, sait quels objets actualiser et se charge d’émettre les interrogations
appropriées => pas observer
laisser chaque objet s’enregistrer lui-même de manière indi-
viduelle => observer

Modèle-Vue-Contrôleur

Cette conception permet de n’effectuer qu’une seule fois le travail de traduction de
la valeur du curseur en valeur de temps crête. L’application actualise un seul objet
Tpeak, et tous les objets de GUI intéressés par un changement peuvent interroger
l’objet pour en connaître la nouvelle valeur.

Exercice 9.4

P. 108

Observer, PropertyChangeSupport

Pour une
grande GUI, envisagez la possibilité de passer à une conception MVC, en permet-
tant à chaque objet intéressé de gérer son besoin d’être notifié au lieu d’introduire
un objet central médiateur.

#### MEDIATOR

P. 116

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