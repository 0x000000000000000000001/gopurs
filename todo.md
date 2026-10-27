# Plan de publication officielle de gopurs

Ce plan suit la préparation de `0.1.0`, sa qualification puis sa publication.
Une case de préparation cochée ne vaut pas validation de la release : les
résultats doivent porter sur les sources et artefacts exacts qui seront publiés.
Le bilan des campagnes est dans `docs/quality-0.1.0.md`.

## 1. Figer le périmètre et le contrat public

- [x] Choisir la version de départ et le schéma de versionnement.
- [x] Lister les surfaces stables : CLI, backend Go, formats intermédiaires, bibliothèques `gopurs-*`.
- [x] Lister explicitement ce qui reste expérimental.
- [x] Fixer les plateformes supportées : OS, architectures, versions de Go et de PureScript.
- [x] Auditer les licences du compilateur, du fork PureScript et de PBO ; consigner séparément les écarts des ports sources dans [l'audit des bibliothèques](docs/library-source-audit.md).
- [x] Ajouter les attributions et notices tierces nécessaires.
- [x] Définir la politique de support et de dépréciation.

## 2. Garantir l'installation depuis zéro

- [x] Tester un clone propre sur une machine sans cache ni binaire préexistant.
- [x] Rendre le build JavaScript indépendant des checkouts voisins.
- [x] Rendre le bootstrap natif et les applications reproductibles depuis des entrées publiques épinglées (fork TAST, PBO, bibliothèques et adaptations Go).
- [x] Épingler les dépendances du build JavaScript.
- [x] Figer et vérifier les révisions du frontend TAST et de tous les ports `gopurs-*` nécessaires à la release.
- [x] Documenter la chaîne de bootstrap : binaire JS, binaire Go précédent, artefacts de référence.
- [ ] Publier les checksums des binaires de bootstrap.
- [x] Vérifier la reproductibilité des sorties, ou documenter précisément les différences.
- [x] Mettre en place une CI Linux et macOS sur les architectures supportées.
- [ ] Vérifier que la CI part d'un environnement vierge, sans cache préchauffé obligatoire.

## 2 bis. Nixifier le projet

- [x] Fournir des sorties `packages`, `apps`, `checks` et `devShells` sur le modèle de phpurs, sur les trois systèmes annoncés.
- [x] Construire un paquet gopurs autonome depuis les sources et les lockfiles, sans dépendance à un checkout voisin.
- [x] Fournir un frontend TAST épinglé et un parcours documenté pour compiler une application avec les ports Go.
- [x] Vérifier localement sur Linux arm64 `nix build`, `nix run` et `nix flake check`, avec un vrai programme PureScript/FFI compilé puis exécuté en Go ; la matrice CI reste à valider ci-dessous.
- [x] Vérifier l'évaluation des trois systèmes sans IFD et archiver les preuves locales Nix ; cette évaluation ne remplace pas leurs builds en CI.
- [ ] Exécuter les validations Nix sur Linux amd64/arm64 et macOS arm64 en CI.
- [x] Documenter les commandes de développement, de build et d'installation Nix et leurs entrées exactes.

## 3. Nettoyer le dépôt

- [x] Retirer les artefacts générés, caches, logs, binaires et sauvegardes du suivi Git.
- [x] Retirer les preuves de benchmark et fichiers d'archive internes des surfaces publiées.
- [x] Rechercher et supprimer les secrets, jetons, URLs privées et chemins personnels.
- [x] Enlever les mentions `WIP` des surfaces officielles ou les convertir en limitations documentées.
- [x] Supprimer le code mort et les options expérimentales visibles.
- [x] Harmoniser le nommage `gopurs` et les intitulés des paquets.
- [x] Vérifier que les fichiers publiés dans les paquets sont limités au nécessaire.
- [x] Éviter tout refactor massif non nécessaire avant la première release.

## 4. Finaliser la documentation

- [x] Écrire un quickstart testé de bout en bout.
- [x] Documenter l'installation, les prérequis et la référence CLI.
- [x] Documenter l'architecture, les formats et le fonctionnement du backend Go.
- [x] Documenter les limitations connues et les différences avec le compilateur officiel.
- [x] Documenter la migration et la compatibilité entre versions.
- [x] Ajouter un guide de dépannage et une FAQ.
- [x] Ajouter `CONTRIBUTING`, `SECURITY`, `CHANGELOG` et le code de conduite.
- [x] Reprendre les benchmarks avec protocole, machine, commandes et scripts reproductibles.
- [x] Vérifier qu'aucune promesse de performance ou de compatibilité n'est ambiguë.

## 5. Passer les portes de qualité

- [ ] Exécuter toute la suite de tests dans un environnement propre.
- [x] Exécuter les 401 fixtures en Release et les 401 en Dev avec le compilateur natif frais : 802/802 passent, snapshots inchangés et exécution Go validée ; preuves archivées dans le bilan qualité.
- [x] Vérifier la conformité avec le compilateur officiel sur un corpus représentatif.
- [x] Vérifier les golden tests, la parité FFI et les formats générés.
- [x] Tester les erreurs, annulations, bornes de configuration mémoire et cas de concurrence, avec les tests Go `-race` et l'émission native.
- [ ] Vérifier les performances à froid et à chaud avec le protocole publié.
- [ ] Faire tourner une release candidate et corriger les régressions bloquantes.
- [ ] Exiger une CI verte sur le tag de release.
- [x] Configurer la CI sur les tags pour couvrir le frontend TAST, le bootstrap natif et l'exécution d'applications, en plus du paquet JavaScript ; son exécution verte reste une porte distincte.

## 6. Préparer les artefacts et la publication

- [ ] Produire les binaires par plateforme avec les flags de release documentés.
- [ ] Générer les checksums, la provenance et un SBOM.
- [ ] Signer le tag et les artefacts.
- [ ] Rédiger les notes de version et le changelog.
- [ ] Tester l'installation uniquement depuis les artefacts publiés.
- [ ] Vérifier les ports sources `gopurs-*` : commits compatibles épinglés, dépendances résolues et notices complètes ; appliquer versions, bornes et allowlists aux paquets effectivement publiés, sans versions communes imposées ([audit et écarts restants](docs/library-source-audit.md)).
- [x] Faire un `npm publish --dry-run` sur l'archive locale préparée.
- [ ] Vérifier la provenance de la publication npm réelle.
- [ ] Préparer la procédure de rollback et conserver les artefacts précédents.
- [x] Aligner le contrat de support avec les parcours réellement distribués : hôte JavaScript npm/Nix et binaires Go ; hôte Rust expérimental.
- [ ] Enregistrer dans Git toutes les sources, configurations et notices nécessaires avant de créer la RC ; vérifier un checkout du commit publié.
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
- [ ] Installation et application Nix validées sur la matrice supportée.
