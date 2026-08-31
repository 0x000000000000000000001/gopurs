import sys

def replace_first(content, search, replace):
    if content.count(search) != 1:
        print(f"ERROR: Found {content.count(search)} occurrences.")
        sys.exit(1)
    return content.replace(search, replace)

with open("src/Gopurs/CodeGen.purs", "r") as f:
    content = f.read()

old_instantiations = """    instantiations = Map.mapMaybe
      ( \\typeMap ->
          let
            set = Set.fromFoldable (Map.keys typeMap)
            safeSet = Set.filter isSafeType set
          in
            if Set.isEmpty safeSet then Nothing else Just safeSet
      )
      rawInstantiations"""

content = replace_first(content, old_instantiations, "")

with open("src/Gopurs/CodeGen.purs", "w") as f:
    f.write(content)
