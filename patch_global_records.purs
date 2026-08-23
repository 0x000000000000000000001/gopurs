globalRecordStructs :: Ref.Ref (Set.Set String)
globalRecordStructs = unsafePerformEffect (Ref.new Set.empty)

globalRecordDecls :: Ref.Ref (Array String)
globalRecordDecls = unsafePerformEffect (Ref.new [])
