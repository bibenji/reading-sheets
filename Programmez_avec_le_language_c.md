cout = "ci aoute"
fourni par iostream
endl : créer un retour à la ligne

P. 48

```
#include <iostream>
#include <string>
```

int maVar(3);
int& maRef(maVar);

cin = "si-inne"

cout << "Blabla" << endl;
cin >> nomVar;

P. 66

int const nombre(2);
unsigned int const autreNombre(2);

`#include <cmath>`

P. 88

```asp
type nomFonction(arguments)
{
	//Instructions effectuées par la fonction
}

int ajouteDeux(int nombreRecu)
{
	int valeur(nombreRecu + 2);
	return valeur;
}
```

fonctions avec même nom, arguments différents s'appelle la surcharge

void direBonjour() { }

les valeurs des variables transmises aux fonctions sont copiées dans de nouvelles cases mémoires

Passage par valeur et passage par référence

```asp
int ajouteDeux(int& a) //Notez le petit & !!!
{
	a+=2;
	return a;
}
```

Avancé : Le passage par référence constante
= évite la copie et empêche modification

```asp
void f1(string const& texte); // Pas de copie et pas de modification possible
{
}
```

.cpp (code source)
.h (header)

```asp
#include "math.h"

int ajouteDeux(int nombreRecu)
{
	int valeur(nombreRecu + 2);
	return valeur;
}
```

```asp
#ifndef MATH_H_INCLUDED
#define MATH_H_INCLUDED

int ajouteDeux(int nombreRecu);

#endif // MATH_H_INCLUDED
```

```asp
#ifndef MATH_H_INCLUDED
#define MATH_H_INCLUDED

#include <string>

void afficherMessage(std::string message);

#endif // MATH_H_INCLUDED
```

```asp
#include <iostream>
#include "math.h"
using namespace std;
int main()
{

}
```

système doxygen pour description des fonctions :

/**
* \brief Fonction qui ajoute 2 au nombre reçu en argument
* \param nombreRecu Le nombre auquel la fonction ajoute 2
* \return nombreRecu + 2
*/
int ajouteDeux(int nombreRecu);

int nombreDeSecondes(int heures, int minutes = 0, int secondes = 0);

```asp
#include <iostream>
using namespace std;

// Prototype avec les valeurs par défaut
int nombreDeSecondes(int heures, int minutes = 0, int secondes = 0);

// Main
int main()
{
	cout << nombreDeSecondes(1, 10, 25) << endl;
	return 0;
}
// Définition de la fonction, SANS les valeurs par défaut
int nombreDeSecondes(int heures, int minutes, int secondes)
{
	int total = 0;
	total = heures * 60 * 60;
	total += minutes * 60;
	total += secondes;
	return total;
}
```

il ne faut spécifier les valeurs par défaut que dans le fichier d'en-tête .h

# Les tableaux

P. 114
