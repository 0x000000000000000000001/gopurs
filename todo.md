# Plan de publication officielle de gopurs

## 1. Figer le périmètre et le contrat public

- [ ] Choisir la version de départ et le schéma de versionnement.
- [ ] Lister les surfaces stables : CLI, backend Go, formats intermédiaires, bibliothèques `gopurs-*`.
- [ ] Lister explicitement ce qui reste expérimental.
- [ ] Fixer les plateformes supportées : OS, architectures, versions de Go et de PureScript.
- [ ] Auditer les licences du compilateur, du fork PureScript, de PBO et des dépendances.
- [ ] Ajouter les attributions et notices tierces nécessaires.
- [ ] Définir la politique de support et de dépréciation.

## 2. Garantir l'installation depuis zéro

- [ ] Tester un clone propre sur une machine sans cache ni binaire préexistant.
- [ ] Supprimer toute dépendance à des chemins locaux ou privés.
- [ ] Épingler les versions d'outillage et les dépendances de build.
- [ ] Documenter la chaîne de bootstrap : binaire JS, binaire Go précédent, artefacts de référence.
- [ ] Publier les checksums des binaires de bootstrap.
- [ ] Vérifier la reproductibilité des sorties, ou documenter précisément les différences.
- [ ] Mettre en place une CI Linux et macOS sur les architectures supportées.
- [ ] Vérifier que la CI part d'un environnement vierge, sans cache préchauffé obligatoire.

## 3. Nettoyer le dépôt

- [ ] Retirer les artefacts générés, caches, logs, binaires et sauvegardes du suivi Git.
- [ ] Retirer les preuves de benchmark et fichiers d'archive internes des surfaces publiées.
- [ ] Rechercher et supprimer les secrets, jetons, URLs privées et chemins personnels.
- [ ] Enlever les mentions `WIP` des surfaces officielles ou les convertir en limitations documentées.
- [ ] Supprimer le code mort et les options expérimentales visibles.
- [ ] Harmoniser le nommage `gopurs` et les intitulés des paquets.
- [ ] Vérifier que les fichiers publiés dans les paquets sont limités au nécessaire.
- [ ] Éviter tout refactor massif non nécessaire avant la première release.

## 4. Finaliser la documentation

- [ ] Écrire un quickstart testé de bout en bout.
- [ ] Documenter l'installation, les prérequis et la référence CLI.
- [ ] Documenter l'architecture, les formats et le fonctionnement du backend Go.
- [ ] Documenter les limitations connues et les différences avec le compilateur officiel.
- [ ] Documenter la migration et la compatibilité entre versions.
- [ ] Ajouter un guide de dépannage et une FAQ.
- [ ] Ajouter `CONTRIBUTING`, `SECURITY`, `CHANGELOG` et le code de conduite.
- [ ] Reprendre les benchmarks avec protocole, machine, commandes et scripts reproductibles.
- [ ] Vérifier qu'aucune promesse de performance ou de compatibilité n'est ambiguë.

## 5. Passer les portes de qualité

- [ ] Exécuter toute la suite de tests dans un environnement propre.
- [ ] Vérifier la conformité avec le compilateur officiel sur un corpus représentatif.
- [ ] Vérifier les golden tests, la parité FFI et les formats générés.
- [ ] Tester les erreurs, annulations, limites mémoire et cas de concurrence.
- [ ] Vérifier les performances à froid et à chaud avec le protocole publié.
- [ ] Faire tourner une release candidate et corriger les régressions bloquantes.
- [ ] Exiger une CI verte sur le tag de release.

## 6. Préparer les artefacts et la publication

- [ ] Produire les binaires par plateforme avec les flags de release documentés.
- [ ] Générer les checksums, la provenance et un SBOM.
- [ ] Signer le tag et les artefacts.
- [ ] Rédiger les notes de version et le changelog.
- [ ] Tester l'installation uniquement depuis les artefacts publiés.
- [ ] Vérifier les paquets `gopurs-*` : versions alignées, bornes de dépendances, allowlist de fichiers.
- [ ] Faire un `npm publish --dry-run` et vérifier la provenance si publication sur npm.
- [ ] Préparer la procédure de rollback et conserver les artefacts précédents.
- [ ] Effectuer un smoke test après publication depuis un environnement vierge.

## 7. Clôturer la publication

- [ ] Vérifier qu'un nouvel utilisateur peut compiler un petit projet avec uniquement le README.
- [ ] Vérifier qu'un build b8x complet passe depuis les artefacts publiés.
- [ ] Publier les limites et les travaux futurs sans les présenter comme terminés.
- [ ] Annoncer la version, les plateformes supportées et le canal de support.
- [ ] Archiver les preuves de release et les résultats de la CI.

## Critères de sortie

- [ ] Clone propre, installation et premier build réussis sans intervention manuelle.
- [ ] CI verte sur toutes les plateformes annoncées.
- [ ] Aucun secret, chemin privé ou artefact interne dans la publication.
- [ ] Licences et attributions complètes.
- [ ] Artefacts reproductibles ou procédure de vérification explicite.
- [ ] Release candidate validée depuis les artefacts publiés.
- [ ] Rollback testé.
- [ ] Documentation et benchmarks conformes à ce qui est réellement mesuré.
