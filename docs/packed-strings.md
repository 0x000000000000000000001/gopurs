# Charge utile `String` packée dans le runtime

`TypeString` ne réserve plus un en-tête `*string` par boîte. La valeur porte
directement ses données :

| Champ | Contenu |
| --- | --- |
| `IntVal` | longueur en octets |
| `UnsafePtr` | données immuables de la chaîne (`unsafe.StringData`) |
| pointeur nul | chaîne vide, lue comme `""` |

L'écriture passe par `Str`, la lecture par `StrValue` :

```go
func StrValue(v Value) string {
	if v.UnsafePtr == nil {
		return ""
	}
	return unsafe.String((*byte)(v.UnsafePtr), int(v.IntVal))
}
```

La taille de `Value` est inchangée, comme pour les flottants déjà packés dans
`IntVal`.

## Contrat

- Les octets référencés doivent rester **immuables pendant toute la vie** de la
  valeur, exigence d'`unsafe.String`. Toutes les chaînes produites par `Str`
  proviennent de chaînes Go immuables ; aucune FFI ne construit de chaîne sur
  un tampon mutable sans copie.
- Un sous-`string` retient tout son support : c'est un compromis de rétention,
  pas un problème de correction.
- Une chaîne vide n'a pas de pointeur ; ne jamais déréférencer `UnsafePtr`
  sans passer par `StrValue`.
- Aucun code ne doit convertir `UnsafePtr` en `*string` : le helper est la
  seule lecture autorisée.

## Surfaces adaptées

- `runtime/runtime.go` : `Str`, `StrValue` et 13 lecteurs (`StrVal`, `AnyVal`,
  `RecordGet`, `Unbox`, `ValueToAny`, `ExtractVariant`, `copyReflectField`).
- `src/Gopurs/Runtime.go` et `Runtime.js` : régénérés par
  `node tools/embed-runtime.mjs` (ou `npm run build:runtime`).
- FFI : `gopurs-argonaut-core` (1 site), `gopurs-foreign` (3), 
  `gopurs-node-buffer` (9), tous remplacés par `gopurs_runtime.StrValue`.
- `tools/ffi-generics.test.mjs` : le garde-fou d'ordonnancement attend
  désormais **zéro allocation** par appel au lieu de deux boîtes de chaînes.

## Contrôles

- Sonde GC autonome : 4 M de valeurs packées, sous-chaînes comprises, relues
  correctement après plusieurs collectes, y compris sous `-race`/checkptr.
- `node --test tools/*.test.mjs` : 170 tests, 168 réussis, 0 échec, 2 ignorés.
- `bin/modtest` : `argonaut-core`, `foreign`, `foreign-object`, `node-buffer`,
  `prelude`, `strings-extra` verts. `gopurs-strings` échoue sur un littéral à
  surrogates isolés ; le même échec se reproduit avec le bundle baseline et
  correspond aux exclusions `StringEdgeCases`/`StringEscapes`.
- Oracles des diagnostics altbak : 17 cas JSON et 12 modules TAST identiques
  dans chaque processus mesuré.
- `bin/go/run -c` : 14 sorties de référence, total 13,138 ms.

## Effet mesuré

Campagnes appariées (5 processus par version, ordre alterné) : décodage
**−10,79 %** (JSON général) et **−8,21 %** (TAST) ; combinés **−4,61 %** et
**−4,36 %** ; allocations par corpus **−2,68 Mio** et **−23,13 Mio**. Le
parseur est inchangé. Détails, campagnes officielles et provenance :
[rapport altbak](../../../altbak.pub-gopurs/docs/benchmark-results/2026-09-22-packed-strings.md).
