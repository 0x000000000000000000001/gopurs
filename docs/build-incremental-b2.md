# B2 — descripteur PBO sans décodage du corps

**20 octobre 2026 — candidat rejeté.** Objectif : ≥5 s sur la feuille. Le PBO
gagne ~2 s, mais le backend total ne montre pas de gain reproductible ; le
seuil n’est pas atteint.

## Prototype

- Un module dont les clés PBO complètes correspondent peut être réutilisé sans
  décoder la partie `pbo/backend-v2` : seules les implémentations, les directives
  et les workers sont lues.
- Le backend minimal porte le nom, les imports, la présence FFI et les maps
  d’implémentations/directives ; l’émetteur recopie les sorties précédentes.
- Le memo porte le digest FFI de l’émetteur. Si ce digest diffère, le module
  repasse par l’émission complète ; le chemin `lookupAll` est couvert aussi.
- Les modules FFI ou dont le digest FFI change ne prennent pas le raccourci.

## Mesures

| Passage | Avant | B2 | Verdict |
| --- | ---: | ---: | --- |
| PBO feuille | 9,876 s | **7,882 s** | ~2 s, sous le seuil |
| Backend feuille | 45,421 s | **45,805 s** | pas de gain total |
| Transition `ffi` | — | 40,400 s | exacte, 1 module émis / 2 703 restitués |

Les 27 sorties sont exactes sur la feuille et sur `ffi` après correction. Les
premières tentatives `ffi` ont révélé un défaut réel : la clé PBO ne couvre pas
le digest FFI de l’émetteur. Le memo FFI et le contrôle dans `emit` corrigent
ce cas, mais le coût/complexité ne se justifie pas sans gain total.

## Décision

**Rejeté et retiré du chemin installé.** Les preuves et les journaux restent
dans `htdocs/.gopurs-evidence/a-20261020-b2*` ; les binaires A/B3 précédents
sont restaurés. B3 lookup reste partiel ; B1 est le prochain candidat lourd.
