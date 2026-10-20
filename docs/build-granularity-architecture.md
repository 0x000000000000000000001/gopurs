# Granularité Go et architecture incrémentale — exploration du 20 octobre 2026

**Statut : exploration privée terminée, aucun changement de production.**
Les preuves complètes sont dans
`htdocs/.gopurs-evidence/a-20261020-granularity/report.md` et
`a-20261020-partition/`.

## Pourquoi ce chantier

A est livré : froid installé **76,163 s**, hit **0,914 s**, feuille backend
**45,421 s**. Le `go build` après feuille de G06 valait **48,490 s** pour un
package généré de 39,8 Mio. Les optimisations backend seules ne peuvent pas
atteindre 30 s de chaîne complète.

## Résultats mesurés

### Densité Go

Mesures du vrai `main`, cache de base chaud, `-p 1`, `GOMAXPROCS=2` :

| Density | Packages | Plus gros package | Go feuille |
| ---: | ---: | ---: | ---: |
| 512 | 123 | 16 Mo | 20,2 s |
| 768 | 77 | 20 Mo | 20,3 s |
| **1024** | **68** | **20 Mo** | **9,6 s** |
| 1536 | 33 | 29 Mo | 22,9 s |
| 2048 | 39 | 27 Mo | 27,3 s |
| 4096 | 17 | 39 Mo | 29,4 s |
| 8192 | 6 | 39 Mo | 30,2 s |

À density 1024 et `-p 14` / `GOMAXPROCS=14`, le cas Ping tombe à **2,67 s**.
Le froid Go est **46,7 s** (`-p 14`), contre ~84 s en `-p 1` à 8192.

### Limite critique : le placement et les SCC

La densité 1024 n’est pas générale :

| Édition | Backend density 1024 | Go `-p 14` |
| --- | ---: | ---: |
| feuille Ping | 43,7 s | **2,67 s** |
| `ffi` | 49,4 s | **32,9 s** |
| `shared` | 77,4 s | **31,2 s** |

À density 128, le package FFI modifié fait 0 Mo mais a 125 dépendances ; le
`main` recompile quand même **33,9 s**. À density 64, le froid est 59,9 s et la
feuille 17 s. Density 1 produit 69 186 packages et est inutilisable.

### Isolation dynamique

Un prototype privé déplace les déclarations d’un module généré dans un package
dédié, recalcule les dépendances et rend les packages :

| Module isolé | Go base | Go feuille | Verdict |
| --- | ---: | ---: | --- |
| `Inter_Cli_Ping_Main` | 88,3 s | **1,16 s** | succès |
| `Util_Type_String_String` | — | — | cycle d’imports Go |

Le cas FFI crée `p_02c4 → p_iso → p_02c4`. Isoler ce module exige de déplacer
tout le SCC ou de casser le cycle généré. L’isolation naïve n’est donc pas une
solution universelle.

### Backend après granularité

Feuille density 1024, 8 jobs : **43,673 s** :

| Phase | Temps |
| --- | ---: |
| Lookup + gel | 2,74 s |
| Préparation + monomorphisation | 9,75 s |
| Clés des modules | 5,34 s |
| PBO producer | 9,73 s |
| Go packages | 10,06 s |
| Publication/reste | ~6,05 s |

Avec 14 jobs backend, le leaf monte à 47,76 s : le PBO passe à 13,2 s. Garder
8 jobs backend.

### Scénarios complets (frontend/tidy ~2 s)

| Scénario | Backend | Go feuille | Total |
| --- | ---: | ---: | ---: |
| A actuel, 6 packages | 45,4 s | ~30 s | ~77 s |
| density 1024, Ping | 43,7 s | 2,67 s | ~48 s |
| isolation Ping | 43,7 s | 1,16 s | ~47 s |
| density 1024, FFI | 49,4 s | 32,9 s | ~84 s |
| shared | 77,4 s | 31,2 s | ~111 s |
| froid density 1024 | 77,1 s | 46,7 s | ~126 s |

Même le meilleur cas n’atteint pas 30 s. Le backend est devenu le mur principal.
Le froid 60 s reste hors de portée avec le Go froid à 46,7 s et le backend
froid à 77 s.

## Décision d’architecture

Deux classes de modules existent :

- **séparables** : Go feuille ~1–3 s possible ;
- **dans un SCC/cycle** : plancher ~33 s avec le modèle actuel.

Aucune densité statique ni isolation naïve ne donne un build court universel.
La suite doit donc être un chantier d’architecture, pas un nouveau réglage.

### C1 — découpage Go SCC-aware

Objectif : rendre l’isolation valide pour tous les modules.

- plan de packages déterministe et cache par module/bundle ;
- extraction du module édité, ou du SCC minimal, sans cycle ;
- gestion explicite de l’ordre d’initialisation, des symboles privés, des
  helpers partagés et des exports ;
- fallback complet si le SCC ne peut pas être divisé ;
- conservation des sorties et de l’ABI Go.

Critère : Go feuille ≤5 s sur un module séparable et ≤15 s sur un module d’un
SCC large, sans régression froide >5 %. Si le SCC FFI reste >25 s, arrêter C1.

### C2 — backend résident/incrémental

Objectif : supprimer les scans globaux de préparation, clés, PBO et
planification à chaque édition.

- processus résident ou état sérialisé de session ;
- seuls le module modifié et sa fermeture prouvée sont retraités ;
- cache des modules préparés, des implémentations/directives et du plan de
  packages ;
- validation/lookup/publication incrémentaux ;
- fallback complet en cas de doute.

Critère : backend feuille ≤20–25 s, puis ≤15 s avec B1/B3 ; régression froide
≤5 %. Sans preuve de dépendances, ne pas livrer.

### C3 — froid

Conditionnel. Le froid actuel (backend 77 s + Go 46,7 s) ne peut pas viser 60 s
par un cache chaud. À n’ouvrir qu’après C1/C2 et avec un budget séparé.

## Rôle de B1/B2/B3

- **B1** : le `planHash` est identique base/feuille à toutes les densités
  testées ; le package modifié fait 4 Mo à 1024. C’est un composant de C2, pas
  une solution seule.
- **B2** : PBO 9,9 → 7,9 s, sans gain total dans l’ancien chemin ; peut devenir
  utile dans C2.
- **B3** : lookup 0,914 → 0,655 s, partiel ; clés préparées (~5,3 s) encore
  entières et dépendantes des preuves P03/P04.

## Protocole et limites

- Trois paires alternées, sorties exactes, pas de campagne lourde concurrente.
- Les chemins de packages changent avec la densité : l’oracle byte-exact par
  chemin n’est pas applicable ; il faut un oracle d’exécution.
- Density 1024 non validée sur les quatre hôtes ni en stress multi-éditions.
- `-p 14` consomme ~7,3 Gio au froid.
- Aucun commit/push de l’assistant.


## C1 — verdict du prototype SCC/closure

Le prototype privé a mesuré les SCC et testé deux stratégies :

| Module | SCC | Déclarations déplacées | Résultat |
| --- | ---: | ---: | --- |
| `Inter_Cli_Ping_Main` | 1 composant | 3 | succès, Go feuille 1,16 s |
| `Util_Type_String_String` | 57 composants | 4 447 | cycle d’imports Go |
| FFI, fermeture de dépendances | — | 180 | cycle d’imports Go |

La fermeture de dépendances ne suffit pas : déplacer un sous-ensemble d’un
package peut créer un cycle **au niveau package**, même si le graphe de
déclarations est acyclique localement. Le package FFI d’origine fait 11 Mo et
reste compilé ; le package extrait ne résout pas la fermeture de l’entrée.

**Verdict : C1 tel que défini est rejeté.** L’isolation SCC/closure ne donne pas
un découpage universel. Il faudrait un refactor de génération pour casser les
cycles au niveau package (helpers partagés, exports, init), ou accepter que les
modules d’un SCC large gardent un plancher Go de ~33 s.

C2 (backend résident) peut encore réduire le backend, mais ne supprime pas ce
plancher Go. C1b (refactor codegen SCC) est un nouveau projet, pas une suite
mécanique de C1.


## C1b — première tentative de casse de cycle

Le cycle FFI a été tracé : le module FFI utilise
`Util.Html.Clean.Clean` depuis `p_02c4`, tandis que `p_02c4` importe le
symbole FFI. Une extraction locale de `Util.Html.Clean.Clean` **et** du module
FFI a été testée.

Résultat : **cycle persistant** via d’autres packages
(`p_03f…`, `p_7525…`, `p_9e27…`, etc.). Le graphe n’a pas une seule arête à
casser ; c’est un réseau de dépendances mutuelles au niveau package.

C1b exige donc un vrai refactor de génération :

- graphe SCC package complet et classification des arêtes ;
- extraction des helpers/types partagés vers des packages bas ;
- inversion de dépendance (interfaces, callbacks, registres) quand un appel
  direct est nécessaire ;
- ordre d’initialisation explicite ;
- validation d’ABI et de sorties.

Aucun prototype ne démontre encore que ce refactor reste borné. C1b est
**bloqué en conception**, pas un simple patch de planificateur.


## C1b — mesure de fermeture globale

| Module | Déclarations dans la fermeture | Unités | Taille | Total du programme |
| --- | ---: | ---: | ---: | ---: |
| Ping | 2 278 | 186 | 7,4 Mo | 185 667 déclarations / 123 Mo |
| FFI | 2 510 | 205 | **10,9 Mo** | idem |

La fermeture FFI fait presque la taille du package FFI d’origine (11 Mo) :
la déplacer ne réduit pas le coût de compilation. Après extraction de
`Util.Html.Clean.Clean`, `p_iso` importe encore plusieurs packages qui
participent au cycle. Il n’y a pas de petite coupe locale.

**Verdict C1b :** un simple SCC/closure/helper ne suffit pas. Il faut une
inversion de dépendance à l’échelle du graphe (interfaces/callbacks/registres),
avec ordre d’init et ABI, ou accepter le plancher Go de ~33 s pour les modules
d’un SCC large.


## C1b — graphe de cycles après extraction

Après isolation de FFI et de `Util.Html.Clean.Clean`, le package isolé `p_iso`
importe encore **12 packages sur cycle** et **51 packages** peuvent l’atteindre.
Une arête `p_02c4 → p_iso` représente à elle seule 51 références de symboles
(types, constructeurs, fonctions). Un pont sur une seule arête ne suffit pas.

Pour rendre l’isolation acyclique, il faudrait :

- traiter les 12 imports directs et leurs fermetures ;
- remplacer types/constructeurs/fonctions par des ponts ou déplacer les
  helpers partagés ;
- préserver l’ordre d’init et l’ABI ;
- valider 51 packages dépendants.

C1b complet n’est donc pas un prototype d’une heure ; c’est une refonte du
graphe de génération. Le plancher actuel reste ~33 s pour FFI.


## C1b — partition topologique et noinline

Dernière variante testée : SCC ordonnés topologiquement, packages de taille
bornée, noms stables (`planHash` identique base/FFI). Le package modifié FFI
fait **0,8 Mo** et la feuille tombe à **46,7 s** (contre 32,9 s en densité 1024).
Sans inlining Go (`-gcflags=all=-l`), elle reste à **25,6 s**, au-dessus de la
cible ≤15 s. Les dépendants recompilent encore.

C1b est rejeté ; une refonte ABI/export plus profonde serait nécessaire.
